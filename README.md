# s-go

## Protocol

### Header

The first byte of the header consists of a three bit version, a downstream flag, and a compressed flag, than a 3 bit type value. After that, a nodeID byte, session byte, frame counter, and fragment number are passed to complete the header. After that, the rest of length is a part of a message that is not related to the header.

If the compression bit is set, the rest of the message has received brotli compression and requires decompression.

### Message

One every fragment zero of a transaction, the first two bytes of the message is the byte length as an unsigned 16-bit integer. After that, a 8-bit checksum follows and the payload than follows.