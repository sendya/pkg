package traceid

import "testing"

func Test_newID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newID(); got != tt.want {
				t.Errorf("newID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNewID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newID()
	}
}
