package main

import (
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"user_server/migration/model"
)

type Config struct {
	DataSource string `yaml:"DataSource"`
}

func main() {
	// 从配置文件加载数据库配置
	config := loadConfig("etc/userServer.yaml")

	// 构建 DSN 连接字符串
	dsn := config.DataSource

	// 连接 MySQL 数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("无法连接到SQL数据库: %v", err)
	}

	// 自动迁移 schema
	err = db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.UserRole{},
		&model.Permissions{},
		&model.RolePermissions{},
	)
	if err != nil {
		log.Fatalf("无法自动迁移数据库架构: %v", err)
	}

	log.Println("数据库迁移成功！")
}

// loadConfig 从 YAML 配置文件中加载数据库配置
func loadConfig(filename string) *Config {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("YAML解析成Go结构体失败: %v", err)
	}

	return &config
}
