package parsing

import (
	"io"
	"log"
	"os"
)

func CustomWrite(f string) (io.Writer, *os.File) {

	file, err := os.Create(f)
	if err != nil {
		log.Fatalln(err)
	}

	dest := io.MultiWriter(os.Stdout, file)

	return dest, file
}
