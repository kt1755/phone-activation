package tiki

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func ReadCSVFile(path string) (data [][]string, err error) {
	f, err := os.Open(path)
	if err != nil {
		log.WithError(err).Errorf("open file %s error", path)
		return nil, errors.Wrapf(err, "open file %s error", path)
	}
	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				log.Infof("read file %s completed!", path)
				break
			} else {
				log.WithError(err).Errorf("read file %s content error", path)
				return nil, errors.Wrapf(err, "read file %s content error", path)
			}
		}

		data = append(data, record)
	}
	return data, nil
}

func WriteCSVFile(path string, data [][]string) error {

	f, err := os.Create(path)
	if err != nil {
		return errors.Wrapf(err, "cannot create file %s", path)
	}

	csvWriter := csv.NewWriter(f)

	for _, record := range data {
		if err = csvWriter.Write(record); err != nil {
			log.WithError(err).Errorf("cannot write record %v to file", record)
		}
	}

	csvWriter.Flush()
	if err = csvWriter.Error(); err != nil {
		log.WithError(err).Errorf("error while write data to file %s", path)
		return err
	}

	return nil
}
