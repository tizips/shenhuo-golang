package helper

import "time"

const ImportPasswordSep = "@"

func PersonImportPassword(mobile string, at time.Time) string {
	tail := mobile
	if len(mobile) >= 5 {
		tail = mobile[len(mobile)-5:]
	}

	return tail + ImportPasswordSep + at.Format("20060102")
}
