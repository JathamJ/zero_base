package utilx

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadFile 下载文件并保存到指定路径
func DownloadFile(url string, filename string) error {
	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("http.Get failed, err: %v", err)
	}
	defer resp.Body.Close()

	// 检查 HTTP 响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// 创建文件
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create file failed, err: %v", err)
	}
	defer out.Close()

	// 将响应体内容写入文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("copy data failed, err: %v", err)
	}
	return nil
}
