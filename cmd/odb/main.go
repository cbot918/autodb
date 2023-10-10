/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/cbot918/autodb/cmd/odb/cmd"
)

func main() {

	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
