package main

import (
	"log"
	"os"
)

func isExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil || os.IsExist(err)
}

func removeFile(fp string) {
	if isExists(fp) {
		log.Println(fp)
		if err := os.RemoveAll(fp); err != nil {
			log.Println("remove project path ", fp, " failure")
			log.Println(err)
		}
	}
}
