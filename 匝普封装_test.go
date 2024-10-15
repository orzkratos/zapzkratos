package zapzkratos

import (
	"testing"
)

func TestNew创建匝普日志(t *testing.T) {
	v匝普日志 := New匝普日志(New日志配置())

	v模块匝普 := v匝普日志.Sub模块匝普()

	v模块匝普.LOG.Info("abc")
	v模块匝普.SUG.Info("xyz")
}
