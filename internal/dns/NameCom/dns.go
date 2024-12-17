package NameCom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DNSRecord struct {
	Host   string `json:"host"`
	Type   string `json:"type"`
	Answer string `json:"answer"`
	TTL    int    `json:"ttl"`
}

func AddRecord(authEncoded, domainName, SubDomain, recordType, value string) {
	client := &http.Client{}

	// 构建 DNS 记录数据
	record := DNSRecord{
		Host:   SubDomain,
		Type:   recordType,
		Answer: value,
		TTL:    300,
	}

	// 将 DNS 记录数据转换为 JSON 格式
	jsonData, err := json.Marshal(record)
	if err != nil {
		fmt.Println("JSON 编码时出错:", err)
		return
	}

	// 构建完整的 API URL
	apiBaseURL := "https://api.name.com/v4/domains/"
	apiURL := fmt.Sprintf("%s%s/records", apiBaseURL, domainName)

	// 创建一个新的 HTTP 请求
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("创建请求时出错:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 设置 Authorization 请求头
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", authEncoded))

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求时出错:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体时出错:", err)
		return
	}

	// 打印响应状态码和响应体
	fmt.Println("响应状态码:", resp.StatusCode)
	fmt.Println("响应体:", string(body))

	if resp.StatusCode == 200 {
		fmt.Println("DNS 记录添加成功")
	} else {
		fmt.Println("重新录入")
	}
}
