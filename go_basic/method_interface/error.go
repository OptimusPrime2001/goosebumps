package main

import (
	"fmt"
	"math"
)

// Bước 1: Tạo kiểu lỗi mới
type ErrNegativeSqrt float64

// Bước 2: Triển khai phương thức Error() cho nó
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Bước 3: Viết hàm Sqrt xử lý lỗi
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	previous := 0.0
	for range 1000 {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-previous) < 1e-10 {
			break
		}
		previous = z
	}
	return z, nil
}

func TestError() {
	fmt.Println(Sqrt(2))  // ✅ Không lỗi
	fmt.Println(Sqrt(-2)) // ❌ Báo lỗi: cannot Sqrt negative number: -2
}
