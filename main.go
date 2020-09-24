package main

import (
	"NewsAppV2/api"
)

func main() {
	s := api.Server{}
	s.Serve()
}
