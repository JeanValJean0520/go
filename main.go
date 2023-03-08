// 변수, 함수, 상수, 메서드 등등을 패키지로 갖고온다.
package main

// import는 기본 패키지가 아닌 외부에 있는 다른 패키지를 포함 시키기 위해 필요하다.
import (
	"fmt"
)

// 매인 함수
func main() {
	// fmt 라는 외부 패키지의 기능을 갖고와서 사용
	fmt.Println("Hello world")
}
