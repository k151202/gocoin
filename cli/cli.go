package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/k151202/gocoin/explorer"
	"github.com/k151202/gocoin/rest"
)

func usage(){
	fmt.Printf("Welcome to Nomad Coin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port: Set the PORT of the sever\n")
	fmt.Printf("-mode: Choose between 'html' and 'rest'\n")
	fmt.Printf("!! for those want two servers run simultaneously, choose 'both'!!\n")
	os.Exit(0)
}

func Start(){
	if len(os.Args) < 2 {
		usage()
	}
	
	port := flag.Int("port", 4000, "Set port of the server")

	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")


	flag.Parse()

	switch *mode{
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "both":
		go rest.Start(*port)
		explorer.Start(*port + 1)
	default: usage()
	}
}