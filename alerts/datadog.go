package alerts

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"errors"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/jatm80/domain-expiration-checker/checks"
)

func SendToDatadog(data *[]checks.Status) error {
	if _, siteOk := os.LookupEnv("DD_SITE"); siteOk {
		if _, apiKeyOk := os.LookupEnv("DD_API_KEY"); apiKeyOk {
			alertDays := GetEnv("ALERT_THRESHOLD_DAYS", "30,60,90")
			a := strings.Split(alertDays, ",")
			for _, v := range *data {
				for _, d := range a {
					if strconv.Itoa(v.Days) == d {
						err := sendEvent(v)
						if err != nil {
							return err
						}
					}
				}
			}
		} else {
			return errors.New("environment variable `DD_API_KEY` not found")
		}
	} else {
		return errors.New("environment variable `DD_SITE` not found")
	}
	return nil
}

func sendEvent(data checks.Status) error {
	body := datadogV1.EventCreateRequest{
		Title: fmt.Sprintf("Domain expiration checker: %s", data.Name),
		Text:  fmt.Sprintf("Domain %s will expire on the %s ( %d days )", data.Name, data.Date, data.Days),
		SourceTypeName: datadog.PtrString("go"),
		Tags: []string{
			"application:domain-expiration-checker",
			fmt.Sprintf("domain:%s", data.Name),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewEventsApi(apiClient)
	_, _, err := api.CreateEvent(ctx, body)
	if err != nil {
		return err
	}
	return nil
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
