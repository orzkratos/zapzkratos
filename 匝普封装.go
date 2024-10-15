package zapzkratos

import (
	"path/filepath"

	"github.com/yyle88/runpath"
	"github.com/yyle88/zaplog"
)

type T日志配置 struct{}

func New日志配置() *T日志配置 {
	return &T日志配置{}
}

type T匝普日志 zaplog.Zap //这里主要做的就是把英文的类型名改为中文类型名，再增加几个简单的逻辑

func New创建匝普日志(cfg *T日志配置) *T匝普日志 {
	//没有用到这个参数，只是预留的选项

	return New匝普日志(zaplog.MustNewZap(zaplog.NewConfig()))
}

func New匝普日志(zap封装 *zaplog.Zap) *T匝普日志 {
	return (*T匝普日志)(zap封装)
}

func (A *T匝普日志) Get基本匝普() *zaplog.Zap {
	return (*zaplog.Zap)(A)
}

func (A *T匝普日志) Sub模块匝普() *zaplog.Zap {
	zpx := A.Get基本匝普()

	return zpx.SubZap2("模块位置", filepath.Base(runpath.Skip(1)))
}
