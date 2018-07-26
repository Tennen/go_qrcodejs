package go_qrcodejs

import "errors"

type QRMode uint

const (
	number      QRMode = 1 << iota
	alphaNumber QRMode = 1 << iota
	bitByte8    QRMode = 1 << iota
	kanji       QRMode = 1 << iota
)

func (mode QRMode) bitCount(qrtype int) (int, error) {
	if 1 <= qrtype && qrtype < 10 {
		switch mode {
		case number:
			return 10, nil
		case alphaNumber:
			return 9, nil
		case bitByte8, kanji:
			return 8, nil
		}
	} else if qrtype < 27 {
		switch mode {
		case number:
			return 12, nil
		case alphaNumber:
			return 11, nil
		case bitByte8:
			return 6, nil
		case kanji:
			return 10, nil
		}
	} else if qrtype < 41 {
		switch mode {
		case number:
			return 14, nil
		case alphaNumber:
			return 13, nil
		case bitByte8:
			return 16, nil
		case kanji:
			return 12, nil
		}
	}
	return 0, errors.New("unsupported qrtype")
}
