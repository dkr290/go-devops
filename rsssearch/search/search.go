package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	//retreive a list of feeds to search through

	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// Create a unbuffered channel to receive match results.

	results := make(chan *Result)
	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	//set up a number of goroutines we need to wait dfor a while
	//they process the individual feeds

	waitGroup.Add(len(feeds))

	//launch goroutine for each feed to find results
	for _, feed := range feeds {

		//retreive a matcher for the search
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		//Launch the go routine to perform the search
		go func(m Matcher, feed *Feed) {
			Match(m, feed, searchTerm, results)
			waitGroup.Done()

		}(matcher, feed)
	}

	//Launch a goroutine to monitor when all the work is done
	go func() {
		//wait for everything to be processed
		waitGroup.Wait()
		//close the channel to signal the Display
		//function that we can exit the program
		close(results)
	}()

	// Start displaying results as they are available and
	// return after the final result is displayed.
	Display(results)

}

// Register is called to register a matcher for use by the program.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
