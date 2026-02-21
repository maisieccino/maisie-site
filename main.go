package main

import "github.com/maisieccino/maisie-site/cmd/api"

func main() {
	if err := api.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
