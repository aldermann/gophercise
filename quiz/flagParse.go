package main

import "flag"

var norand bool
var questionCount int
var timeOut int
var filePath string

func flagParse() {
	flag.BoolVar(&norand, "norand", false, "Specify whether to shuffle the questions")
	flag.IntVar(&questionCount, "n", 0, "Specify the number of questions. Input 0 to use all the questions.")
	flag.IntVar(&timeOut, "t", 30, "Specify the time for each question. Default is 30s")
	flag.StringVar(&filePath, "p", "problem.csv", "Specify the CSV file that contain the questions. Input 'mock' to use the mock questions instead")
	flag.Parse()
}
