package s

import (
	"bytes"
)

const (
	MessageTypeAck    = 0x0
	MessageTypeStream = 0x1
)

const Version = 0x0

func CreateTransaction(pyld []byte, conn *Connection) ([][]byte, error) {
	buf := bytes.NewBuffer(pyld)
	var messages [][]byte
	checksum := generateChecksum(pyld)
	frag := 0
	// Transaction will take multiple messages
	for {
		var m Message
		if frag == 0 { // TODO MOVE COMPRESSION  HERE TO FILL A PACKET TO MAX PAYLOAD DURING COMPRESSION INSTEAD OF JUST MAKING THE SAME CNT OF SMALLER PACKETS
			m = Message{Header{true, conn.Compress, MessageTypeStream, Version, conn.ID, conn.SessionID, conn.LastFrame, uint8(frag)}, buf.Next(int(conn.MaximumPayload) - 7), uint16(len(pyld)), checksum}
		} else {
			m = Message{Header{true, conn.Compress, MessageTypeStream, Version, conn.ID, conn.SessionID, conn.LastFrame, uint8(frag)}, buf.Next(int(conn.MaximumPayload) - 4), 0x0, 0x0}
		}
		marshalled, err := m.Marshal()
		if err != nil {
			return messages, err
		}
		messages = append(messages, marshalled)
		if buf.Len() == 0 {
			break
		}
		frag++
	}

	conn.LastFrame++

	// The returned messages can be sent asap to a devices waiting for s data in the same session
	// be sure to order the messages correctly, starting from element 0
	return messages, nil
}

func generateChecksum(pyld []byte) byte {
	var c byte
	for _, i := range pyld {
		c = c ^ i
	}
	return c
}
