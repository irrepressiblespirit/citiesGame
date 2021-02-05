package helper

import (
	"io/ioutil"
	"os"
)

func WriteToFile(fileName string, data []byte) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, er := file.Write(data)
	if er != nil {
		panic(er)
	}
}

func ReadFromFile(fileName string) []byte {
	fContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return fContent
}
