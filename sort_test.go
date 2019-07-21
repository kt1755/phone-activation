package tiki

import (
	"reflect"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		data []PhoneUsage
	}
	tests := []struct {
		name string
		args args
		want []PhoneUsage
	}{
		// TODO: Add test cases.
		{
			name: "Test ",
			args: args{
				data: []PhoneUsage{
					PhoneUsage{
						Phone: "0978000001",
						Duration: Usage{
							ActivationDate: time.Date(2019, time.July, 21, 0, 35, 0, 0, time.UTC),
						},
					},
					PhoneUsage{
						Phone: "0978000001",
						Duration: Usage{
							ActivationDate: time.Date(2019, time.July, 20, 0, 35, 0, 0, time.UTC),
						},
					},
					PhoneUsage{
						Phone: "0978000001",
						Duration: Usage{
							ActivationDate: time.Date(2019, time.July, 22, 0, 35, 0, 0, time.UTC),
						},
					},
					PhoneUsage{
						Phone: "0978000001",
						Duration: Usage{
							ActivationDate: time.Date(2019, time.July, 18, 0, 35, 0, 0, time.UTC),
						},
					},
					PhoneUsage{
						Phone: "0978000001",
						Duration: Usage{
							ActivationDate: time.Date(2019, time.July, 20, 1, 35, 0, 0, time.UTC),
						},
					},
				},
			},
			want: []PhoneUsage{
				PhoneUsage{
					Phone: "0978000001",
					Duration: Usage{
						ActivationDate: time.Date(2019, time.July, 18, 0, 35, 0, 0, time.UTC),
					},
				},
				PhoneUsage{
					Phone: "0978000001",
					Duration: Usage{
						ActivationDate: time.Date(2019, time.July, 20, 0, 35, 0, 0, time.UTC),
					},
				},
				PhoneUsage{
					Phone: "0978000001",
					Duration: Usage{
						ActivationDate: time.Date(2019, time.July, 20, 1, 35, 0, 0, time.UTC),
					},
				},
				PhoneUsage{
					Phone: "0978000001",
					Duration: Usage{
						ActivationDate: time.Date(2019, time.July, 21, 0, 35, 0, 0, time.UTC),
					},
				},
				PhoneUsage{
					Phone: "0978000001",
					Duration: Usage{
						ActivationDate: time.Date(2019, time.July, 22, 0, 35, 0, 0, time.UTC),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
