package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestA(x *testing.T) {
	result := HelloWorld("Go")
	if result != "Hello Go" {
		// panic("Result is not 'Hello Go'")
		// x.FailNow()
		x.Fatal("Result must be:", "Hello Go")
	}
	fmt.Println("TestA Done")
}

func TestA2(x *testing.T) {
	result := HelloWorld("Rust")
	if result != "Hello Rust" {
		// panic("Result is not 'Hello Rust'")
		// x.Fail()
		x.Error("Result must be:", "Hello Rust")
	}
	fmt.Println("TestA2 Done")
}

func TestHelloWorldAssert(t *testing.T)  {
	result := HelloWorld("React")
	assert.Equal(t, "Hello React", result, "Expected value should be: %s", result)
	fmt.Println("TestHelloWorldAssert done")
}

func TestHelloWorldRequire(t *testing.T)  {
	result := HelloWorld("React")
	require.Equal(t, "Hello React", result, "Expected value should be: %s", result)
	fmt.Println("TestHelloWorldRequire done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on windows")
	}
	fmt.Println("TestSkip done")
}

func TestMain(m *testing.M)  {
	fmt.Println("--------------------Before Unit Test--------------------")
	
	m.Run()
	
	fmt.Println("--------------------After Unit Test--------------------")
}

func TestSubTest(t *testing.T) {
	t.Run("Sub Test 1", func(t *testing.T) {
		result := HelloWorld("React")
		require.Equal(t, "Hello React", result, "Expected value should be: %s", result)
		t.Run("Sub Test 1 - Sub Test 1", func(t *testing.T) {
			result := HelloWorld("React")
			require.Equal(t, "Hello React", result, "Expected value should be: %s", result)
		})
	})
	t.Run("Sub Test 2", func(t *testing.T) {
		result := HelloWorld("React")
		assert.Equal(t, "Hello React", result, "Expected value should be: %s", result)
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct{
		name string
		expected string
		request string
	}{
		{
			name: "Sub Test React",
			expected: "Hello React",
			request: "React",
		},
		{
			name: "Sub Test Vue",
			expected: "Hello Vue",
			request: "Vue",
		},
		{
			name: "Sub Test Svelte",
			expected: "Hello Svelte",
			request: "Svelte",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, HelloWorld(test.request))
		})
	}
}

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Go")
	}
}

func BenchmarkA2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rust")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("BenchMarkSub 1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Go")
		}
	})
	b.Run("BenchMarkSub 2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rust")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct{
		name, request string
	}{
		{"Go", "Go"},
		{"Rust", "Rust"},
		{"JavaScript", "JavaScript"},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}