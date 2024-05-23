package main

import "fmt"

func main() {
	// Key-value pair to store data
	// Any value can be used as a key, struct, int etc
	serviceProviders := map[string]string{
		"GCP": "https://console.cloud.google.com",
		"AWS": "https://aws.com",
	}

	fmt.Println(serviceProviders["GCP"])

	// Add new service provider
	serviceProviders["Azure"] = "https://portal.azure.com"
	fmt.Println("Added Azure to the list: ", serviceProviders)

	// Delete existing service provider
	delete(serviceProviders, "GCP")
	fmt.Println("Deleted GCP from the list of service providers: ", serviceProviders)

}
