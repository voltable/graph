package expressions

func IsArithmetic(i interface{}) bool {
	switch i.(type) {
	case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint:
		return true
	case float32, float64:
		return true
	case complex64, complex128:
		return true
	default:
		return false
	}
}
