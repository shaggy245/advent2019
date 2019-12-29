package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getOrbits(planet string, orbits map[string][]string) []string {
	var nodeNetwork []string
	for _, orbitee := range orbits[planet] {
		nodeNetwork = append(nodeNetwork, orbitee)
		nodeNetwork = append(nodeNetwork, getOrbits(orbitee, orbits)...)
	}
	return nodeNetwork
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

func findSantaOrbits(orbits map[string][]string) []string {
	return getOrbits("SAN", orbits)
}

func findYourOrbits(orbits map[string][]string) []string {
	return getOrbits("YOU", orbits)
}

func findOrbitalTransfers(you []string, santa []string) int {
	transfers := 0
	for i, yOrbit := range you {
		for j, sOrbit := range santa {
			if yOrbit == sOrbit {
				transfers = i + j
				break
			}
		}
		if transfers > 0 {
			break
		}
	}
	return transfers
}

func main() {
	inputfile := os.Args[1]
	f, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile(f)
	orbits := buildNetwork(getInput(f))
	santaOrbits := findSantaOrbits(orbits)
	yourOrbits := findYourOrbits(orbits)
	fmt.Println("Minimum transfers: ", findOrbitalTransfers(yourOrbits, santaOrbits))
}
