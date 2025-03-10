package twocrystalballs

import "math"

func TwoCrystalBalls(breaks []bool) int {
	length := len(breaks)
	jmpAmount := int(math.Floor(math.Sqrt(float64(length))))

	i := jmpAmount
	for ; i < length; i += jmpAmount {
		if breaks[i] {
			break
		}
	}
	i -= jmpAmount

	for j := 0; j < jmpAmount && i < length; j, i = j+1, i+1 {
		if breaks[i] {
			return i
		}
	}

	return -1

}
