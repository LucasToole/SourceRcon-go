/* Copyright (c) 2020 Lucas Toole. See LICENSE for details */

package rcon

import (
	"fmt"
)

const SERVERDATA_AUTH           int32 := 3
const SERVERDATA_AUTH_RESPONSE  int32 := 2
const SERVERDATA_EXECCOMMAND    int32 := 2
const SERVERDATA_RESPONSE_VALUE int32 := 0

func Test() {
	fmt.Println("In Test()")
}

/* Use when you have your own connection code */
func InitRcon() {
	
}

/* For when you don't want to write connection code */
func RconInitConnection() {
	
}

func RconSend() {
	
}

func RconRecieve() {
	
}


