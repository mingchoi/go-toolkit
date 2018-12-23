package jsondatetime

import (
	"strings"
	"time"
)

type DateTime struct{ time.Time }

func (d *DateTime) UnmarshalJSON(b []byte) error {
	str := strings.Replace(string(b), "\"", "", -1)

	var t time.Time
	var err error
	if len(str) == 7 {
		t, err = time.ParseInLocation("2006-01", str, time.Now().Location())
	} else if len(str) == 10 {
		t, err = time.ParseInLocation("2006-01-02", str, time.Now().Location())
	} else {
		t, err = time.ParseInLocation("2006-01-02T15:04", str, time.Now().Location())
	}

	if err != nil {
		return err
	}

	*d = DateTime{t}
	return nil
}
