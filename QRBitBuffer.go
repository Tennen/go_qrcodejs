package go_qrcodejs

type QRBitBuffer struct {
	buffer   []uint
	bitcount int
}

func (this *QRBitBuffer) GetFromIndex(index uint) bool {
	bufferIndex := index / 8
	return ((this.buffer[bufferIndex] >> (7 - index%8)) & 1) == 1
}

func (this *QRBitBuffer) Put(bit bool) {
	bufIndex := this.bitcount / 8
	if len(this.buffer) <= bufIndex {
		this.buffer = append(this.buffer, 0)
	}
	if bit {
		this.buffer[bufIndex] |= (uint(0x80) >> uint(this.bitcount%8))
	}
	this.bitcount += 1
}

func (this *QRBitBuffer) PutNumber(num uint, length int) {
	for i := 0; i < length; i++ {
		this.Put(((num >> uint(length-i-1)) & 1) == 1)
	}
}
