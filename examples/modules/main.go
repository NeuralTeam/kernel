package main

import (
	"log"
)

func main() {
	modules, err := New()
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(modules.Get())
}
