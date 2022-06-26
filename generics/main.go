package main

import "fmt"

func SumAllInt64s(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumAllFloat64s(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func genericFuncForSums[K comparable, V float64|int64 ](m map[K]V) V{
	var total V
	for _, a := range m{
		total += a
	}
	return total
}

func main() {
	dic1 := map[string]int64{
		"a": 1,
		"b": 2,
	}
	dic2 := map[string]float64{
		"a": 1.5,
		"b": 2.0,
	}

	fmt.Println("Sums %v %v", SumAllInt64s(dic1), SumAllFloat64s(dic2))
	fmt.Println("Sums %v %v",
		genericFuncForSums[string,int64](dic1),
		genericFuncForSums[string,float64](dic2))
}
