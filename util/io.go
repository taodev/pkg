package util

import "io"

// WriteAtLeast 尝试将 buf 中的数据写入 w，直到成功写入至少 min 字节或发生错误。
// 若写入过程中已写入部分数据（n > 0）但遇到 EOF，会返回 io.ErrUnexpectedEOF。
// 当实际写入字节数 n >= min 时，错误会被置为 nil（表示已满足最小写入要求）。
//
// 参数：
//
//	w: 实现了 io.Writer 接口的写入目标
//	buf: 待写入的数据缓冲区（实际写入长度不超过 len(buf)）
//	min: 需要至少写入的字节数（需满足 min <= len(buf)，否则可能无法完成）
//
// 返回：
//
//	n: 实际成功写入的字节数（可能大于等于 min 或小于 min 但发生错误）
//	err: 错误信息（成功写入 >= min 字节时为 nil；写入未达 min 时返回具体错误）
func WriteAtLeast(w io.Writer, buf []byte, min int) (n int, err error) {
	for n < min && err == nil {
		var nn int
		nn, err = w.Write(buf[n:])
		n += nn
		if n >= min {
			err = nil
		} else if n > 0 && err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}
	return
}

// WriteFull 尝试将 p 中的全部数据完整写入 w，直到成功写入 len(p) 字节或发生错误。
// 该函数是 WriteAtLeast 的便捷封装，等价于调用 WriteAtLeast(w, p, len(p))。
// 当成功写入全部 len(p) 字节时返回的 err 为 nil；若写入未完成则返回具体错误。
//
// 参数：
//
//	w: 实现了 io.Writer 接口的写入目标
//	p: 待完整写入的数据缓冲区（要求写入长度为 len(p)）
//
// 返回：
//
//	n: 实际成功写入的字节数（可能等于 len(p) 或小于 len(p) 但发生错误）
//	err: 错误信息（成功写入全部数据时为 nil；写入未完成时返回具体错误）
func WriteFull(w io.Writer, p []byte) (n int, err error) {
	return WriteAtLeast(w, p, len(p))
}
