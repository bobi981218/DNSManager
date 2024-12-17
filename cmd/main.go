package main

import (
	"TencentDNSManager/internal/config"
	"TencentDNSManager/internal/dns/NameCom"
	"TencentDNSManager/internal/dns/alibabadns"
	"TencentDNSManager/internal/dns/tencentdns"
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
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

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("现支持解析服务商如下： ")
	fmt.Println("腾讯云")
	fmt.Println("阿里云")
	fmt.Println("Name.com")

	for {
		fmt.Print("请输入解析服务商: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		command := strings.TrimSpace(input)

		switch command {
		case "腾讯云":
			Test1()

		case "阿里云":
			Test2()

		case "name.com":
			Test3()

		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command. Please try again.")
		}
	}
}

func Test1() {

	//生产密钥管道
	credential := tencentdns.TencentInit()
	//录入DNS信息
	d := read()
	//添加域名
	tencentdns.TencenAddDNS(d.Domain, credential)
	//写入DNS记录
	tencentdns.TencenAddRecord(d.Domain, d.SubDomain, d.RecordType, d.Value, d.RecordLine, credential)

	//测试用
	//tencentdns.TencenAddDNS("aoyebobi.pro", credential)
	//tencentdns.TencenAddRecord("aoyebobi.pro", "test", "A", "47.109.55.9", "默认", credential)

}
func Test2() {
	//生产密钥管道
	client := alibabadns.AalibabatInit()
	//录入DNS信息
	d := read()
	//添加域名
	alibabadns.AlibabaAddDomain(d.Domain, client)
	//写入DNS记录
	alibabadns.AlibabaAddRecord(d.Domain, d.SubDomain, d.RecordType, d.Value, d.RecordLine, client)

	//测试用
	//alibabadns.AlibabaAddDomain("aoyebobi.pro", client)
	//alibabadns.AlibabaAddRecord("aoyebobi.pro", "test", "A", "47.109.55.9", "default", client)

}
func Test3() {

	//读取配置
	comConfig := config.NAMEComConfig()

	// 对组合字符串进行 Base64 编码
	authString := fmt.Sprintf("%s:%s", comConfig.Username, comConfig.Token)
	authEncoded := base64.StdEncoding.EncodeToString([]byte(authString))

	d := read()                                                                  //录入DNS信息
	NameCom.AddRecord(authEncoded, d.Domain, d.SubDomain, d.RecordType, d.Value) //写入DNS记录
}

func read() *config.DNS {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("输入域名:")
	domain, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取域名失败:", err)
		return nil
	}
	domain = strings.TrimSpace(domain) // 去除前后空白字符，包括换行符

	fmt.Println("输入子域名，例如 \"www\"")
	SubDomain, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取子域名失败:", err)
		return nil
	}
	SubDomain = strings.TrimSpace(SubDomain)

	fmt.Println("输入类型，例如 \"A\"")
	recordType, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取类型失败:", err)
		return nil
	}
	recordType = strings.TrimSpace(recordType)

	fmt.Println("输入录值，例如 IP 地址")
	value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取值失败:", err)
		return nil
	}
	value = strings.TrimSpace(value)

	fmt.Println("输入线路，例如 腾讯云：\"默认\"，阿里云\"default\"")
	RecordLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取线路失败:", err)
		return nil
	}
	RecordLine = strings.TrimSpace(RecordLine)

	return &config.DNS{
		Domain:     domain,
		SubDomain:  SubDomain,
		RecordType: recordType,
		Value:      value,
		RecordLine: RecordLine,
	}
}
