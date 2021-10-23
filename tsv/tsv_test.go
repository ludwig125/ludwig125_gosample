package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReadTSV(t *testing.T) {
	tests := map[string]struct {
		data    string
		wantErr bool
		want    []Data
	}{
		"1": {
			data: "a1\tb1\na2\tb2\na3\tb3\n",
			want: []Data{
				{A: "a1", B: "b1"},
				{A: "a2", B: "b2"},
				{A: "a3", B: "b3"},
			},
		},
		"empty": {
			data:    "",
			wantErr: false,
			want:    nil,
		},
		"abnormal": {
			data:    "a1\tb1\na2",
			wantErr: true, //  record on line 2: wrong number of fields
			want:    nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ConvertDataFromTSV(strings.NewReader(tt.data))
			if err != nil {
				if !tt.wantErr {
					t.Errorf("error: %v, wantErr: %t", err, tt.wantErr)
				}
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("got: %v,want: %v, diff: %s", got, tt.want, diff)
			}
		})
	}
}
