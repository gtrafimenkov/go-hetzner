// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2018 Gennady Trafimenkov <gennady.trafimenkov@gmail.com>

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gtrafimenkov/go-hetzner"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s username password\n", os.Args[0])
		os.Exit(1)
	}

	username, password := os.Args[1], os.Args[2]

	client := hetzner.NewClient(username, password)
	servers, _, err := client.Server.ListServers()
	if err != nil {
		log.Fatal(err)
	}

	for _, srv := range servers {
		fmt.Printf("%s\t%s\n", srv.ServerName, srv.ServerIP)
	}
}
