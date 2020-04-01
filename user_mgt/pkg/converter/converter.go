package converter

import "strconv"

func StringToFloat64(input string) float64{
	f, err:= strconv.ParseFloat(input,64)
	if err !=nil {
		return -1
	}
	return f
}
