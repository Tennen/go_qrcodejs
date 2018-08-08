package go_qrcodejs

type QRModel struct {
	TypeNumber        int
	ErrorCorrectLevel QRErrorCorrectLevel
	modules           [][]bool
	moduleCount       int
	encodedText       QR8bitByte
	dataCache         []int
}

func (self *QRModel) Init(text string, typeNumber int, errorCorrectLevel QRErrorCorrectLevel) {
	bit := QR8bitByte{}
	bit.Init(text)
	encoded := bit
	self.encodedText = encoded
	self.TypeNumber = typeNumber
	self.ErrorCorrectLevel = errorCorrectLevel
}

func (self *QRModel) CreateData() []int {
	rsBlock := self.ErrorCorrectLevel.getRSBlocksOfType(self.TypeNumber)
	buffer := QRBitBuffer{}
	buffer.PutNumber(uint(self.encodedText.mode), 4)
	length, err := self.encodedText.mode.bitCount(self.TypeNumber)
	if err != nil {
		panic(err)
	}
	buffer.PutNumber(uint(self.encodedText.Count), length)
	self.encodedText.Write(&buffer)
	return []int{}
}
