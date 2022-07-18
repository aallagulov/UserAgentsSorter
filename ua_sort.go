package main

import (
	uastat "UserAgentsSorter/uastat"
	"container/heap"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filePtr := flag.String("file", "data/whatismybrowser-user-agent-database.csv", "path to file with useragents data")
	cntPtr := flag.Int("cnt", 10, "how many top useragents to collect")
	flag.Parse()

	f, err := os.Open(*filePtr)
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

	h := &uastat.Heap{}
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
		if h.Len() > *cntPtr {
			heap.Pop(h)
		}
	}

	for h.Len() > 0 {
		r := heap.Pop(h)
		p := r.(uastat.Record)
		fmt.Print(p.TimesSeen, "\t", p.UserAgent, "\n")
	}
}
