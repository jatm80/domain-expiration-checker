package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

type Status struct {
	Name string
	ExpirationDate string
	DaysToExpire int
}


func main() {

  if len(os.Args) > 0 {
	domains := strings.Split(os.Args[1], ",")
	
	for _,d := range domains {
		result, err:=checkDomain(d)
		if err != nil {
			log.Fatal(err)
			}
			fmt.Printf("%#v \n",result)
		}
 	} else {
		log.Fatalln("Invalid number of arguments,  need to pass at least one domain to check")
	}
}

func checkDomain(name string) (*Status, error) {
	whois_raw, err := whois.Whois(name)
	if err != nil {
		return nil, err
	}


	result, err := whoisparser.Parse(whois_raw)
	if err != nil {
		return nil, err
	}

    days, err := calculateDays(result.Domain.ExpirationDate)
	if err != nil {
		return nil, err
	}

	return &Status{
		Name: name,
		ExpirationDate: result.Domain.ExpirationDate,
		DaysToExpire: days,
	},nil
}

func calculateDays(date string) (int,error) {
	parseDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return 0,errors.New("no expiration date available for this domain")
	}
	now := time.Now()
	duration := parseDate.Sub(now)
	days := int(duration.Hours() / 24)

	return days,nil
}