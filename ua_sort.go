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
	line, err := csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}

	h := &uastat.UAStatHeap{}
	heap.Init(h)

	for line != nil {
		line, err = csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		heap.Push(h, line)
		if h.Len() > 10 {
			heap.Pop(h)
		}
	}

	for h.Len() > 0 {
		r := heap.Pop(h)
		p := r.(uastat.Record)
		fmt.Print(p.TimesSeen, "\t", p.UserAgent, "\n")
	}
}
