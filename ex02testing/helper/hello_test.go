package helper

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Before testing")
	m.Run()
	log.Println("After testing")
}

func BenchmarkHello(b *testing.B) {
	b.Run("hello Eko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Eko")
		}
	})

	b.Run("hello Eko hello Kurniawan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Eko")
			Hello("Kurniawan")
		}
	})

	b.Run("table driven benchmark", func(b *testing.B) {
		tests := []struct {
			name    string
			request string
		}{
			{
				name:    "hello Eko",
				request: "Eko",
			},
			{
				name:    "hello Kurniawan",
				request: "Kurniawan",
			},
			{
				name:    "hello Eko Kurniawan Khannedy",
				request: "Eko Kurniawan Khannedy",
			},
			{
				name:    "hello Budi Nugraha",
				request: "Budi Nugraha",
			},
		}

		for _, test := range tests {
			b.Run(test.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					Hello(test.request)
				}
			})
		}
	})
}

func TestHello(t *testing.T) {
	t.Log(t.Name(), "start...")

	t.Run("table driven test", func(t *testing.T) {
		tests := []struct {
			name     string
			expected string
		}{
			{
				name:     "Budi",
				expected: "Hello, Budi",
			},
			{
				name:     "Joko",
				expected: "Hello, Joko",
			},
			{
				name:     "Rully",
				expected: "Hello, Rully",
			},
		}

		for _, test := range tests {
			t.Run("test hello "+test.name, func(t *testing.T) {
				result := Hello(test.name)
				assert.Equal(t, test.expected, result)
				t.Log(t.Name(), "done") // assert calls Error, therefore this line will be printed if test fails
			})
		}
	})

	t.Run("test hello Eko", func(t *testing.T) {
		result := Hello("Eko")
		assert.Equal(t, "Hello, Eko", result)
		t.Log(t.Name(), "done") // assert calls Error, therefore this line will be printed if test fails
	})

	t.Run("test hello Khannedy", func(t *testing.T) {
		result := Hello("Khannedy")
		require.Equal(t, "Hello, Khannedy", result)
		t.Log(t.Name(), "done") // require calls ErrorNow, therefore this line will not be printed if test fails
	})

	t.Log(t.Name(), "done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Cannot run on Linux")
	}
	t.Log(t.Name(), "done")
}
