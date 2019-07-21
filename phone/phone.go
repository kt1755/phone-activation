package phone

import (
	"tiki"
	"time"

	log "github.com/sirupsen/logrus"
)

const DATE_LAYOUT = "2006-01-02"

func CollectData(sortedPhoneUsages []tiki.PhoneUsage) map[string]*tiki.PhoneUsage {
	rs := make(map[string]*tiki.PhoneUsage)

	for _, p := range sortedPhoneUsages {
		mp := rs[p.Phone]
		if mp == nil {
			rs[p.Phone] = &tiki.PhoneUsage{
				Phone:    p.Phone,
				Duration: p.Duration,
			}
		} else {
			if p.Duration.ActivationDate.Equal(mp.Duration.DeactivationDate) {
				mp.Duration.DeactivationDate = p.Duration.DeactivationDate
			} else {
				mp.Duration.ActivationDate = p.Duration.ActivationDate
				mp.Duration.DeactivationDate = p.Duration.DeactivationDate
			}
		}
	}
	return rs
}

func TransformToCsvData(m map[string]*tiki.PhoneUsage) [][]string {
	d := make([][]string, 0)

	for _, p := range m {
		d = append(d, []string{p.Phone, ParseTimeToStr(p.Duration.ActivationDate)})
	}
	return d
}

func TransformToPhoneUsageData(data [][]string) []tiki.PhoneUsage {
	phoneUsage := make([]tiki.PhoneUsage, 0)
	for i, _ := range data {
		log.Info(data[i])
		phoneUsage = append(phoneUsage, tiki.PhoneUsage{
			Phone: data[i][0],
			Duration: tiki.Usage{
				ActivationDate:   ParseStrToTime(data[i][1]),
				DeactivationDate: ParseStrToTime(data[i][2]),
			},
		})
	}
	return phoneUsage
}

func ParseStrToTime(ds string) time.Time {
	if ds == "" {
		return time.Time{}
	}
	t, err := time.Parse(DATE_LAYOUT, ds)
	if err != nil {
		log.Infof("parse time %s error", ds)
		return time.Time{}
	}
	return t
}

func ParseTimeToStr(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format(DATE_LAYOUT)
}
