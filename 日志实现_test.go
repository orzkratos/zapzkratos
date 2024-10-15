package zapzkratos

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yyle88/zaplog"
)

func TestNew奎沱日志(t *testing.T) {
	v奎沱日志 := New奎沱日志(zaplog.LOG, "测试日志实现")

	v日志主簿 := log.NewHelper(v奎沱日志)

	v日志主簿.Infow("k", "v", "k1", "v2")
	v日志主簿.Infow("k", "v", "v2")
}
