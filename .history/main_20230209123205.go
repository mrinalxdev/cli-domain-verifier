package main

import(
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"os"
)


func main(){
	scanner := bufio .NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")


	for scanner.Scan(){
		checkDomain(scanner.Text())
	}

	if err := scanner.Err
}

func checkDomain (domain string){

}