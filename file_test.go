package stp

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

type s struct {
	Index int   `json:"index"`
	TS    int64 `json:"ts"`
}

func TestWriteFileByOverwriting(t *testing.T) {
	type args struct {
		path    string
		handler func([]byte) ([]byte, error)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{
				path: "TestWriteFileByOverwriting.json",
				handler: func(oldContentBytes []byte) ([]byte, error) {
					maxIndex, contentSlice := 0, make([]*s, 0, 8)
					if len(oldContentBytes) != 0 {
						err := json.Unmarshal(oldContentBytes, &contentSlice)
						if err != nil {
							return nil, err
						}
						NewArray(contentSlice).ForEach(func(v *s, i int) {
							if maxIndex == 0 || v.Index > maxIndex {
								maxIndex = v.Index
							}
						})
						if maxIndex > 2 {
							maxIndex, contentSlice = 0, contentSlice[:0]
						}
					}
					contentSlice = append(contentSlice, &s{Index: maxIndex + 1, TS: time.Now().Unix()})
					newContentBytes, err := json.Marshal(contentSlice)
					if err != nil {
						return nil, err
					}
					return newContentBytes, nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFileByOverwriting(tt.args.path, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("WriteFileByOverwriting() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriteFileByAppend(t *testing.T) {
	type args struct {
		path    string
		handler func([]byte) ([]byte, error)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{
				path: "TestWriteFileByAppend.json",
				handler: func(oldContentBytes []byte) ([]byte, error) {
					maxIndex, contentSlice := 0, make([]*s, 0, 8)
					if len(oldContentBytes) != 0 {
						err := json.Unmarshal(oldContentBytes, &contentSlice)
						if err != nil {
							return nil, err
						}
						NewArray(contentSlice).ForEach(func(v *s, i int) {
							if maxIndex == 0 || v.Index > maxIndex {
								maxIndex = v.Index
							}
						})
						if maxIndex > 2 {
							maxIndex, contentSlice = 0, contentSlice[:0]
						}
					}
					contentSlice = append(contentSlice, &s{Index: maxIndex + 1, TS: time.Now().Unix()})
					newContentBytes, err := json.Marshal(contentSlice)
					if err != nil {
						return nil, err
					}
					return newContentBytes, nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFileByAppend(tt.args.path, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("WriteFileByAppend() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_ReadFile_d12b1533c0e904e38b8b6ce42b62d37d(t *testing.T) {
	type args struct {
		path    string
		handler func([]byte) ([]*s, error)
	}
	tests := []struct {
		name  string
		args  args
		want0 []*s
		want1 error
	}{
		{
			name: "test case 1",
			args: args{
				path: "TestWriteFileByOverwriting.json",
				handler: func(b []byte) ([]*s, error) {
					s := make([]*s, 0, 8)
					err := json.Unmarshal(b, &s)
					return s, err
				},
			},
			want0: nil,
			want1: nil,
		},
		{
			name: "test case 2",
			args: args{
				path: "TestWriteFileByOverwriting.json",
				handler: func(b []byte) ([]*s, error) {
					s := make([]*s, 0, 8)
					err := json.Unmarshal(b, &s)
					return s, err
				},
			},
			want0: nil,
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0, got1 := ReadFile(tt.args.path, tt.args.handler)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("ReadFile() got0 = %v, want0 %v", got0, tt.want0)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReadFile() got1 = %v, want1 %v", got1, tt.want1)
			}
		})
	}
}
