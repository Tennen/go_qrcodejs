package go_qrcodejs

type QRPolynomial struct {
	numbers []int
	count   int
}

func (this *QRPolynomial) Init(nums []int, shift int) {
	var offset int
	length := len(nums)
	for offset < length && nums[offset] == 0 {
		offset += 1
	}
	this.numbers = make([]int, length-offset+shift)
	for i := 0; i < length-offset; i++ {
		this.numbers[i] = nums[i+offset]
	}
	this.count = length
}

func (this *QRPolynomial) GetFromIndex(index int) int {
	return this.numbers[index]
}

func (this *QRPolynomial) Multiplying(index int) int {