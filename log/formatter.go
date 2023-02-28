package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gtoxlili/give-advice/common/pool"
	"github.com/sirupsen/logrus"
)

/**
1. 时间
2. 日志级别
3. 所在进程号
4. 分割符
5. 所在协程
6. 所在方法
7. 日志内容
*/

type Formatter struct {
	// default: time.Stamp
	TimestampFormat string
	// default: false
	NoUppercaseLevel bool
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	buffer := pool.GetBuffer()
	defer pool.PutBuffer(buffer)

	// 1. 时间
	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.Stamp
	}
	buffer.WriteString(entry.Time.Format(timestampFormat))
	buffer.WriteByte(' ')

	// 2. 日志级别
	levelText := entry.Level.String()
	levelColor := getColorByLevel(entry.Level)
	if !f.NoUppercaseLevel {
		levelText = strings.ToUpper(levelText)
	}
	// [DEBUG] [ INFO] [ WARN] [ERROR] 向右对齐
	buffer.WriteString(fmt.Sprintf("\x1b[%dm[%7s]\x1b[0m", levelColor, levelText))
	buffer.WriteByte(' ')

	// 3. 所在进程号 直接展示
	buffer.WriteString(fmt.Sprintf("\x1b[%dm%d\x1b[0m", Purple, os.Getpid()))
	buffer.WriteString(" -")

	// 4. 所在方法 向右对其42个字符
	if entry.HasCaller() {
		fc := strings.Split(entry.Caller.Function, "/")
		buffer.WriteString(
			fmt.Sprintf(" (%s:%d %s)",
				entry.Caller.File,
				entry.Caller.Line,
				// 去掉包名 */funcName
				fc[len(fc)-1],
			))
		buffer.WriteString(" :")
	}

	// 5. 额外信息 [KEY:VALUE]
	if len(entry.Data) > 0 {
		for k, v := range entry.Data {
			buffer.WriteString(fmt.Sprintf(" %s:%v |", k, v))
		}
		buffer.Truncate(buffer.Len() - 1)
		buffer.WriteByte('>')
	}

	// 6. 日志内容
	buffer.WriteString(" " + entry.Message)
	buffer.WriteByte('\n')

	return buffer.Bytes(), nil
}

type Color uint8

const (
	Green  Color = 32
	Red    Color = 31
	Yellow Color = 33
	Blue   Color = 34
	Purple Color = 35
	Gray   Color = 37
	Cyan   Color = 36
)

func getColorByLevel(level logrus.Level) Color {
	switch level {
	case logrus.DebugLevel, logrus.TraceLevel:
		return Gray
	case logrus.InfoLevel:
		return Green
	case logrus.WarnLevel:
		return Yellow
	default:
		return Red
	}
}
