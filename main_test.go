package main

import (
	"testing"
)

func Test_scrap(t *testing.T) {
	type args struct {
		browser string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "chromeSuccess",
			args: args{
				browser: "chrome",
			},
			wantErr: false,
		},
		{
			name: "firefoxSuccess",
			args: args{
				browser: "firefox",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUa, err := scrap(tt.args.browser)
			if (err != nil) != tt.wantErr {
				t.Errorf("scrap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotUa) <= 0 {
				t.Errorf("scrap() = total %v", len(gotUa))
			}
		})
	}
}
