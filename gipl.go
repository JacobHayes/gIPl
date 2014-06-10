package main

import (
	"bufio"
	"fmt"
	"github.com/JacobHayes/locus"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func open(filename string) *os.File {
	file, err := os.Open(filename)
	check(err)

	return file
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	check(scanner.Err())
	line := scanner.Text()
	check(scanner.Err())

	return line
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	key_file := open(`api`)
	defer key_file.Close()
	key_scanner := bufio.NewScanner(key_file)
	key := readLine(key_scanner)

	fmt.Print("Location Precision - City/[Country]: ")
	locations, err := locus.LookupLocationsFile(`ips`, strings.ToLower(readLine(stdin)), key)
	check(err)

	for _, location := range locations {
		fmt.Println(location)
	}
}
