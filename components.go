package s

type Ack Header

type Message struct {
	Header              Header
	Payload             []byte
	Length              int
	TransactionLength   uint16
	TransactionChecksum uint8
}

type Header struct {
	Downstream bool
	Compressed bool
	Type       uint8
	Version    uint8
	NodeID     uint8
	SessionID  uint8
	Frame      uint8
	Fragment   uint8
}
