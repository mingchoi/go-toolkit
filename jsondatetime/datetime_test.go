package jsondatetime_test

import (
	"testing"
	"time"

	"github.com/mingchoi/toolkit/jsondatetime"
)

func TestUnmarshalJSON(t *testing.T) {
	// YYYY-MM
	b := []byte("2018-02")
	expected, _ := time.ParseInLocation("2006-01", "2018-02", time.Now().Location())

	d := jsondatetime.DateTime{}
	err := d.UnmarshalJSON(b)
	if err != nil {
		t.Error(err)
	}

	if !d.Time.Equal(expected) {
		t.Error(err)
	}

	// YYYY-MM-DD
	b = []byte("2000-12-31")
	expected, _ = time.ParseInLocation("2006-01-02", "2000-12-31", time.Now().Location())

	d = jsondatetime.DateTime{}
	err = d.UnmarshalJSON(b)
	if err != nil {
		t.Error(err)
	}

	if !d.Time.Equal(expected) {
		t.Error(err)
	}

	// YYYY-MM-DDTHH-mm
	b = []byte("2036-09-30T23:59")
	expected, _ = time.ParseInLocation("2006-01-02T15:04", "2036-09-30T23:59", time.Now().Location())

	d = jsondatetime.DateTime{}
	err = d.UnmarshalJSON(b)
	if err != nil {
		t.Error(err)
	}

	if !d.Time.Equal(expected) {
		t.Error(err)
	}

}
