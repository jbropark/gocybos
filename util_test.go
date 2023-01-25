package gocybos

import "testing"

func TestIsUserAnAdmin(t *testing.T) {
	v, err := IsUserAnAdmin()
	if err != nil {
		panic(err)
	}
	t.Logf("IsAdmin: %v", v)
}

func shift(x, n int) bool {
	return (x >> n) == 0
}

func div(x, n int) bool {
	return (x / n) == 0
}

func compare(x, n int) bool {
	return x > n
}

func BenchmarkUIntToTimeHM(b *testing.B) {
	for i := 0; i < 100_000_000; i++ {
		UIntToTimeHM(1011)
	}
}

func BenchmarkUIntToTimeHMS(b *testing.B) {
	for i := 0; i < 100_000_000; i++ {
		UIntToTimeHMS(101100)
	}
}

func BenchmarkShift(b *testing.B) {
	for i := 0; i < 1_000_000_000; i++ {
		shift(10000, 12)
	}
}

func BenchmarkDiv(b *testing.B) {
	for i := 0; i < 1_000_000_000; i++ {
		div(10000, 12)
	}
}

func BenchmarkBit(b *testing.B) {
	for i := 0; i < 1_000_000_000; i++ {
		_ = (i & 0b1111000000) == 0
	}
}

func BenchmarkCompare(b *testing.B) {
	for i := 0; i < 1_000_000_000; i++ {
		compare(10000, 10000)
	}
}
