package loaderservice

import (
	"time"

	"github.com/cockroachdb/errors"
)

func (r *Service) getImageDate(filepath string) (time.Time, error) {
	const createDateTag = "CreateDate"

	fileInfos := r.et.ExtractMetadata(filepath)
	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			continue
		}

		for k, v := range fileInfo.Fields {
			if k == createDateTag {
				timeStr, ok := v.(string)
				if !ok {
					continue
				}

				timeParsed, err := time.Parse("2006:01:02:15:04:05", timeStr)
				if err != nil {
					continue
				}

				timeParsed = timeParsed.UTC()

				return timeParsed, nil
			}
		}
	}

	return time.Time{}, errors.New("no date found")
}
