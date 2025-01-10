/*
Parse input into two lists of ints. Then sort the lists. Then iterate over them computing
the pairwise distances.
*/
package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"strconv"
	"sort"
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

	sort.Ints(llist)
	sort.Ints(rlist)

	result := 0
	for i := 0; i < len(llist); i++ {
		if llist[i] > rlist[i] {
			result += llist[i] - rlist[i]
		} else {
			result += rlist[i] - llist[i]
		}
	}

	fmt.Printf("%v\n", result)
}