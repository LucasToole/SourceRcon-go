/* Copyright (c) 2020 Lucas Toole. See LICENSE for details */

package rcon

import (
	"math/rand"
	"bytes"
)

type rconPacket struct {
	size int32
	id int32
	reqType int32
	body string
}


func constructPacket(reqType int32, body string) *rconPacket{
	return &rconPacket {
		size: int32(len(body) + 10),
		id: rand.Int31(),
		reqType: reqType,
		body: body,
	}
}

/* Rcon requires all data be in Little Endian */
func encodePacket(p rconPacket) []byte {
	// TODO: This.
	
}

func encodePacket() {
	
}
