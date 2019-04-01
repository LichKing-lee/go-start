// 한 주의 요일에 대한 상수 생성자 iota를 작성한다.
package main

import (
	"fmt"
	"time"
)

func main() {
	const(
		SUN time.Weekday = iota
		MON
		TUE
		WED
		THU
		FRI
		SAT
	)

	fmt.Println(SUN)
	fmt.Println(MON)
	fmt.Println(TUE)
	fmt.Println(WED)
	fmt.Println(THU)
	fmt.Println(FRI)
	fmt.Println(SAT)
	fmt.Println(MON)
}
