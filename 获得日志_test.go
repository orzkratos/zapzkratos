package zapzkratos

import (
	"testing"

	"github.com/yyle88/zaplog"
)

func Test_T匝普日志_Get奎沱主簿(t *testing.T) {
	v匝普日志 := New创建匝普日志(New日志配置())

	v奎沱主簿 := v匝普日志.Get奎沱主簿("测试用例")

	v奎沱主簿.Info("woca", "[a]", "[b]", "[c]")
	v奎沱主簿.Infow("k", "v", "k1", "v2")
}

func TestT匝普日志_Get奎沱秘书(t *testing.T) {
	v匝普日志 := New匝普日志(zaplog.LOGGER)

	v奎沱秘书 := v匝普日志.Get奎沱秘书("测试秘书")

	v奎沱秘书.Info("我爱杨亦乐")
	v奎沱秘书.Infow("我爱", "杨亦乐")
}
