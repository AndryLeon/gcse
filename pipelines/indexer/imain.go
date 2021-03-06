package main

import (
	"log"
	"runtime"

	"github.com/daviddengcn/gcse"
)

func main() {
	runtime.GOMAXPROCS(2)
	log.Println("indexer started...")

	if err := gcse.IndexSegments.ClearUndones(); err != nil {
		log.Printf("ClearUndones failed: %v", err)
	}

	if err := clearOutdatedIndex(); err != nil {
		log.Printf("clearOutdatedIndex failed: %v", err)
	}
	doIndex()

	log.Println("indexer exits...")
}
