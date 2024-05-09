package stp

import (
	"encoding/json"
	"testing"
	"time"
)

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
					type s struct {
						Index int   `json:"index"`
						TS    int64 `json:"ts"`
					}
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
					type s struct {
						Index int   `json:"index"`
						TS    int64 `json:"ts"`
					}
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
