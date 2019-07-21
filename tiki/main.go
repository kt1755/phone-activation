package main

import (
	"path/filepath"
	"tiki"
	"tiki/phone"

	log "github.com/sirupsen/logrus"
)

const DATE_LAYOUT = "2006-01-02"

func main() {

	log.Print("Application is running...")

	filePath, _ := filepath.Abs("../datatest/data.csv")
	data, err := tiki.ReadCSVFile(filePath)
	if err != nil {
		log.WithError(err).Errorf("get data error")
	}

	log.Infof("raw string %+v", data)

	phoneUsage := phone.TransformToPhoneUsageData(data)

	log.Infof("rawdata: %+v", phoneUsage)

	tiki.QuickSort(phoneUsage)
	log.Infof("sorted data: %+v", phoneUsage)

	mapUsage := phone.CollectData(phoneUsage)

	rsData := phone.TransformToCsvData(mapUsage)

	rsPath, _ := filepath.Abs("../datatest/main_result.csv")
	err = tiki.WriteCSVFile(rsPath, rsData)
	if err != nil {
		panic(err)
	}
}

// func parseStrToTime(ds string) time.Time {
// 	if ds == "" {
// 		return time.Time{}
// 	}
// 	t, err := time.Parse(DATE_LAYOUT, ds)
// 	if err != nil {
// 		log.Infof("parse time %s error", ds)
// 		return time.Time{}
// 	}
// 	return t
// }

// func parseTimeToStr(t time.Time) string {
// 	if t.IsZero() {
// 		return ""
// 	}

// 	return t.Format(DATE_LAYOUT)
// }
