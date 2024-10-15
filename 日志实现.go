package zapzkratos

import (
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yyle88/erero"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// T奎沱日志实现 是将 zap.Logger 转换为 kratos log.Logger 的适配器
type T奎沱日志实现 struct {
	zap日志 *zap.Logger
	msg模块 string
}

// New奎沱日志 创建一个新的 T奎沱日志实现 的对象实现 kratos log.Logger 的接口
func New奎沱日志(zap日志 *zap.Logger, msg模块 string) log.Logger {
	return &T奎沱日志实现{
		zap日志: zap日志,
		msg模块: msg模块,
	}
}

// Log 实现 kratos log.Logger 接口
func (a *T奎沱日志实现) Log(arg日志级别 log.Level, keyvals ...interface{}) error {
	var zap日志级别 zapcore.Level
	switch arg日志级别 {
	case log.LevelDebug:
		zap日志级别 = zap.DebugLevel
	case log.LevelInfo:
		zap日志级别 = zap.InfoLevel
	case log.LevelWarn:
		zap日志级别 = zap.WarnLevel
	case log.LevelError:
		zap日志级别 = zap.ErrorLevel
	case log.LevelFatal:
		zap日志级别 = zap.DPanicLevel
	default:
		zap日志级别 = zap.DebugLevel //假如不确定级别就设置为最低级别的
	}

	// 使用 zap.Logger 记录日志
	zap按级日志 := a.zap日志.Check(zap日志级别, a.msg模块)
	if zap按级日志 == nil {
		return erero.Errorf("日志级别错误 zap=%v arg=%v", zap日志级别, arg日志级别)
	}

	var fields = make([]zap.Field, 0, (len(keyvals)+1)/2) //避免奇数
	for idx := 0; idx < len(keyvals); idx += 2 {
		if idx+1 < len(keyvals) {
			kes, ok := keyvals[idx].(string)
			if !ok {
				kes = newBK(idx) //因为日志是有可能要打印json的，因此这里的键还是组个正确的
			}
			fields = append(fields, zap.Any(kes, keyvals[idx+1]))
		}
	}
	if len(keyvals)%2 == 1 { //说明是奇数的，说明最后1个只有值没有键，而不是只有值没有键，因为键更难漏掉
		fields = append(fields, zap.Any(newBK(len(keyvals)-1), keyvals[(len(keyvals)-1)]))
	}

	zap按级日志.Write(fields...)
	return nil
}

func newBK(idx int) string {
	return "BAD_KEY_" + strconv.Itoa(idx)
}
