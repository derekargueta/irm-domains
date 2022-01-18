package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scannedtext, err := os.Open("scannedtext.txt")
	fixedtext, err2 := os.OpenFile("fixedtext.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) //       path 3
	if err != nil {
		fmt.Println("djhfdksjfh")
	}
	if err2 != nil {
		fmt.Println("djhfdksjfh")
	}
	reader := bufio.NewScanner(scannedtext)
	fixedtextwriter := bufio.NewWriter(fixedtext)

	for reader.Scan() {
		if strings.Count(reader.Text()[8:len(reader.Text())], "/") == 0 {
			fixedtextwriter.WriteString(reader.Text()[8:len(reader.Text())] + "\n")
			fmt.Println(reader.Text()[8:len(reader.Text())])

		}

		fmt.Println(strings.Count(reader.Text()[8:len(reader.Text())], "/"))

	}
	fixedtextwriter.Flush()
}
