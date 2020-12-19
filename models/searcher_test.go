package models

import (
	"index/suffixarray"
	"testing"
)

func TestSearcher_Load(t *testing.T) {
	type fields struct {
		CompleteWorks string
		SuffixArray   *suffixarray.Index
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid filename",
			fields: fields{},
			args: args{
				filename: "../completeworks.txt",
			},
			wantErr: false,
		},
		{
			name: "invalid filename",
			fields: fields{},
			args: args{
				filename: "../completeworks.md",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Searcher{
				CompleteWorks: tt.fields.CompleteWorks,
				SuffixArray:   tt.fields.SuffixArray,
			}
			if err := s.Load(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}