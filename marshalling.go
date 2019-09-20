package s

import (
	"bytes"
	"errors"

	"gopkg.in/kothar/brotli-go.v0/enc"
)

func (h Header) Marshal() ([]byte, error) {
	p := make([]byte, 5)
	if h.Version > 7 {
		return []byte{}, errors.New("Version too large")
	}
	p[0] = h.Version << 5
	if h.Downstream {
		p[0] |= 0x8
	}
	if h.Compressed {
		p[0] |= 0x4
	}
	if h.Type > 7 {
		return []byte{}, errors.New("Type too large")
	}
	p[0] |= h.Type
	p[1] = h.NodeID
	p[2] = h.SessionID
	p[3] = h.Frame
	p[4] = h.Fragment
	return p, nil
}

func (m Message) Marshal() ([]byte, error) {
	p := bytes.Buffer{}
	h, err := m.Header.Marshal()
	_, err = p.Write(h)
	if err != nil {
		return []byte{}, err
	}
	mess := bytes.Buffer{}
	if m.Header.Fragment == 0 {
		size := make([]byte, 2)
		size[0] = uint8(m.TransactionLength)
		size[1] = uint8(m.TransactionLength >> 8)
		mess.Write(size)
	}
	if m.Header.Compressed {
		mess.Write(m.Payload)
		c, err := enc.CompressBuffer(nil, mess.Bytes(), make([]byte, 0))
		if err != nil {
			return []byte{}, err
		}
		p.Write(c)
	} else {
		p.Write(mess.Bytes()) // TODO ensure this is working properly when not  frag 0
		p.Write(m.Payload)
	}
	return p.Bytes(), nil
}
