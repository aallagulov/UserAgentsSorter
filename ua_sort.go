package main

import (
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
	UserAgent  string
	TimesSeen  int64
	LastSeen   time.Time
	LastSeenTS int64
	value      int64
}

func appendUAStatRecord(uaStatList *[]UAStatRecord, line []string) {
	var rec UAStatRecord

	rec.UserAgent = line[1]

	intField, err := strconv.ParseInt(line[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	rec.TimesSeen = intField

	time, err := time.Parse(layout, line[37])
	// intField, err := strconv.ParseInt(field, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	rec.LastSeen = time
	rec.LastSeenTS = time.Unix()

	rec.value = rec.TimesSeen*10000000000 + rec.LastSeenTS

	*uaStatList = append(*uaStatList, rec)
}

func main() {
	f, err := os.Open("data/whatismybrowser-user-agent-database.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	var uaStatList []UAStatRecord

	// skipping first "headers" line
	_, err = csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 2; i++ {
		data, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(data[2], "\t", data[37], "\t", data[1], "\n")
		appendUAStatRecord(&uaStatList, data)
	}

	fmt.Print(uaStatList, "\n")
}
