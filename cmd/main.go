package main

import (
	"fmt"
	"log"
)

func main() {
	url, err := GetCurrentSafariURL()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
