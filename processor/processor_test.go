package processor

import (
	"reflect"
	"testing"
)

var p Processor = Processor{}

func Test_normalizeURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				url: "https://www.instagram.com/reel/DEnG0T_N8YS/?igsh=ZWE0bjM3a3p3aW82",
			},
			want: "https://www.instagram.com/p/DEnG0T_N8YS/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := normalizeURL(tt.args.url); got != tt.want || err != nil {
				t.Errorf("normalizeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessor_loadContent(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		p       *Processor
		args    args
		wantErr bool
	}{
		{
			name:    "Test 1",
			p:       &p,
			args:    args{"https://www.instagram.com/reel/DEnG0T_N8YS/?igsh=ZWE0bjM3a3p3aW82"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.LoadContent(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("Processor.loadContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_findFiles(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantU   []string
		wantErr bool
	}{
		{
			name:    "Test 1",
			args:    args{"C:/src/tgInstaagramLoad/cmd/downloads"},
			wantU:   []string{"cmd/downloads/2025-01-09_16-10-42_UTC.mp4"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotU, err := findFiles(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("findFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotU, tt.wantU) {
				t.Errorf("findFiles() = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}

func TestProcessor_LoadContent(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		p       *Processor
		args    args
		wantErr bool
	}{
		{
			name:    "Test 1",
			p:       &p,
			args:    args{"https://www.instagram.com/reel/DEnG0T_N8YS/?igsh=ZWE0bjM3a3p3aW82"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.LoadContent(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("Processor.LoadContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
