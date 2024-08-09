# Domain Expiration Checker

## Overview

The **Domain Expiration Checker** is a lightweight application written in Go designed to monitor the expiration dates of specified domain names. The application sends the information to Datadog for monitoring and alerting.

## Environment Variables

To configure the application, the following environment variables need to be set:

- **`DD_SITE`**: The Datadog site URL. Example: `"datadoghq.com"`
- **`DD_API_KEY`**: Your Datadog API key for authentication and data submission.
- **`DATADOG_ENABLED`**: Set to `"true"` to enable Datadog integration. Default: false
- **`DOMAINS`**: A comma-separated list of domain names to monitor. Example: `"example1.com,example2.com,example3.com"`
- **`ALERT_THRESHOLD_DAYS`**: A comma-separated list of days to alert before domain expires. Default: "30,60,90"

## Usage

1. **Setup Environment Variables**: Set the required environment variables in your Kubernetes CronJob or deployment configuration.

2. **CronJob Example**:

   ```yaml
   apiVersion: batch/v1
   kind: CronJob
   metadata:
     name: domain-expiration-checker
   spec:
     schedule: "1 1 * * *"
     jobTemplate:
       spec:
         template:
           spec:
             containers:
             - name: domain-expiration-checker
               image: jatm80/domain-expiration-checker:1
               imagePullPolicy: IfNotPresent
               env:
               - name: DD_SITE
                 value: "datadoghq.com"
               - name: DD_API_KEY
                 value: "your-api-key"
               - name: DATADOG
                 value: "true"
               - name: DOMAINS
                 value: "example1.com,example2.com,example3.com"
             restartPolicy: OnFailure
   ```

3. **Running**: The application runs according to the defined schedule, checking domain expirations and sending metrics to Datadog if enabled.

## Monitoring and Alerts

Ensure that your Datadog integration is properly configured to receive and display metrics from the application. You can set up custom dashboards and alerts based on the expiration dates of the monitored domains.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.