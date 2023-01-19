package main

import (
	"flag"
	"fmt"
	"walk/certificate"
	"walk/config"
)

func main() {
	var method string

	fmt.Println("Hello, world!")
	certificate.CreateCertificate()

	flag.StringVar(&method, "X", config.DefaultWalkConfig.Method, "Method to be used to communicate with the server")
	flag.Parse()

	fmt.Println("Method used for connecting", method)

	url := flag.Arg(0)

	fmt.Println("Trying to fetch data from", url)
}
