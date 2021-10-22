package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	numDigits, numTickets int
	fileName              string
)

func main() {
	if !dataInput() {
		fmt.Println("Wrong program input")
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rand.Seed(time.Now().Unix())
	for i := 0; i < numTickets; i++ {
		nextRand := rand.Intn(int(math.Pow10(numDigits)))

		_, err := file.WriteString(addZerosIfNeeded(nextRand) + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Check output in file %s\n", fileName)
}

func addZerosIfNeeded(num int) string {
	strNum := strconv.Itoa(num)
	for {
		if len(strNum) >= numDigits {
			break
		}
		strNum = "0" + strNum
	}

	return strNum
}

func dataInput() bool {
	if len(os.Args) != 4 {
		fmt.Println("specify 3 parameters: <numDigits> <numTickets> <fileName>")
		return false
	}

	var err error

	fileName = os.Args[3]
	numDigits, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("<numDigits> to int convert error")
		return false
	}
	numTickets, err = strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("<numTickets> to int convert error")
		return false
	}

	return true
}
