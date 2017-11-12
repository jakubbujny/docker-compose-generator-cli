package main

import "io/ioutil"

func WriteOrAppendToFile(fileName string, content string) error {
	return ioutil.WriteFile(fileName, []byte(content), 0644)
}
