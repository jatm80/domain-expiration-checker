package checks

import (
	"errors"
	"time"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

type Status struct {
	Name string
	Date string
	Days int
}

type Domains struct {
	Name []string
}

func (c *Domains) CheckExpiration() (*[]Status, error) {

	var r []Status
	var s Status

	for _, d := range c.Name {

		s.Name = d

		whois_raw, err := whois.Whois(d)
		if err != nil {
			return nil, err
		}

		result, err := whoisparser.Parse(whois_raw)
		if err != nil {
			return nil, err
		}

		days, err := calculateDays(result.Domain.ExpirationDate)
		if err != nil {
			s.Date = "NA"
			s.Days = -1
		} else {
			s.Date = result.Domain.ExpirationDate
			s.Days = days
		}

		r = append(r, s)
	}

	return &r, nil
}

func calculateDays(date string) (int, error) {
	parseDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return -1, errors.New("no expiration date available for this domain")
	}
	now := time.Now()
	duration := parseDate.Sub(now)
	days := int(duration.Hours() / 24)

	return days, nil
}
