package s

const (
	MessageTypeAck    = 0x0
	MessageTypeStream = 0x1
)

const Version = 0x0

const MaxFrameSize = 1499

// Actually inits as zero, c++ take note
var lastFrame uint

func CreateTransaction(pyld []byte, nodeID uint8, sessionID uint8, compress bool) ([][]byte, error) {
	var messages [][]byte
	return messages, nil
}
