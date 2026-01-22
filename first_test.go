// 이 파일은 main 패키지에 속함. Go는 패키지 없는 코드를 허용 안 함
package main

// testing 패키지를 사용하겠다고 선언.testing.T 타입을 알려줌
import "testing"

// 테스트 함수 규칙: Test로 시작, *test.T 매개 변수, *testing.T 매개변수
func TestMyGrammar(t *testing.T) {
	println("Hello, Go!")
	t.Log("Hello, Go!") // t를 활용한 테스트
}
