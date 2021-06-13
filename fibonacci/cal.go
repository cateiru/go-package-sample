package fibonacci

// CalType is calculation type.
// 1: 再帰
// 2: for-loop
type Calculate struct {
	CalType int
}

func FibonacciCal(caltype int) *Calculate {
	cal := new(Calculate)
	cal.CalType = caltype
	return cal
}

func (cal *Calculate) Cal(input int64) (int64, error) {
	var result int64

	switch cal.CalType {
	case 1:
		result = calRecursion(input)
	case 2:
		break
	}

	return int64(result), nil
}

func calRecursion(a int64) int64 {
	if a <= 1 {
		return a
	}
	return calRecursion(a-1) + calRecursion(a-2)
}
