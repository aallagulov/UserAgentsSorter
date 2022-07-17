package main

import (
	"container/heap"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	layout = "2006-01-02 15:04:05.000000"
)

type UAStatRecord struct {
	// UserAgent  string
	// TimesSeen  int64
	// LastSeenTS int64
	value int64
}

type UAStatHeap []UAStatRecord

// Len, Less, Swap для реализации интерфейса sort.Interface
func (h UAStatHeap) Len() int           { return len(h) }
func (h UAStatHeap) Less(i, j int) bool { return h[i].value > h[j].value }
func (h UAStatHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *UAStatHeap) Push(linei interface{}) {
	line := linei.([]string)

	var rec UAStatRecord

	// rec.UserAgent = line[1]

	intField, err := strconv.ParseInt(line[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	// rec.TimesSeen = intField

	time, err := time.Parse(layout, line[37])
	if err != nil {
		log.Fatal(err)
	}
	// rec.LastSeenTS = time.Unix()

	// rec.value = rec.TimesSeen*10000000000 + rec.LastSeenTS
	rec.value = intField*10000000000 + time.Unix()
	*h = append(*h, rec)
}

func (h *UAStatHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

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

	h := &UAStatHeap{}
	heap.Init(h)

	for i := 0; i < 7; i++ {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		heap.Push(h, line)
	}

	l := h.Len()
	for i := 0; i < l; i++ {
		fmt.Print(h, "\n")
		fmt.Print(h.Pop(), "\n")
		fmt.Print(h, "\n\n")
	}
}
