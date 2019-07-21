package tiki

import (
	"path/filepath"
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestReadCSVFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name     string
		args     args
		wantData [][]string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Test",
			args: args{
				path: func(relativeFileFromCurrentFolder string) string {
					p, _ := filepath.Abs(relativeFileFromCurrentFolder)
					log.Infof(p)
					return p
				}("./datatest/data_test.csv"),
			},
			wantData: [][]string{
				[]string{"0987000001", "2016-03-01", "2016-05-01"},
				[]string{"0987000002", "2016-02-01", "2016-03-01"},
				[]string{"0987000003", "2016-01-01", "2016-03-01"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := ReadCSVFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCSVFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("ReadCSVFile() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestWriteCSVFile(t *testing.T) {
	type args struct {
		path string
		data [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test",
			args: args{
				path: func(relativeFileFromCurrentFolder string) string {
					p, _ := filepath.Abs(relativeFileFromCurrentFolder)
					log.Infof(p)
					return p
				}("./datatest/data_result_test.csv"),
				data: [][]string{
					[]string{"1", "2", "3"},
					[]string{"4", "5", "6"},
					[]string{"7", "8", "9"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteCSVFile(tt.args.path, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("WriteCSVFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
