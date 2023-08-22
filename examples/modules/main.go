package main

import (
	"log"
)

func main() {
	modules, err := New()
	if err != nil {
		log.Fatalln(err.Error())
	}
	modules.Get().Range(func(key, value any) bool {
		log.Println(key, value)
		return true
	})
}
