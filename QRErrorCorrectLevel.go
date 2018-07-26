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
