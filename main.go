package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var usage int8
	dir := flag.String("dir", "", "dir where be scanning")
	fileName := flag.String("file", "", "file name for scanning")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage -flag argument \n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *dir != "" {
		if *fileName != "" {
			data, err := os.ReadDir(*dir)

			if err != nil {
				fmt.Println("error : ", err)
				return
			}

			for _, entry := range data {

				if entry.Name() == *fileName {
					usage++
					fmt.Println("i find needed file :", entry.Name())
					continue
				}
			}
			if usage != 0 {
				fmt.Println("i find", usage, "with", *fileName, "name")
			}
		}
	}
}
