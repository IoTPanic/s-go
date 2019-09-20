package s

import "bytes"

const (
	MessageTypeAck    = 0x0
	MessageTypeStream = 0x1
)

const Version = 0x0

const MaxFrameSize = 1499

func CreateTransaction(pyld []byte, conn *Connection) ([][]byte, error) {
	buf := bytes.NewBuffer(pyld)
	var messages [][]byte
	checksum := generateChecksum(pyld)

	// Transaction will take multiple messages
	for frag := 0; ; frag++ {
		var m Message
		if frag == 0 {
			m = Message{Header{true, conn.Compress, MessageTypeStream, Version, conn.ID, conn.SessionID, conn.LastFrame, uint8(frag)}, buf.Next(MaxFrameSize - 7), uint16(len(pyld)), checksum}
		} else {
			m = Message{Header{true, conn.Compress, MessageTypeStream, Version, conn.ID, conn.SessionID, conn.LastFrame, uint8(frag)}, buf.Next(MaxFrameSize - 4), 0x0, 0x0}
		}
		marshalled, err := m.Marshal()
		if err != nil {
			return messages, err
		}
		messages = append(messages, marshalled)
		if buf.Len() == 0 {
			break
		}
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
