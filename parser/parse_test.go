package parser

import (
	"monkey-lang-clone/ast"
	"monkey-lang-clone/lexer"
	"testing"
)

func TestLetStatment(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	// parser를 생성하기 위해 사용자를 입력을 인자로 넘겨 lexer를 생성하고
	l := lexer.New(input)

	// lexer를 이용해 parser를 생성한다.
	p := New(l)

	// 프로그램을 구문 분석을 한다.
	program := p.ParseProgram()

	// 만약 구문 분석한 프로그램이 nil이라면,
	if program == nil {
		// 에러 실패 뜨게 하기
		t.Fatalf("ParseProgram() returned nil")
	}

	// 만약 Statment의 크기가 3이 아니라면
	if len(program.Statements) != 3 {
		// 에러 내기
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	// 테스트 데이터 생성
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	// 모든 테스트 셋을 하니씩 돌리며
	for i, tt := range tests {
		// 구문 분석이 완료된, Statements들을 변수에 할당
		letStmt := program.Statements[i]

		// 테스트 ㄱㄱ~
		if testLetStatement(t, letStmt, tt.expectedIdentifier) == false {
			return
		}
	}

}

// testLetStatement 함수는
func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	// TokneLiteral이 let과 다르다면,
	if s.TokenLiteral() != "let" {
		// 에러!!
		t.Errorf("s.Tokenliteral not 'let'. get=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

// Monkey 소스코드를 입력으로 제공하고 나서, 파서가 만들어냈으면 하는 AST 형태를 기댓값으로 설정한다.
