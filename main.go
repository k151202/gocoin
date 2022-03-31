package main

import (
	"github.com/k151202/gocoin/cli"
	"github.com/k151202/gocoin/db"
)

func main(){
	defer db.Close()
	cli.Start()
}