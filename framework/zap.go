package framework

import (
	"fmt"
	"github.com/twelvet-s/gins/framework/global"
	"github.com/twelvet-s/gins/framework/internal"
	"github.com/twelvet-s/gins/framework/utils/file"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	if ok, _ := file.PathExists(global.CONFIG.Gins.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.CONFIG.Gins.Zap.Director)
		_ = os.Mkdir(global.CONFIG.Gins.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.CONFIG.Gins.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
