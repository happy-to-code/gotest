/**
 * @Author: FB
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2019/9/7 14:01
 */
package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"strconv"
)

func SendMail(mailTo []string, subject string, body string) error {
	// 定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "zhangyfei126@yeah.net",
		"pass": "a2271778df21437",
		"host": "smtp.yeah.net",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) // 转换端口类型为int


	m := gomail.NewMessage()
	// 这种方式可以添加别名，即“XX官方”
	// 说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	m.SetHeader("From", m.FormatAddress(mailConn["user"], "张益飞官方"))
	// m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">")
	// 这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	// m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo...)    // 发送给多个用户
	m.SetHeader("Subject", subject) // 设置邮件主题
	m.SetBody("text/html", body)    // 设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}
func main() {
	// 定义收件人
	mailTo := []string{
		"yida8080@gmail.com",
		"zhangyifei@wutongchain.com",
	}
	// 邮件主题为"Hello"
	subject := "Hello.  测试gomail发送邮件！！邮件来之网易邮箱"
	// 邮件正文
	body := "Hello,by gomail sent"

	err := SendMail(mailTo, subject, body)
	if err != nil {
		log.Println(err)
		fmt.Println("send fail")
		return
	}

	fmt.Println("send successfully")
}
