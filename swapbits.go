func SwapBits(x byte) byte {
	return 16*(x%16) + x/16
}
