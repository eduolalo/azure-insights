package main

import (
	"os"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

var telemetryClient appinsights.TelemetryClient
var key string = os.Getenv("APP_INSIGHTS_KEY")

func main() {

	config := appinsights.NewTelemetryConfiguration(key)
	config.EndpointUrl = "https://southcentralus-3.in.applicationinsights.azure.com/"
	telemetryClient = appinsights.NewTelemetryClientFromConfig(config)
	telemetryClient.Context().Tags.Cloud().SetRole("Moldava")
	telemetryClient.TrackTrace("Hello World!", appinsights.Information)
}
