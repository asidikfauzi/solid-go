package order

type Payment interface{
	Pay(float64)
}