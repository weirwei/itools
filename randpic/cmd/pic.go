package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

const (
	baseURL = "https://api.unsplash.com"
)

type UnsplashResponse struct {
	URLs struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	} `json:"urls"`
}

func getRandomPic() error {
	// 获取 API key
	accessKey := os.Getenv("UNSPLASH_ACCESS_KEY")
	if accessKey == "" {
		return fmt.Errorf("请设置 UNSPLASH_ACCESS_KEY 环境变量")
	}

	// 构建 API URL
	url := fmt.Sprintf("%s/photos/random", baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %v", err)
	}

	// 添加查询参数
	q := req.URL.Query()
	if query != "" {
		q.Add("query", query)
	}
	if collectionID != "" {
		q.Add("collections", collectionID)
	}
	req.URL.RawQuery = q.Encode()

	// 添加认证头
	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", accessKey))

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应失败: %v", err)
	}

	// 解析 JSON
	var photo UnsplashResponse
	if err := json.Unmarshal(body, &photo); err != nil {
		return fmt.Errorf("解析响应失败: %v", err)
	}

	// 准备输出内容
	var output strings.Builder
	if detail {
		output.WriteString("随机图片链接:\n")
		fmt.Fprintf(&output, "原始尺寸: %s\n", photo.URLs.Raw)
		fmt.Fprintf(&output, "全尺寸: %s\n", photo.URLs.Full)
		fmt.Fprintf(&output, "常规尺寸: %s\n", photo.URLs.Regular)
		fmt.Fprintf(&output, "小尺寸: %s\n", photo.URLs.Small)
		fmt.Fprintf(&output, "缩略图: %s\n", photo.URLs.Thumb)
	} else {
		output.WriteString(photo.URLs.Raw)
	}

	// 输出结果
	result := output.String()

	// 如果指定了复制到剪贴板
	if toClip {
		if err := clipboard.WriteAll(result); err != nil {
			return fmt.Errorf("复制到剪贴板失败: %v", err)
		}
	} else {
		fmt.Println(result)
	}
	return nil
}
