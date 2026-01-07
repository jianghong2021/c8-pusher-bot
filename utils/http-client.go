package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// Get 发送GET请求
func Get(url string, headers map[string]string) (string, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Post 发送POST请求
func Post(url string, data map[string]string, headers map[string]string) (string, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	form, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(form))
	if err != nil {
		return "", err
	}

	// 设置请求头
	req.Header.Set("content-type", "application/json; charset=utf-8")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
