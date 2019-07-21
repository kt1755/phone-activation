package phone

import (
	"reflect"
	"testing"
	"tiki"
	"time"
)

func TestCollectData(t *testing.T) {
	type args struct {
		sortedPhoneUsages []tiki.PhoneUsage
	}
	tests := []struct {
		name string
		args args
		want map[string]*tiki.PhoneUsage
	}{
		// TODO: Add test cases.
		{
			name: "Test",
			args: args{
				sortedPhoneUsages: []tiki.PhoneUsage{
					tiki.PhoneUsage{
						Phone: "1",
						Duration: tiki.Usage{
							ActivationDate:   time.Date(2016, time.January, 1, 00, 00, 00, 00, time.UTC),
							DeactivationDate: time.Date(2016, time.March, 1, 00, 00, 00, 00, time.UTC),
						},
					},
					tiki.PhoneUsage{
						Phone: "1",
						Duration: tiki.Usage{
							ActivationDate:   time.Date(2016, time.March, 1, 00, 00, 00, 00, time.UTC),
							DeactivationDate: time.Date(2016, time.April, 1, 00, 00, 00, 00, time.UTC),
						},
					},
					tiki.PhoneUsage{
						Phone: "1",
						Duration: tiki.Usage{
							ActivationDate:   time.Date(2016, time.April, 1, 00, 00, 00, 00, time.UTC),
							DeactivationDate: time.Time{},
						},
					},
				},
			},
			want: map[string]*tiki.PhoneUsage{
				"1": &tiki.PhoneUsage{
					Phone: "1",
					Duration: tiki.Usage{
						ActivationDate:   time.Date(2016, time.January, 1, 00, 00, 00, 00, time.UTC),
						DeactivationDate: time.Time{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CollectData(tt.args.sortedPhoneUsages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CollectData() = %v, want %v", got, tt.want)
			}
		})
	}
}
