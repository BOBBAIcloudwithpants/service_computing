package hw2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPackInFields(t *testing.T) {
	output := PackInFields("i love you my friend           carmen         鲍勃")
	expected := []string{"i", "love", "you", "my", "friend", "carmen", "鲍勃"}

	if !reflect.DeepEqual(output, expected) {
		t.Errorf("expected '%q' but got '%q'", expected, output)
	}
}

func BenchmarkPackInFields(b *testing.B) {
	for i := 0; i<b.N ;i++ {
		PackInFields("i love you my friend           carmen         鲍勃")
	}
}

func ExamplePackInFields() {
	outcome := PackInFields("i love you my friend           ")
	fmt.Printf("'%q'\n", outcome)
	// Output: '["i" "love" "you" "my" "friend"]'
}

func TestSwitchToTitle(t *testing.T) {
	output := SwitchToTitle("play with me")
	expected := "Play With Me"

	if output != expected {
		t.Errorf("expected '%q' but got '%q'", expected, output)
	}
}

func BenchmarkSwitchToTitle(b *testing.B) {
	for i := 0;i < b.N; i++ {
		SwitchToTitle("adf asdfqewr qwer df asdf ewrq dsf qwerq 1adfa q ewqrewfsda sadf a")
	}
}

func ExampleSwitchToTitle() {
	outcome := SwitchToTitle("paul maccartney george harrison john lennon ringo")
	fmt.Println(outcome)
	// Output: Paul Maccartney George Harrison John Lennon Ringo
}
