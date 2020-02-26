/* Copyright (c) 2020 Lucas Toole. See LICENSE for details */

package rcon

import (
	"math/rand"
	"bytes"
	"encoding/binary"
	"time"
)

type rconPacket struct {
	size int32
	id int32
	reqType int32
	body string
}


func constructPacket(reqType int32, body string) *rconPacket{
	rand.Seed(time.Now().UnixNano())
	return &rconPacket {
		size: int32(len(body) + 10),
		id: rand.Int31(),
		reqType: reqType,
		body: body,
	}
}

/* Rcon requires all data be in Little Endian */
func encodePacket(p *rconPacket) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, p.size)
	err = binary.Write(buf, binary.LittleEndian, p.id)
	err = binary.Write(buf, binary.LittleEndian, p.reqType)
	_, err = buf.WriteString(p.body) /* Ignore String size */
	err = buf.WriteByte(0)
	err = buf.WriteByte(0)

	return buf.Bytes(), err
	
}

func decodePacket(buf []byte) (*rconPacket) {
	buffer := bytes.NewBuffer(buf)
	var binsize int32
	var binid int32
	var bintype int32

	binary.Read(buffer, binary.LittleEndian, &binsize)
	binary.Read(buffer, binary.LittleEndian, &binid)
	binary.Read(buffer, binary.LittleEndian, &bintype)
	binbody := string(buf[15:cap(buf)])
	
	return &rconPacket {
		size: binsize,
		id: binid,
		reqType: bintype,
		body: binbody,
	}
}
