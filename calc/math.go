package calc

//return sum of two integers
func Add(numbers ...int) (int,string) {
	sum :=0

	for _, num := range numbers {
		sum = sum + num
	}
	return sum, "returns 2nd param"
}