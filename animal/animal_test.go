package animal

import (
	"testing"
)

func Test_sleeping(t *testing.T) {
	c := NewCat("Sinmigo", "white", 20)
	c.Sleeping()
}

func Test_eating(t *testing.T) {
	c := NewCat("Sinmigo", "white", 20)
	c.Eating()
}

func BenchmarkBigLen(b *testing.B) {
	//c := NewCat("Sinmigo", "white", 20)

	for i := 0; i < b.N; i++ {
		//c.Sleeping()
	}
}
