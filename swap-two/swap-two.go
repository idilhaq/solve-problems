package swaptwo

func SwapTwo(x, y int) (newX, newY int) {
	// x = 1, y = 2
	x = x + y // x = 1 + 2 = 3
	y = x - y // y = 3 - 2 = 1
	x = x - y // x = 3 - 1 = 2

	return x, y // 1, 2
}

func SwapTwoWithGolang(x, y int) (newX, newY int) {
	x, y = y, x

	return x, y
}

func SwapTwoWithBitwiseXOR(x, y int) (newX, newY int) {
	x = x ^ y
	y = x ^ y
	x = x ^ y

	return x, y
}
