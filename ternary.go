// Package ternary provides methods to mimic the use of a ternary operator.
package ternary

// String returns first if cond is true and second if not
func String(cond bool, first, second string) string {
	if cond {
		return first
	}
	return second
}

// Interface returns first if cond is true and second if not
func Interface(cond bool, first, second interface{}) interface{} {
	if cond {
		return first
	}
	return second
}

// Int returns first if cond is true and second if not
func Int(cond bool, first, second int) int {
	if cond {
		return first
	}
	return second
}

// Float32 returns first if cond is true and second if not
func Float32(cond bool, first, second float32) float32 {
	if cond {
		return first
	}
	return second
}
