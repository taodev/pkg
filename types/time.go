package types

import (
	"time"
)

type Time int64

func (t *Time) Format(v string) {
	tm, err := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
	if err != nil {
		panic(err)
	}
	*t = Time(tm.Unix())
}

func (t Time) String() string {
	return time.Unix(int64(t), 0).Format("2006-01-02 15:04:05")
}

func (t Time) ToTime() time.Time {
	return time.Unix(int64(t), 0)
}

func (t *Time) FromTime(v time.Time) {
	*t = Time(v.Unix())
}

// TimeWindow 计算时间戳所属的对齐窗口基准值
// 参数：
//
//	ts - 原始时间戳（单位需与tolerance一致，如秒/毫秒）
//	tolerance - 时间窗口总长度（例如设置为60表示60单位的窗口）
//
// 返回值：当前时间戳所属窗口的中心基准值（窗口中心点位置）
// 逻辑说明：通过将时间戳对齐到窗口中心，确保同一窗口内的时间戳归属于相同基准值
func TimeWindow(ts, tolerance int64) (out int64) {
	// 窗口中心偏移量（窗口大小的一半），用于调整对齐基准点
	offset := tolerance / 2
	// 计算逻辑：
	// 1. (ts - offset) 将原始时间戳调整为相对于窗口中心的偏移值
	// 2. / tolerance 通过整数除法计算完整窗口数（向下取整）
	// 3. * tolerance 还原为基准窗口的起始位置
	// 4. + offset 将基准位置调整到窗口中心
	return ((ts-offset)/tolerance)*tolerance + offset
}
