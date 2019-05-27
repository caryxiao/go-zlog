# go-zlog

> 日志插件, 用于输出日志到指定位置, 方便调试, 本库是对 [logrus](github.com/sirupsen/logrus) 进行了封装



日志所用到的库

[logrus](github.com/sirupsen/logrus) 日志

[logrus-prefixed-formatter](https://github.com/x-cray/logrus-prefixed-formatter) 日志格式化



使用方法

```go
import "github.com/caryxiao/go-zlog"
func main() {
  zlog.SetOutput("/tmp/zlog.log")
  zlog.SetLevel(5)
  zlog.SetFormat("[%level%]: %time% - [%trace_id%] %msg%")
  zlog.Logger.Debug("test debug")
}

```

