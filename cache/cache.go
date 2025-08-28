package cache

import (
	"log/slog"
	"sync"
)

const (
	ReadThrought  = 1
	WriteThrought = 2
	CacheAside    = 3
	WriteAround   = 4
	WriteBack     = 5
)

type Adapter interface {
	Get() (err error) // 从缓存获取
	Set() (err error) // 保存到缓存
	Del() (err error) // 删除缓存

	DBLoad() (err error)   // 从持久层加载
	DBSave() (err error)   // 保存到持久层
	DBRemove() (err error) // 删除持久层
}

// 读取穿透模式
// 先从缓存中获取数据，如果缓存中不存在，则从持久层加载数据，并将数据保存到缓存中
func ReadThroughtGet(adapter Adapter) (err error) {
	// 先从缓存中获取数据
	if err = adapter.Get(); err == nil {
		return nil
	}

	// 从持久层加载数据
	if err = adapter.DBLoad(); err != nil {
		return err
	}

	// 将数据同步到缓存中
	return adapter.Set()
}

// 保存到持久层
func ReadThroughtSet(adapter Adapter) (err error) {
	return adapter.DBSave()
}

// 删除持久层和缓存的数据
func ReadThroughtDel(adapter Adapter) (err error) {
	if err = adapter.DBRemove(); err != nil {
		return err
	}

	return adapter.Del()
}

// 写穿透模式
// 先将数据保存到持久层，然后再将数据保存到缓存中
func WriteThroughtGet(adapter Adapter) (err error) {
	if err = adapter.Get(); err == nil {
		return nil
	}

	if err = adapter.DBLoad(); err != nil {
		return err
	}

	return adapter.Set()
}

// 同步保存到持久层和缓存
func WriteThroughtSet(adapter Adapter) (err error) {
	if err = adapter.DBSave(); err != nil {
		return err
	}

	return adapter.Set()
}

// 删除持久层和缓存的数据
func WriteThroughtDel(adapter Adapter) (err error) {
	if err = adapter.DBRemove(); err != nil {
		return err
	}

	return adapter.Del()
}

// 旁路缓存模式
// 先从缓存中获取数据，如果缓存中不存在，则从持久层加载数据，并将数据保存到缓存中
func CacheAsideGet(adapter Adapter) (err error) {
	if err = adapter.Get(); err == nil {
		return nil
	}

	if err = adapter.DBLoad(); err != nil {
		return err
	}

	return adapter.Set()
}

func CacheAsideSet(adapter Adapter) (err error) {
	if err = adapter.DBSave(); err != nil {
		return err
	}

	return adapter.Set()
}

func CacheAsideDel(adapter Adapter) (err error) {
	if err = adapter.DBRemove(); err != nil {
		return err
	}

	return adapter.Del()
}

// 绕写模式
// 先从缓存中获取数据，如果缓存中不存在，则从持久层加载数据，并将数据保存到缓存中
// 然后再将数据保存到持久层
func WriteAroundGet(adapter Adapter) (err error) {
	if err = adapter.Get(); err == nil {
		return nil
	}

	if err = adapter.DBLoad(); err != nil {
		return err
	}

	return adapter.Set()
}

func WriteAroundSet(adapter Adapter) (err error) {
	return adapter.DBSave()
}

func WriteAroundDel(adapter Adapter) (err error) {
	return adapter.DBRemove()
}

// 回写模式
// 先从缓存中获取数据，如果缓存中不存在，则从持久层加载数据，并将数据保存到缓存中
// 然后再将数据保存到持久层
func WriteBackGet(adapter Adapter) (err error) {
	if err = adapter.Get(); err == nil {
		return nil
	}

	if err = adapter.DBLoad(); err != nil {
		return err
	}

	return adapter.Set()
}

func WriteBackSet(adapter Adapter) (err error) {
	if err = adapter.Set(); err != nil {
		return err
	}

	chanWriteBack <- adapter.DBSave
	return nil
}

func WriteBackDel(adapter Adapter) (err error) {
	if err = adapter.Del(); err != nil {
		return err
	}

	chanWriteBack <- adapter.DBRemove
	return nil
}

var (
	wgWriteBack   sync.WaitGroup
	chanWriteBack = make(chan func() error, 100)
)

func Init() {
	wgWriteBack.Go(func() {
		for f := range chanWriteBack {
			if err := f(); err != nil {
				slog.Error("WriteBack", "error", err)
			}
		}
	})
}

func Exit() {
	close(chanWriteBack)
	wgWriteBack.Done()
}
