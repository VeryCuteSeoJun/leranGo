package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	kbReader := bufio.NewReader(os.Stdin)
	fmt.Print("나이를 입력하세요 : ")
	strAge, err := kbReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	strAge = strings.TrimSpace(strAge)
	iAge, err := strconv.Atoi(strAge)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("당신의 나이 : %d세\n", iAge)
}
