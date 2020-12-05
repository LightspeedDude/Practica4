package code

func cBtD(n int) int {
	r := 0
	d := 0
	c := 1
	for n != 0 {
		r = n % 10
		n /= 10
		d += r * c
		c *= 2
	}
	return d
}

func cDtO(n int) int {
	r := 0
	o := 0
	c := 1
	for n != 0 {
		r = n % 8
		n /= 8
		o += r * c
		c *= 10
	}
	return o
}
