/* Copyright (c) 2020 Lucas Toole. See LICENSE for details */

package rcon

import (
	"fmt"
	"net"
)

const AUTH           int32 = 3
const AUTH_RESPONSE  int32 = 2
const EXECCOMMAND    int32 = 2
const RESPONSE_VALUE int32 = 0

/* Use when you have your own connection code */
func InitRcon(conn net.Conn, password string) (int32){
	RconSend(conn, AUTH, password)

	var recvid, recvType int32
	for {
		recvid, recvType, _ = RconRecieve(conn)
		if recvType == 2 {
			break
		}
	}
	if recvid == -1 {
		fmt.Println("Rcon Auth Failed")
	}

	return recvid
}

/* For when you don't want to write connection code */
func RconInitConnection(address, port, password string) (net.Conn, int32, error) {
	conn, err := net.Dial("tcp4", (address + ":" + port))

	if err != nil {
		fmt.Println(err)
		return conn, -2, err /* Placeholder error number... Maybe */
	}

	authAcceptance := InitRcon(conn, password)
	
	return conn, authAcceptance, err
}

func RconSend(conn net.Conn, reqType int32, body string) (int32){
	packet := constructPacket(reqType, body)

	binpacket, err := encodePacket(packet)
	
	if err != nil {
		fmt.Println(err)
	}

	conn.Write(binpacket)

	return packet.id
	
}

func RconRecieve(conn net.Conn) (int32, int32, string){
	buf := make([]byte, 4096)
	
	num, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	if num > 4096 { // TODO: Multi-Packet Response
		fmt.Println("Response too large for single packet.")
		fmt.Println("Current version does not support multi-packet responses.")
	}
	packet := decodePacket(buf)

	return packet.id, packet.reqType, packet.body
}


