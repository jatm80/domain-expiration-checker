package metrics

import (
	"context"
	"os"
	"time"

	"errors"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/jatm80/domain-expiration-checker/checks"
)

func SendToDatadog(data *[]checks.Status) error{
	if _, siteOk := os.LookupEnv("DD_SITE"); siteOk {
		if _, apiKeyOk := os.LookupEnv("DD_API_KEY"); apiKeyOk {
			for _,v := range *data {
				err := submit(v)
				if err != nil {
					return err
				}
			}
		}
	} else {
		return errors.New("environment variable `DD_SITE` or `DD_API_KEY` not found")
	}
	return nil
}

func submit(data checks.Status) error{
		body := datadogV2.MetricPayload{
			Series: []datadogV2.MetricSeries{
				{
					Metric: "domain.expiration.check",
					Type:   datadogV2.METRICINTAKETYPE_GAUGE.Ptr(),
					Points: []datadogV2.MetricPoint{
						{
							Timestamp: datadog.PtrInt64(time.Now().Unix()),
							Value:     datadog.PtrFloat64(float64(data.Days)),
						},
					},
					Resources: []datadogV2.MetricResource{
						{
							Name: datadog.PtrString(data.Name),
							Type: datadog.PtrString("domain"),
						},
					},
				},
			},
		}
		ctx := datadog.NewDefaultContext(context.Background())
		configuration := datadog.NewConfiguration()
		apiClient := datadog.NewAPIClient(configuration)
		api := datadogV2.NewMetricsApi(apiClient)
		_, _, err := api.SubmitMetrics(ctx, body, *datadogV2.NewSubmitMetricsOptionalParameters())
	
		if err != nil {
			return err
		}
    return nil
}
