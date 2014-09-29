package mt

import (
	"testing"
)

func BenchmarkMt(b *testing.B) {
	mt := New(22, 0)
	for i := 0; i < b.N; i++ {
		mt.IntN(10000)
	}
}

func Equals(b1 []uint32, b2 []uint32) bool {
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}

func TestMt(t *testing.T) {
	src := New(322, 0)
	size := uint64(1000000)
	v1 := make([]uint32, size)
	v2 := make([]uint32, size)

	for i := uint64(0); i < size; i++ {
		v1[i] = src.IntN(size)
	}
	t.Log(src.Counter)
	split := uint64(500000)
	mt1 := New(322, split)
	for i := uint64(split); i < size; i++ {
		v2[i] = mt1.IntN(size)
	}
	t.Log(Equals(v1[split:], v2[split:]))
}
