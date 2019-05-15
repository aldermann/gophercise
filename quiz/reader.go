package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"math/rand"
	"os"
	"time"
)

type Inputter interface {
	Input() []problem
}

type MockInputter struct {}

func (MockInputter) Input() (res []problem) {
	res = []problem{{"1+1", "2"}, {"2+2", "4"}}
	return
}

type CSVInputter struct {
	filePath string
	randomized bool
}

func (c CSVInputter) Input() (res []problem) {
	csvFile, err := os.Open(c.filePath)
	CheckError(err)
	res = make([]problem, 0, 20)
	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		CheckError(err)
		res = append(res, problem{record[0], record[1]})
	}
	if c.randomized {
		rand.Seed(time.Now().UTC().UnixNano())
		rand.Shuffle(len(res), func(i, j int) {
			res[i], res[j] = res[j], res[i]
		})
	}
	return
}

