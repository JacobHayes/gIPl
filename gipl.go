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
	check(scanner.Err())
	scanner.Scan()
	line := scanner.Text()
	check(scanner.Err())

	return line
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	key_file := open(`api`)
	defer key_file.Close()
	key := readLine(bufio.NewScanner(key_file))

	fmt.Print("Enter an IP address to lookup: ")
	ip := readLine(stdin)

	fmt.Print("Location Precision - City/[Country]: ")

	if precision := strings.ToLower(readLine(stdin)); precision == "city" {
		fmt.Println(locus.LookupLocation(ip, "City", key))
	} else {
		fmt.Println(locus.LookupLocation(ip, "Country", key))
	}
}
