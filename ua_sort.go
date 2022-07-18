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
	for i := 0; i < 10; i++ {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		heap.Push(h, line)
	}

	for h.Len() > 0 {
		fmt.Print(heap.Pop(h), "\n")
	}
}
