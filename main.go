// 변수, 함수, 상수, 메서드 등등을 패키지로 갖고온다.
package main

// import는 기본 패키지가 아닌 외부에 있는 다른 패키지를 포함 시키기 위해 필요하다.
import (
	"fmt"
	"log"
	"os"
	"time"
)

func getTime() time.Time {
	nowTime := time.Now()
	return nowTime
}

func print(data string) {
	fmt.Println(data)
}

func repeat(repeat int) {
	for index := 0; index < repeat; index++ {
		fmt.Println(index)
	}
}

func main() {
	file, err := os.Open("./README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(file.Name())
}
