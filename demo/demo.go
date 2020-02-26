/* Copyright (c) 2020 Lucas Toole. See LICENSE for details */

package main

import (
	"../rcon"
)

func main() {
	conn := rcon.RconInitConnection("10.0.0.128", "27015", "test")
	rcon.RconSend(conn, rcon.EXECCOMMAND, "say Hello")
}
 
