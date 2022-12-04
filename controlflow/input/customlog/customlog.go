package customlog

import (
	"fmt"
	"log"
	"os"
	"path"
)

var (
	LOGFILE = path.Join(os.TempDir(), "mGO.log")
)

func CreateLog() {

	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	flags := log.Lshortfile | log.Ldate | log.Ltime
	iLog := log.New(f, "iLog", flags)
	iLog.Println("From custom log")

	iLog.SetFlags(log.Lshortfile | log.LstdFlags)
	iLog.Println("Another log entry")

}
