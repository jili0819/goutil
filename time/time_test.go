package time

import (
	"reflect"
	"testing"
	"time"
)

func TestFormatDefaultTime(t *testing.T) {
	type args struct {
		t      time.Time
		format string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test01", args: args{
			format: DateFormat,
			t:      time.Now(),
		},
			want: "2021-12-21"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDefaultTime(tt.args.t, tt.args.format); got != tt.want {
				t.Errorf("FormatDefaultTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimestampToTime(t *testing.T) {
	type args struct {
		timestamp int64
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
		{name: "test01",
			args: args{
				timestamp: 1640073672,
			}, want: time.Unix(1640073672, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimestampToTime(tt.args.timestamp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimestampToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimestampToTimeStr(t *testing.T) {
	type args struct {
		timestamp int64
		format    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test01",
			args: args{
				timestamp: 1640073672,
				format:    DateFormat,
			}, want: "2021-12-21"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimestampToTimeStr(tt.args.timestamp, tt.args.format); got != tt.want {
				t.Errorf("TimestampToTimeStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
