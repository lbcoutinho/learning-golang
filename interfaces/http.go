package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type fileWriter struct {}

func (fileWriter) Write(bs []byte) (int, error)  {
	err := ioutil.WriteFile("writerData.txt", bs, os.ModePerm)
	if err != nil {
		return 0, err
	}
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Version 1
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	// Version 2
	//io.Copy(os.Stdout, resp.Body)

	// Version 3 with Writer type
	fr := fileWriter{}
	written, err := io.Copy(fr, resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} else {
		fmt.Println(written, "bytes written to the file")
	}
}
