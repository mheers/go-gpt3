package gpt3

// IntPtr converts an integer to an *int as a convenience
func IntPtr(i int) *int {
	return &i
}

// FloatPtr converts an float32 to an *float32 as a convenience
func FloatPtr(f float32) *float32 {
	return &f
}
