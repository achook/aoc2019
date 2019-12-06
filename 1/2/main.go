package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculateFuel(mass uint64) uint64 {
	mass /= 3

	if mass < 2 {
		return 0
	}

	mass -= 2
	mass += calculateFuel(mass)

	return mass
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum uint64

	for scanner.Scan() {
		line := scanner.Text()
		weight, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			panic(err)
		}

		weight = calculateFuel(weight)

		sum += weight
	}

	fmt.Println(sum)
}
