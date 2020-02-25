/* Copyright (c) 2020 Lucas Toole. See LICENSE for details */

package rcon

import (
	"fmt"
	"net"
	"os"
	"bytes"
)

const SERVERDATA_AUTH           int32 = 3
const SERVERDATA_AUTH_RESPONSE  int32 = 2
const SERVERDATA_EXECCOMMAND    int32 = 2
const SERVERDATA_RESPONSE_VALUE int32 = 0

func Test() {
	fmt.Println("In Test()")
}

/* Use when you have your own connection code */
func InitRcon(conn Conn, password string) {
	RconSend(conn, SERVERDATA_AUTH, password)
}

/* For when you don't want to write connection code */
func RconInitConnection(address, port, password string) Conn{
	
	conn, err := net.Dial("tcp4", (address + ":" + port))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	InitRcon(conn, password)
	
	return conn
}

func RconSend(conn Conn, reqType int32, body string) {
	var p rconPacket
	p = constructPacket(reqType, body)
	
}

func RconRecieve() {
	
}


