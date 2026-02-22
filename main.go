// 이 파일이 프로그램의 시작점(Entry Point) 임을 선언하는 코드
// 빌드 시, "실행파일" 이 생성된다.
package main

import "fmt"

// 1. 구조체 정의 (데이터 설계)
// 특징: 객체지향의 '클래스' 역할을 수행하지만, 상속은 없음
type Player struct {
	Name  string
	Level int
}

// 2. 메서드 정의 (기능 추가)
// (p Player)는 이 함수를 Player 구조체에 귀속시킨다는 뜻 (Receiver)
func (p Player) Info() {
	fmt.Printf("ID: %s | LV: %d\n", p.Name, p.Level)
}

func main() {
	// 3. 구조체 인스턴스 생성
	p1 := Player{Name: "GoMaster", Level: 1}

	// 4. 메서드 호출
	p1.Info()
}
