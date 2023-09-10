package datetime

import (
	"reflect"
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    time.Time
		wantErr bool
	}{
		{
			name:    "invalid date",
			s:       "invalid123",
			wantErr: true,
		},
		{
			name: "invalid date",
			s:    "2023-09-09",
			want: time.Date(2023, 9, 9, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDate(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
