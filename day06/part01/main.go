package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func countOrbits(planet string, orbits map[string][]string) int {
	orbitCount := 0
	for _, orbitee := range orbits[planet] {
		orbitCount = countOrbits(orbitee, orbits) + 1
	}
	return orbitCount
}

func buildNetwork(inputs []string) map[string][]string {
	orbits := make(map[string][]string)
	for _, x := range inputs {
		orbit := strings.Split(x, ")")
		orbits[orbit[1]] = append(orbits[orbit[1]], orbit[0])
	}
	return orbits
}

func getInput(f *os.File) []string {
	var orbits []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		orbits = append(orbits, scanner.Text())
	}
	return orbits
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	inputfile := os.Args[1]
	f, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile(f)
	orbits := buildNetwork(getInput(f))
	orbitTotal := 0
	for orbitor := range orbits {
		orbitTotal += countOrbits(orbitor, orbits)
	}
	fmt.Println(orbitTotal)
}
