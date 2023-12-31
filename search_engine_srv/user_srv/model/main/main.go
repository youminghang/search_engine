package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github/youminghang/search_engine/search_engine_srv/user_srv/model"
	"io"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

// 连接数据库以及同步数据库
func main() {
	dsn := "root:123456@tcp(139.9.49.75:3306)/se_user_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&model.User{})
	fmt.Println("建表完成")
	// Using custom options
	/*
		options := &password.Options{16, 100, 32, sha512.New}
		salt, encodedPwd := password.Encode("generic password", options)
		newpassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
		fmt.Printf("%d %d %d", len(newpassword), len(salt), len(encodedPwd))

		passwordInfo := strings.Split(newpassword, "$")
		fmt.Println(passwordInfo)
		check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
		fmt.Println(check) // true
	*/
}
