package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func cleanUp(dir string) {
	itens, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range itens {
		if item.IsDir() {
			cleanUp(dir + "/" + item.Name())
		} else {
			duration := time.Now().Sub(item.ModTime())
			month, _ := time.ParseDuration("730h")

			if duration >= month {
				fmt.Println("Removing file:", dir+"/"+item.Name())
				os.Remove(dir + "/" + item.Name())
			}
		}
	}
}

func main() {
	args := os.Args[1:]
	fmt.Println("Argumentos:", args)

	for _, dir := range args {
		cleanUp(dir)
	}
}
