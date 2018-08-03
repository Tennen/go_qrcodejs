package go_qrcodejs

type QRMath struct {
	EXP_TABLE [256]int
	LOG_TABLE [256]int
}

var qrmath *QRMath

func (this *QRMath) Init() {
	for i := 0; i < 8; i++ {
		this.EXP_TABLE[i] = 1 << uint(i)
	}
	for i := 8; i < 256; i++ {
		this.EXP_TABLE[i] = this.EXP_TABLE[i-4] ^ this.EXP_TABLE[i-5] ^ this.EXP_TABLE[i-6] ^ this.EXP_TABLE[i-8]
	}
	for i := 0; i < 255; i++ {
		this.LOG_TABLE[this.EXP_TABLE[i]] = i
	}
}

func (this *QRMath) Glog(n int) int {
	return this.LOG_TABLE[n]
}

func (this *QRMath) Gexp(n int) int {
	for n < 0 {
		n += 255
	}
	for n >= 256 {
		n -= 255
	}
	return this.EXP_TABLE[n]
}
func getMath() *QRMath {
	if qrmath == nil {
		qrmath = &QRMath{}
		qrmath.Init()
	}
	return qrmath
}
