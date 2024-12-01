package resourse

func PercentageChange(a []int) []float32 {
	var retarray []float32
	for index := 0; index < len(a)-1; index++ {
		change := (float32(a[index])/float32(a[index+1]) - 1) * 100
		retarray = append(retarray, change)
	}
	return retarray
}
