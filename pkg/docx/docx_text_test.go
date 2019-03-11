package docx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocxToText(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   map[string]string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "002",
			args:    args{fileName: "../../test/data/002.docx"},
			want:    " ",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := DocxToText(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("DocxToText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.True(t, len(got) > 0)
			/*
				if got != tt.want {
					t.Errorf("DocxToText() got = %v, want %v", got, tt.want)
				}
			*/
			/*
				if !reflect.DeepEqual(got1, tt.want1) {
					t.Errorf("DocxToText() got1 = %v, want %v", got1, tt.want1)
				}
			*/
		})
	}
}
