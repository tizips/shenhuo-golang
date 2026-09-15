package helper

const DefaultOrder uint8 = 50

func Order(n uint8) uint8 {
	if n == 0 {
		return DefaultOrder
	}
	return n
}
