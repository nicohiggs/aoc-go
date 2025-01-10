/*
Parse input into two lists of ints. Then iterate over the second list building a map[int]int of its
values and how often they appear. Then iterate over the first lists values checking the map and using
its matching values to compute the similarity score.
*/
package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	llist := []int{}
	rlist := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lnum, err := strconv.Atoi(line[:5])
		if err != nil {
			log.Fatal(err)
		}
		rnum, err := strconv.Atoi(line[8:])
		if err != nil {
			log.Fatal(err)
		}

		llist = append(llist, lnum)
		rlist = append(rlist, rnum)
	}

	lrmap := make(map[int]int)
	for i := 0; i < len(rlist); i++ {
		value, exists := lrmap[rlist[i]]
		if exists {
			lrmap[rlist[i]] = value + 1
		} else {
			lrmap[rlist[i]] = 1
		}
	}

	result := 0
	for i := 0; i < len(llist); i++ {
		value, exists := lrmap[llist[i]]
		if exists {
			result += llist[i] * value
		}
	}

	fmt.Printf("%v\n", result)
}