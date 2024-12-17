package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DNS struct {
	Domain     string
	SubDomain  string
	RecordType string
	Value      string
	RecordLine string
}

// 所有配置参数（可添加）
type Config struct {
	SECRET_ID       string
	SECRET_KEY      string
	AccessKeyID     string
	AccessKeySecret string
	Username        string
	TokenName       string
	Token           string
}

type Tencent struct {
	SECRET_ID  string
	SECRET_KEY string
}

type Alibaba struct {
	AccessKeyID     string
	AccessKeySecret string
}

type NAMEcom struct {
	Username  string
	TokenName string
	Token     string
}

// 读取配置文件所有信息（配置改变时更改）
func loadConfig() (*Config, error) {
	// 加载 .env 文件(自己的路径)
	err := godotenv.Load("/Users/bobi/codes/TencentDNSManager/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return nil, err
	}

	// 从环境变量中读取配置
	secretID := os.Getenv("SECRET_ID")
	secretKey := os.Getenv("SECRET_KEY")
	accessKeyID := os.Getenv("Access_Key_ID")
	accessKeySecret := os.Getenv("Access_Key_Secret")
	Username := os.Getenv("Username")
	TokenName := os.Getenv("Token_Name")
	Token := os.Getenv("Token")

	return &Config{
		SECRET_ID:       secretID,
		SECRET_KEY:      secretKey,
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		Username:        Username,
		TokenName:       TokenName,
		Token:           Token,
	}, nil
}

// 腾讯云配置列表
func TencentConfig() (*Tencent, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	return &Tencent{
		SECRET_ID:  config.SECRET_ID,
		SECRET_KEY: config.SECRET_KEY,
	}, nil
}

// 阿里云配置列表
func AlibabaConfig() (*Alibaba, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	return &Alibaba{
		AccessKeyID:     config.AccessKeyID,
		AccessKeySecret: config.AccessKeySecret,
	}, nil
}

// DNSimple配置列表
func NAMEComConfig() *NAMEcom {
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return nil
	}

	return &NAMEcom{
		Username:  config.Username,
		TokenName: config.TokenName,
		Token:     config.Token,
	}
}
