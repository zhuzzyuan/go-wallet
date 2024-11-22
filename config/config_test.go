package config

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "World", want: "Hello, World"},
	}
	for _, tt := range tests {
		if got := Hello(tt.name); got != tt.want {
			t.Errorf("Hello() = %v, want %v", got, tt.want)
		}
	}
}
