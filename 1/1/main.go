package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum uint64

	for scanner.Scan() {
		line := scanner.Text()
		weight, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		weight = weight / 3
		weight = weight - 2

		sum += uint64(weight)
	}

	fmt.Println(sum)
}
