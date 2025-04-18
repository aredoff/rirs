package main

import (
	"log"

	"github.com/aredoff/rirs"
	"github.com/aredoff/rirs/fs"
)

func main() {

	folder, err := fs.New("/tmp/rirs")
	if err != nil {
		log.Fatal(err)
	}

	rir, err := rirs.New(folder)
	if err != nil {
		log.Fatal(err)
	}

	err = rir.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
