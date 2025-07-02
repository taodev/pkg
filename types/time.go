package types

import "time"

type Time int64

func (t *Time) Format(v string) {
	tm, err := time.ParseInLocation("2006-01-02", v, time.Local)
	if err != nil {
		panic(err)
	}
	*t = Time(tm.Unix())
}

func (t Time) String() string {
	return time.Unix(int64(t), 0).Format("2006-01-02")
}
