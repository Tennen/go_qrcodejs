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

func (this *QRPolynomial) Multiplying(e *QRPolynomial) *QRPolynomial {
	length := this.count + e.count - 1
	nums := make([]int, length)
	for i := 0; i < this.count; i++ {
		for j := 0; j < e.count; j++ {
			math := getMath()
			nums[i+j] = math.Gexp(math.Glog(this.GetFromIndex(i)) + math.Glog(e.GetFromIndex(j)))
		}
	}
	res := QRPolynomial{}
	res.Init(nums, 0)
	return &res
}

func (this *QRPolynomial) Moded(e *QRPolynomial) *QRPolynomial {
	if this.count < e.count {
		return this
	}
	math := getMath()
	ratio := math.Glog(this.GetFromIndex(0)) - math.Glog(e.GetFromIndex(0))
	num := make([]int, this.count)
	for i := 0; i < this.count; i++ {
		num[i] = this.GetFromIndex(i)
	}
	for i := 0; i < e.count; i++ {
		num[i] ^= math.Gexp(math.Glog(e.GetFromIndex(i) + ratio))
	}
	res := &QRPolynomial{}
	res.Init(num, 0)
	res = res.Moded(e)
	return res
}

func (this *QRPolynomial) ErrorCorrectPolynomial(errorCorrectLength int) *QRPolynomial {
	a := &QRPolynomial{}
	a.Init([]int{1}, 0)
	math := getMath()
	for i := 0; i < errorCorrectLength; i++ {
		temp := &QRPolynomial{}
		temp.Init([]int{1}, math.Gexp(i))
		a = a.Multiplying(temp)
	}
	return a
}
