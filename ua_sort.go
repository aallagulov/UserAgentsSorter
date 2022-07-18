package main

import (
	uastat "UserAgentsSorter/uastat"
	"container/heap"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data/whatismybrowser-user-agent-database.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	// skipping first "headers" line
	_, err = csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}

	h := &uastat.UAStatHeap{}
	heap.Init(h)
	fmt.Print("PUSH\n")
	for i := 0; i < 5; i++ {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(h, "\n")
		fmt.Print("add this ", line[2], "\n")
		heap.Push(h, line)
		fmt.Print(h, "\n\n")
	}

	l := h.Len()
	fmt.Print("POP\n")
	for i := 0; i < l; i++ {
		fmt.Print(h, "\n")
		fmt.Print("pop this - ", h.Pop(), "\n")
		fmt.Print(h, "\n\n")
	}
}
