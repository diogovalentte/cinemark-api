// Package main implements the main function
package main

import api "github.com/diogovalentte/cinemark-api/src"

func main() {
	router := api.SetupRouter()
	router.SetTrustedProxies(nil)

	router.Run()
}
