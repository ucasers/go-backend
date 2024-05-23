package test

import (
	"testing"

	"github.com/ucasers/go-backend/backend"
)

func TestRun(t *testing.T) {
	// 调用 Run 方法
	backend.Run()

	// 这里可以添加更多测试逻辑，例如检查服务器是否成功启动
	// 由于 Run 方法会启动服务器并阻塞，这里暂时不添加更多测试逻辑
}
