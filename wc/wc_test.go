package wc

import (
	"os"
	"testing"
)

var testContent = getTestFileContent()

func getTestFileContent() []byte {
	testContent, err := os.ReadFile("./test.txt")
	if err != nil {
		panic(err)
	}
	return testContent
}

func TestCountBytes(t *testing.T) {
	testContent := getTestFileContent()
	tests := []struct {
		Name string
		Arg  []byte
		Want int
	}{
		{
			Name: "Count Bytes Test File",
			Arg:  testContent,
			Want: 342190,
		},
	}

	for _, test := range tests {
		if CountBytes(test.Arg) != test.Want {
			t.Errorf("Test %s failed, got %v but wanted %v", test.Name, CountBytes(test.Arg), test.Want)
		}
	}
}

func TestCountLines(t *testing.T) {
	testContent := getTestFileContent()
	tests := []struct {
		Name string
		Arg  []byte
		Want int
	}{
		{
			Name: "Count Lines Test File",
			Arg:  testContent,
			Want: 7145,
		},
	}

	for _, test := range tests {
		if CountLines(test.Arg) != test.Want {
			t.Errorf("Test %s failed, got %v but wanted %v", test.Name, CountLines(test.Arg), test.Want)
		}
	}
}

func TestCountWords(t *testing.T) {
	testContent := getTestFileContent()
	tests := []struct {
		Name string
		Arg  []byte
		Want int
	}{
		{
			Name: "Count Words Test File",
			Arg:  testContent,
			Want: 58164,
		},
	}

	for _, test := range tests {
		if CountWords(test.Arg) != test.Want {
			t.Errorf("Test %s failed, got %v but wanted %v", test.Name, CountWords(test.Arg), test.Want)
		}
	}
}

var sink int

func BenchmarkCountLines(t *testing.B) {
	var result int
	for i := 0; i < t.N; i++ {
		result = CountLines(testContent)
	}
	sink = result
}

func BenchmarkCountWords(t *testing.B) {
	var result int
	for i := 0; i < t.N; i++ {
		result = CountWords(testContent)
	}
	sink = result
}
