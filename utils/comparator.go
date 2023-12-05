package utils

import "time"

// Comparator will make type assertion (see IntComparator for example),
// which will panic if a or b are not of the asserted type.
//
// Should return a number:
//
//	negative , if a < b
//	zero     , if a == b
//	positive , if a > b
type Comparator[T any] func(a, b T) int

func StringComparator(a, b string) int {
	min := len(b)
	if len(a) < len(b) {
		min = len(a)
	}
	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(a[i]) - int(b[i])
	}
	if diff == 0 {
		diff = len(a) - len(b)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}

// IntComparator provides a basic comparison on int
func IntComparator(a, b int) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Int8Comparator provides a basic comparison on int8
func Int8Comparator(a, b int8) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Int16Comparator provides a basic comparison on int16
func Int16Comparator(a, b int16) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Int32Comparator provides a basic comparison on int32
func Int32Comparator(a, b int32) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Int64Comparator provides a basic comparison on int64
func Int64Comparator(a, b int64) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// UIntComparator provides a basic comparison on uint
func UIntComparator(a, b uint) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// UInt8Comparator provides a basic comparison on uint8
func UInt8Comparator(a, b uint8) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// UInt16Comparator provides a basic comparison on uint16
func UInt16Comparator(a, b uint16) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// UInt32Comparator provides a basic comparison on uint32
func UInt32Comparator(a, b uint32) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// UInt64Comparator provides a basic comparison on uint64
func UInt64Comparator(a, b uint64) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Float32Comparator provides a basic comparison on float32
func Float32Comparator(a, b float32) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Float64Comparator provides a basic comparison on float64
func Float64Comparator(a, b float64) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// ByteComparator provides a basic comparison on byte
func ByteComparator(a, b byte) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// RuneComparator provides a basic comparison on rune
func RuneComparator(a, b rune) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// TimeComparator provides a basic comparison on time.Time
func TimeComparator(a, b time.Time) int {
	switch {
	case a.After(b):
		return 1
	case a.Before(b):
		return -1
	default:
		return 0
	}
}
