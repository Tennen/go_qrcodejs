package go_qrcodejs

type QRErrorCorrectLevel int

const (
	// Error resilience level: 15%
	M = iota
	// Error resilience level: 7%
	L
	// Error resilience level: 30%
	H
	// Error resilience level: 25%
	Q
)

func (this QRErrorCorrectLevel) getRsBlockTableOfType(typeNumber int) []int {
	switch this {
	case L:
		return RS_BLOCK_TABLE[(typeNumber-1)*4+0]
	case M:
		return RS_BLOCK_TABLE[(typeNumber-1)*4+1]
	case Q:
		return RS_BLOCK_TABLE[(typeNumber-1)*4+2]
	case H:
		return RS_BLOCK_TABLE[(typeNumber-1)*4+3]
	}
	return []int{}
}

func (this QRErrorCorrectLevel) getRSBlocksOfType(typeNumber int) (list []QRRSBlock) {
	rsBlock := this.getRsBlockTableOfType(typeNumber)
	length := len(rsBlock) / 3
	for i := 0; i < length; i++ {
		count := rsBlock[i*3+0]
		totalCount := rsBlock[i*3+1]
		dataCount := rsBlock[i*3+2]
		for j := 0; j < count; j++ {
			list = append(list, QRRSBlock{totalCount: totalCount, dataCount: dataCount})
		}
	}
	return
}
