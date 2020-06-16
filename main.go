package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("you must specify directory")
		return
	}

	dir := os.Args[1]
	if dirs, err := ioutil.ReadDir(dir); err != nil {
		log.Println(err)
		return
	} else if err := findProject(dir, dirs); err != nil {
		log.Println(err)
	}
}
