package test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/71010068/tool-box/log"
	"gopkg.in/yaml.v2"
)

// TestLog 测试日志库
func TestLog(t *testing.T) {

	// 定义结构体，引入log日志库LogConfig配置
	type TestLog struct {
		Log log.LogConfig `ymal:"log"`
	}
	// 读取yaml文件中的日志配置
	logYaml, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		fmt.Printf("读取配置文件失败，err：%v", err)
		return
	}

	// 初始化logger结构体实例
	logger := &TestLog{}
	// 解析配置到到logger结构体
	if err := yaml.Unmarshal(logYaml, logger); err != nil {
		fmt.Printf("配置文件解析失败，err：%v", err)
		return
	}

	// 初始化日志
	log := logger.Log.NewLog()

	// 使用
	log.Info("hello")
	log.Debug("hello", log.AnyField("msg", logger))
	log.DebugWithStack(log.AnyField("msg", logger))

}
