package alibabadns

import (
	"TencentDNSManager/internal/config"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"log"
)

func AalibabatInit() *alidns.Client {

	// 加载配置
	cfg, err := config.AlibabaConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	//创建客户端，使用配置文件中的ID和key
	client, err := alidns.NewClientWithAccessKey("cn-hangzhou",
		cfg.AccessKeyID,
		cfg.AccessKeySecret)
	if err != nil {
		fmt.Println("Failed to create client:", err)
		return nil
	}

	return client
}

func AlibabaAddDomain(domain string, c *alidns.Client) {
	// 创建 AddDomainRequest 请求
	request := alidns.CreateAddDomainRequest()
	request.Scheme = "https"

	// 设置请求参数
	request.DomainName = domain // 替换为你的域名

	// 发送请求
	response, err := c.AddDomain(request)
	if err != nil {
		fmt.Println("域名已添加至域名列表")
		return
	}

	// 打印响应结果
	fmt.Printf("%s\n", response)
	fmt.Println("域名添加已完成")

}

func AlibabaAddRecord(domain, SubDomain, recordType, value, RecordLine string, c *alidns.Client) {
	// 创建 AddDomainRecordRequest 请求
	request := alidns.CreateAddDomainRecordRequest()
	request.Scheme = "https"

	// 设置请求参数
	request.DomainName = domain // 替换为你的域名
	request.RR = SubDomain      // 子域名，例如 "www"
	request.Type = recordType   // 记录类型，例如 "A"
	request.Value = value       // 记录值，例如 IP 地址
	request.Line = RecordLine   // 记录线路，例如 "default"
	//request.TTL = "600"

	// 发送请求
	response, err := c.AddDomainRecord(request)
	if err != nil {
		fmt.Println("Failed to add DNS record:", err)
		fmt.Println("重新录入")
		return
	}

	// 打印响应结果
	fmt.Printf("%s\n", response)
	fmt.Println("DNS记录添加完成")

}
