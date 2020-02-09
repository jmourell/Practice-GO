package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
	assertCorrect := func(t *testing.T, got, want string) {
	t.Helper()
    if got != want {
			t.Errorf("got %q want	%q", got , want)
		}
	}
	t.Run("repeating 5 times on default", func(t *testing.T){
		got := Repeat("a", 0)
		expected := "aaaaa"
		assertCorrect(t,got,expected)
	})
	t.Run("function should do the number specified", func(t *testing.T){
		got := Repeat("b", 3)
		expected := "bbb"
		assertCorrect(t,got,expected)
	})
}

func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}

func ExampleRepeat() {
	
	repeate := Repeat("d", 8)
	fmt.Println(repeate)
	//Output: dddddddd
}