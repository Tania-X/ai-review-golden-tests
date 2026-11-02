// Package main 提供 HTTP 请求示例.
package main

import "net/http"

// fetch 发起 GET 请求并返回响应。
func fetch(url string) (*http.Response, error) {
	client := &http.Client{}
	return client.Get(url)
}
