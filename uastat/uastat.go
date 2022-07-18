package uastat

import (
	"log"
	"strconv"
	"time"
)

type Record struct {
	// UserAgent  string
	TimesSeen int64
	// LastSeenTS int64
	value int64
}

const (
	layout = "2006-01-02 15:04:05.000000"
)

type UAStatHeap []Record

// Len, Less, Swap для реализации интерфейса sort.Interface
func (h UAStatHeap) Len() int           { return len(h) }
func (h UAStatHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h UAStatHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *UAStatHeap) Push(linei interface{}) {
	line := linei.([]string)

	var rec Record

	// rec.UserAgent = line[1]

	intField, err := strconv.ParseInt(line[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	rec.TimesSeen = intField

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
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
