package main

import (
	"TencentDNSManager/internal/config"
	"TencentDNSManager/internal/dns/NameCom"
	"TencentDNSManager/internal/dns/alibabadns"
	"TencentDNSManager/internal/dns/tencentdns"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var provider string
var domain string
var subDomain string
var recordType string
var value string
var recordLine string
var TTL string

func init() {
	flag.StringVar(&provider, "p", "", "DNS provider (e.g., 腾讯, 阿里云, name.com)")
	flag.StringVar(&domain, "d", "", "Domain name")
	flag.StringVar(&subDomain, "sd", "", "Sub-domain name")
	flag.StringVar(&recordType, "rt", "", "Record type (e.g., A, CNAME)")
	flag.StringVar(&value, "v", "", "Record value (e.g., IP address)")
	flag.StringVar(&recordLine, "rl", "", "Record line (e.g., 默认, default)")
	flag.StringVar(&TTL, "ttl", "300", "Record line (e.g., 300,600)")
}

func main() {
	flag.Parse()

	if provider == "" || domain == "" || subDomain == "" || value == "" {
		fmt.Println("Usage: dnsmanager -p <provider> -d <domain> -sd <subDomain> -rt <recordType> -v <value> ")
		os.Exit(1)
	}

	if provider == "腾讯云" {
		if recordLine == "" {
			recordLine = "默认"
		}

	} else if provider == "阿里云" {
		if recordLine == "" {
			recordLine = "default"
		}
		
	}

	if recordType == "" {
		recordType = "A"
	}

	if subDomain == "" {
		subDomain = "@"
	}
	// 设置信号处理函数
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// 开启协程监听信号
	go func() {
		for sig := range signalChan {
			switch sig {
			case syscall.SIGINT:
				fmt.Println("Exiting due to Ctrl+C...")
			case syscall.SIGTERM:
				fmt.Println("Exiting due to SIGTERM...")
			default:
				fmt.Println("Exiting due to signal:", sig)
			}
			os.Exit(0)
		}
	}()

	switch provider {
	case "腾讯":
		Test1(domain, subDomain, recordType, value, recordLine)

	case "阿里云":
		Test2(domain, subDomain, recordType, value, recordLine)

	case "name.com":
		Test3(domain, subDomain, recordType, value, recordLine)

	default:
		fmt.Println("Unknown provider. Please try again.")
		os.Exit(1)
	}
}

func Test1(domain, subDomain, recordType, value, recordLine string) {
	// 生产密钥管道
	credential := tencentdns.TencentInit()
	// 写入DNS记录
	tencentdns.TencenAddDNS(domain, credential)
	tencentdns.TencenAddRecord(domain, subDomain, recordType, value, recordLine, credential)
}

func Test2(domain, subDomain, recordType, value, recordLine string) {
	// 生产密钥管道
	client := alibabadns.AalibabatInit()
	// 写入DNS记录
	alibabadns.AlibabaAddDomain(domain, client)
	alibabadns.AlibabaAddRecord(domain, subDomain, recordType, value, recordLine, client)
}

func Test3(domain, subDomain, recordType, value, recordLine string) {
	// 读取配置
	comConfig := config.NAMEComConfig()

	// 对组合字符串进行 Base64 编码
	authString := fmt.Sprintf("%s:%s", comConfig.Username, comConfig.Token)
	authEncoded := base64.StdEncoding.EncodeToString([]byte(authString))

	// 写入DNS记录
	NameCom.AddRecord(authEncoded, domain, subDomain, recordType, value)
}
