// 이 파일이 프로그램의 시작점(Entry Point) 임을 선언하는 코드
// 빌드 시, "실행파일" 이 생성된다.
package main

import "fmt"

// --------------[struct]--------------

// 구조체 정의 (데이터 설계)
// 특징: 객체지향의 '클래스' 역할을 수행하지만, 상속은 없음
type Player struct {
	Name  string
	Level int
}

// 메서드 정의 (기능 추가)
// (p Player)는 이 함수를 Player 구조체에 귀속시킨다는 뜻 (Receiver)
func (p Player) Info() {
	fmt.Printf("ID: %s | LV: %d\n", p.Name, p.Level)
}

// [복사] (메서드) 원본은 그대로, 가짜(복제본)만 수정됨
func (p Player) MethodLevelup() {
	p.Level++
}

// [참조] (메서드) * 를 붙여서 메모리 주소를 받음. 원본 수정!
func (p *Player) MethodRealLevelup() {
	p.Level++
}

// [복사] (함수) 원본은 그대로, 가짜(복제본)만 수정됨
func FunctionLevelup(p Player) {
	p.Level++
}

// [참조] (함수) * 를 붙여서 메모리 주소를 받음. 원본 수정!
func FunctionRealLevelup(p *Player) {
	p.Level++
}

// --------------[interface]--------------

// 인터페이스 정의
type Sounder interface {
	MakeSound()
}

type Mover interface {
	Move()
}

// 구조체 정의
type Dog struct{}
type Robot struct{}

// 각자 방식대로 메서드 구현
func (d Dog) MakeSound() {
	fmt.Println("멍")
}

func (r Robot) MakeSound() {
	fmt.Println("위잉")
}

func (r Robot) Move() {
	fmt.Println("로봇 이동 중..")
}

func DoSound(s Sounder) {
	s.MakeSound() // 어떤 타입이 오든 Sounder 인터페이스만 만족하면 OK!
}

func main() {
	fmt.Println("\n--------------[struct 실습]--------------")

	// 구조체 인스턴스 생성
	p1 := Player{Name: "GoMaster", Level: 1}

	// 메서드 호출
	p1.Info()

	p1.MethodLevelup()
	fmt.Println(p1.Level) // 결과: 1 (안 변함)

	p1.MethodRealLevelup()
	fmt.Println(p1.Level) // 결과: 2 (변함)

	fmt.Println("\n--------------[interface 실습]--------------")
	d := Dog{}
	r := Robot{}

	DoSound(d) // 결과: 멍
	DoSound(r) // 결과: 위잉

	fmt.Println("")

	var s Sounder = r
	var m Mover = r

	s.MakeSound() // 결과: 위잉
	m.Move()      // 결과: 로봇 이동 중..
}
