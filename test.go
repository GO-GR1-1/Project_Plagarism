package main

import (
	"fmt"
    "io/ioutil"
    "log"
    "strings"
)

func textFilesInDirectory() []string {
	files, err := ioutil.ReadDir(".")
    if err != nil {
        log.Fatal(err)
    }
	
	var textFiles []string
    
    for _, file := range files {
        if strings.Contains(file.Name(), ".txt") {
        	textFiles = append(textFiles, file.Name())
        }
    }

    return textFiles
}