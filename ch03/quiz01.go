// 숫자 4의 제곱에 대한 상수 생성자 iota를 작성한다
package main

import "fmt"

func main() {
	const(
		TEST1 = 1 << iota
		_
		TEST2
		_
		TEST3
		_
		TEST4
		_
		TEST5
	)

	fmt.Println(TEST1)
	fmt.Println(TEST2)
	fmt.Println(TEST3)
	fmt.Println(TEST4)
	fmt.Println(TEST5)
}
