package tencentdns

import (
	"TencentDNSManager/internal/config"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"log"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

func TencentInit() *common.Credential {
	// 加载配置
	cfg, err := config.TencentConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 实例化一个认证对象，使用从配置中加载的 SecretId 和 SecretKey
	return common.NewCredential(
		cfg.SECRET_ID,
		cfg.SECRET_KEY,
	)
}
func TencenAddDNS(domain string, c *common.Credential) {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"

	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(c, regions.Guangzhou, cpf)

	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	request := dnspod.NewCreateDomainRequest()

	// 域名
	request.Domain = common.StringPtr(domain) // 替换为你想要添加的域名

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.CreateDomain(request)
	//处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Println("域名已添加至服务列表")
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	//if err != nil {
	//	log.Println(err)
	//}
	// 打印返回的json字符串

	fmt.Printf("%s\n", response.ToJsonString())
}
func TencenAddRecord(domain, SubDomain, recordType, value, RecordLine string, c *common.Credential) {
	// 实例化一个 client 选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	// 这里指定地域信息，如果不填，默认为广州，但建议您传入地域信息
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"

	// 实例化要请求产品的 client 对象，clientProfile 是可选的
	client, err := dnspod.NewClient(c, "", cpf)

	if err != nil {
		log.Fatalf("Failed to create client: %v\n", err)
		return
	}

	request := dnspod.NewCreateRecordRequest()

	// 填充请求参数
	request.Domain = common.StringPtr(domain)         // 替换为你的域名
	request.SubDomain = common.StringPtr(SubDomain)   // 子域名，例如 "www"
	request.RecordType = common.StringPtr(recordType) // 记录类型，例如 "A"
	request.Value = common.StringPtr(value)           // 记录值，例如 IP 地址
	request.RecordLine = common.StringPtr(RecordLine) // 记录线路，例如 "默认"

	// 调用接口进行操作
	response, err := client.CreateRecord(request)
	if err != nil {

		log.Printf("Failed to create DNS record: %v\n", err)
		fmt.Println("重新录入")
		return
	}

	fmt.Println(response.ToJsonString())
	fmt.Println("继续录入")

}
