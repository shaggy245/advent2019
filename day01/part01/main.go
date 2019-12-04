package main
// aoc 2019 day 1 part 1
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
		fuel_total = fuel_total + (fuel_hop/3 - 2)
	}
	fmt.Println(fuel_total)
	f.Close()
}
