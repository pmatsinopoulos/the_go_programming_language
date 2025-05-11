package count_diff_bits

import "testing"

func TestCountNumberOfDifferentBitsCase1(t *testing.T) {
	var a [32]byte = [32]byte{0x01, 0x02}
	var b [32]byte = [32]byte{0x01, 0x02}

	var result int = CountDifferentBits(a, b)
	if result != 0 {
		t.Error("CountDifferentBits failed")
	}
}

func TestCountNumberOfDifferentBitsCase2(t *testing.T) {
	var a [32]byte = [32]byte{0x01, 0x02}
	var b [32]byte = [32]byte{0x01, 0x01}

	var result int = CountDifferentBits(a, b)
	if result != 2 {
		t.Error("CountDifferentBits failed")
	}
}

func TestCountNumberOfDifferentBitsCase3(t *testing.T) {
	var a [32]byte = [32]byte{0x01, 0x02, 0x11}
	var b [32]byte = [32]byte{0x01, 0x01, 0x38}

	var result int = CountDifferentBits(a, b)
	if result != 5 {
		t.Error("CountDifferentBits failed")
	}
}

func TestCountNumberOfDifferentBitsCase4(t *testing.T) {
	var a [32]byte = [32]byte{0x01, 0x02, 0x11, 0xFa, 0xeb}
	var b [32]byte = [32]byte{0x01, 0x01, 0x38, 0xFa, 0x1e}

	var result int = CountDifferentBits(a, b)
	var expected int = 11
	if result != expected {
		t.Errorf("CountDifferentBits failed: got %d, want %d", result, expected)
	}
}
