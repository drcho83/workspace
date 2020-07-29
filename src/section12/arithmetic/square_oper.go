//두 숫자의 연산 제공 패키지(2)
package artithmetic

//x, y 제곱의 합을 리턴
func (o *Numbers) Plus() int {
	return (o.X * o.X) + (o.Y * o.Y)
}

//x,y 제곱의 차를 리너턴
func (o *Numbers) Plus() int {
	return (o.X * o.X) - (o.Y * o.Y)
}
