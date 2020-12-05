package binario

func convertBinarioaDecimal(n int) int {
	remainder := 0
	decimal := 0
	counter := 1
	for n != 0 {
		remainder = n % 10
		n = n / 10
		decimal = decimal + remainder*counter
		counter = counter * 2
	}
	return decimal
}

func convertDecimalaOctal(n int) int {
	remainder := 0
	octal := 0
	counter := 1
	for n != 0 {
		remainder = n % 8
		n = n / 8
		octal = octal + remainder*counter
		counter = counter * 10
	}
	return octal
}
