package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// benchmark table
func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Akbar)",
			request: "Akbar",
		},
		{
			name:    "HelloWorld(Syidiq)",
			request: "Syidiq",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.name)
			}
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("akbar")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "Joni",
			request:  "Joni",
			expected: "Hello Joni",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

//test pertama yang dijalankan
func TestMain(m *testing.M) {
	fmt.Println("Before test")
	m.Run()
	fmt.Println("After test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run in Mac OS")
	}
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("akbar")
	assert.Equal(t, "Hello akbar", result)
	fmt.Println("Eksekusi selesai")
}
