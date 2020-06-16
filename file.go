package main

import "os"

func isExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil || os.IsExist(err)
}
