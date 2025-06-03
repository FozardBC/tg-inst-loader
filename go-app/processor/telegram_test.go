package processor

import (
	"testing"
)

func Test_validateMsg(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{"https://www.instagram.com/p/CS1J9Z1J1Z1/"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateMsg(tt.args.msg); got != tt.want {
				t.Errorf("validateMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
