package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jatm80/domain-expiration-checker/alerts"
	"github.com/jatm80/domain-expiration-checker/checks"
)

func main() {

	if domains, exist := os.LookupEnv("DOMAINS"); exist {
		d := strings.Split(domains, ",")
		c := checks.Domains{
			Name: d,
		}
		r, err := c.CheckExpiration()
		if err != nil {
			log.Fatalln(err)
		}

		if alerts.GetEnv("DATADOG_ENABLED","false") == "true" {
			err = alerts.SendToDatadog(r)
			if err != nil {
				log.Println(err)
			}
		}

		fmt.Printf("%#v \n", *r)

	} else {
		log.Fatalln("environment variable DOMAINS not found, export DOMAINS=\"command separated domain list\"")
	}
}
