package main

import (
	"log"
	"os"
)

func main() {
	if err := root().Execute(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
