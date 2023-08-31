// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-31 11:17
// @Description:
package conf

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

// 初始化日志：指定日志输出格式，日志文件生成方式以及日志输出模式来生成日志核心，最后借助核心生成zap
func InitLogger() *zap.SugaredLogger {
	//日志模式指定：到debug级还是仅到info模式
	logMode := zapcore.DebugLevel
	if !viper.GetBool("mode.develop") {
		logMode = zapcore.InfoLevel
	}
	//创建核心：指定输出格式、输出源以及输出等级
	core := zapcore.NewCore(getEncode(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), logMode)
	//将使用模式切换为Sugar然后返回
	return zap.New(core).Sugar()

}

// 自定义日志打印格式
func getEncode() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		// 指定时间格式
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 指定日志存放位置及方式:返回一个书写器Writer能将日志按照设定的长度与位置保存起来
func getWriteSyncer() zapcore.WriteSyncer {
	// 设置日志存放位置
	stSeparator := string(filepath.Separator) //文件分隔符:'/'。类似于HTTP状态码，并不直接赋值
	stRootDir, _ := os.Getwd()
	stLogFilePath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format(time.DateOnly)

	// 日志分割
	lumberjackSyncer := &lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxSize:    viper.GetInt("log.MaxSize"),    // 日志文件最大尺寸(M)，超限会分割
		MaxBackups: viper.GetInt("log.MaxBackups"), // 保留旧文件的最大个数
		MaxAge:     viper.GetInt("log.MaxAge"),     // 保留旧文件的最大天数
		Compress:   false,                          // 指定是否对分割的日志文件进行压缩。字段被设置为 false，即不对日志文件进行压缩
	}

	return zapcore.AddSync(lumberjackSyncer)
}
