package core

import (
	"fmt"
	"github.com/twelvet-s/gins/core/internal"
	"github.com/twelvet-s/gins/g"
	"github.com/twelvet-s/gins/utils/file"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	if ok, _ := file.PathExists(g.CONFIG.Gins.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", g.CONFIG.Gins.Zap.Director)
		_ = os.Mkdir(g.CONFIG.Gins.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if g.CONFIG.Gins.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
