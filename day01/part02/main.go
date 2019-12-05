package main

// aoc 2019 day 1 part 2
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calculateFuel(masses []int) int {
	var fuel_total int
	for _, fuel_hop := range masses {
		for fuel_hop > 0 {
			fuel_hop = fuel_hop/3 - 2
			if fuel_hop < 0 {
				fuel_hop = 0
			}
			fuel_total = fuel_total + fuel_hop
		}
	}
	return fuel_total
}

func main() {
	inputfile := os.Args[1]
	f, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	var masses []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		masses = append(masses, mass)
	}

	fuel_total := calculateFuel(masses)

	fmt.Println(fuel_total)
	f.Close()
}
