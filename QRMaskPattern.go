package go_qrcodejs

type QRMaskPattern int

const (
	mask0 QRMaskPattern = iota
	mask1
	mask2
	mask3
	mask4
	mask5
	mask6
	mask7
)

func (this QRMaskPattern) getMask(i int, j int) bool {
	switch this {
	case mask0:
		return (i+j)%2 == 0
	case mask1:
		return i%2 == 0
	case mask2:
		return j%3 == 0
	case mask3:
		return (i+j)%3 == 0
	case mask4:
		return (i/2+j/3)%2 == 0
	case mask5:
		return (i*j)%2+(i*j)%3 == 0
	case mask6:
		return ((i*j)%2+(i*j)%3)%2 == 0
	case mask7:
		return ((i*j)%3+(i+j)%2)%2 == 0
	}
	return false
}
