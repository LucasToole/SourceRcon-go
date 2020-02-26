/* Copyright (c) 2020 Lucas Toole. See LICENSE for details */

package main

import (
	"github.com/LucasToole/SourceRcon-go/rcon"
)

func main() {
	conn := rcon.RconInitConnection("10.0.0.128", "27015", "test")
	rcon.RconSend(conn, rcon.EXECCOMMAND, "say Hello")
}
 
