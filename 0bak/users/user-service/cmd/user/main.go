// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 8907ffca23
// Version Date: Wed 27 Nov 2019 21:28:21 UTC

package main

import (
	"flag"

	// This Service
	"github.com/sabnak227/jwt-demo/bak/users/user-service/svc/server"
)

func main() {
	// Update addresses if they have been overwritten by flags
	flag.Parse()

	server.Run(server.DefaultConfig)
}
