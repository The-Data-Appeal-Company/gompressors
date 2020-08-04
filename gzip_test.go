package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestGzipCompressor_Compress(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "shouldCompressMessage",
			args: args{
				data: readFileOrError(t, "test_data/input.json"),
			},
			want:    readFileOrError(t, "test_data/input.json"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GzipCompressor{}
			got, err := g.Compress(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			decompress, err := g.Decompress(got)
			assert.NoError(t, err)
			if !reflect.DeepEqual(decompress, tt.want) {
				t.Errorf("Compress() got = %v, want %v", decompress, tt.want)
			}
		})
	}
}

func TestGzipCompressor_Decompress(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "shouldDecompressGzippedMessage",
			args: args{
				input: readFileOrError(t, "test_data/input.json.gz"),
			},
			want:    readFileOrError(t, "test_data/input.json"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GzipCompressor{}
			got, err := g.Decompress(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decompress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func readFileOrError(t *testing.T, fileName string) []byte {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
	return data
}
