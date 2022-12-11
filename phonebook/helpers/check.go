package helpers

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func MatchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}

// Initialized by the user – returns a pointer
// If it returns nil, there was an error

func Inits(N, S, T string) *Entry {

	//both of them should have a value

	if T == "" || S == "" {

		return nil
	}

	// Give Lastyaccess the value

	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Entry{Name: N, Surname: S, Tel: T, LastAccess: LastAccess}
}

func Insert(s *Entry) error {

	//if already exists do not add it
	_, ok := index[(*s).Tel]
	if ok {
		return fmt.Errorf("%s already exists", s.Tel)
	}
	return nil

}

func Search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func List() {
	for _, v := range data {
		fmt.Println(v)
	}
}
