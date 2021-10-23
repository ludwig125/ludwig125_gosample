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
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ConvertDataFromTSV(strings.NewReader(tt.data))
			if (err != nil) != tt.wantErr {
				t.Errorf("error: %v, wantErr: %t", err, tt.wantErr)
				return
			}
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("got: %v,want: %v, diff: %s", got, tt.want, diff)
			}
			// ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
			// infos, err := errSem(ctx, tt.cities)
			// if err != nil {
			// 	t.Fatal(err)
			// }
			// for _, v := range infos {
			// 	fmt.Println(*v)
			// }

		})
	}
}
