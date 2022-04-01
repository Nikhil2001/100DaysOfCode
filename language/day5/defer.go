package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Readfile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	return ioutil.ReadAll(f)

}

func main() {

	data, _ := Readfile("sample.txt")
	fmt.Println(string(data))

}
