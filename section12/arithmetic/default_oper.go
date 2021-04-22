// 두 숫자의 사칙연산 계산 제공 패키지 - 1
package arithmetic

// x,y 두 개의 Integer 구조체
type Numbers struct {
	X int
	Y int
}

// x,y 합을 계산해서 반환
func (o *Numbers) Plus() int {
	return o.X + o.Y
}

// x,y 차를 계산해서 반환
func (o *Numbers) Minus() int {
	return o.X - o.Y
}

// x,y 곱을 계산해서 반환
func (o *Numbers) Mulit() int {
	return o.X * o.Y
}

// x,y 나누기해서 몫을 반환
func (o *Numbers) Divide() int {
	return o.X / o.Y
}

