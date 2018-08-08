package go_qrcodejs

type QR8bitByte struct {
	mode       QRMode
	data       string
	parsedData []byte
	Count      int
}

func (this *QR8bitByte) Init(data string) {
	this.data = data
	this.parsedData = []byte(data)
	this.Count = len(this.parsedData)
}

func (this *QR8bitByte) Write(buffer *QRBitBuffer) {
	for _, val := range this.parsedData {
		buffer.PutNumber(uint(val), 8)
	}
}
