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

	if err := scanner.Err(); err != nil{
		log.Fatal("Error: could not read from the input: %v\n", err)
	}
}

func checkDomain (domain string){

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err!= nil {
		log.Printf("Error: %v\n", err)
	}
	if len()
}