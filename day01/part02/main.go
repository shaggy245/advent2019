package main

// aoc 2019 day 1 part 2
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputfile := os.Args[1]
	f, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	var fuel_total int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fuel_hop, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		for fuel_hop > 0 {
			fuel_hop = fuel_hop/3 - 2
			if fuel_hop < 0 {
				fuel_hop = 0
			}
			fuel_total = fuel_total + fuel_hop
		}
	}
	fmt.Println(fuel_total)
	f.Close()
}
