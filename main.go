package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}

	words := strings.Fields(wr.String())
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}
