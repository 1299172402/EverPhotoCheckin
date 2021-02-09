package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/winterssy/EverPhotoCheckin/internal/client"
	"github.com/winterssy/EverPhotoCheckin/internal/push"
)

const (
	EnvEverPhotoMobile   = "EverPhotoMobile"
	EnvEverPhotoPassword = "EverPhotoPassword"
	EnvEverPhotoToken    = "EverPhotoToken"
	EnvServerChanKey     = "ServerChanKey"
)

var (
	_mobile   string
	_password string
	_token    string
)

func init() {
	flag.StringVar(&_mobile, "mobile", "", "your mobile phone number")
	flag.StringVar(&_password, "password", "", "your password")
	flag.StringVar(&_token, "token", "", "your token")
}

func valueOrDefault(value, def string) string {
	if value != "" {
		return value
	}
	return def
}

func createBot() (bot *client.Bot, err error) {
	_token = valueOrDefault(_token, os.Getenv(EnvEverPhotoToken))
	if _token != "" {
		bot = client.NewWithToken(_token)
		return
	}

	_mobile = valueOrDefault(_mobile, os.Getenv(EnvEverPhotoMobile))
	_password = valueOrDefault(_password, os.Getenv(EnvEverPhotoPassword))
	bot, err = client.New(_mobile, _password)
	return
}

var _scKey = os.Getenv(EnvServerChanKey)

func pushMessage(ok bool, desp string) {
	if _scKey == "" {
		return
	}

	const (
		_textSuccess = "【时光相册】签到成功"
		_textFailure = "【时光相册】签到失败"
	)
	var err error
	if ok {
		err = push.Push(_scKey, _textSuccess, desp)
	} else {
		err = push.Push(_scKey, _textFailure, "错误详情 >> "+desp)
	}

	if err != nil {
		log.Print("Server酱推送失败：" + err.Error())
	}
}

func errDesp(msg string, err error) string {
	return msg + "：" + err.Error()
}

func main() {
	flag.Parse()

	var desp string
	bot, err := createBot()
	if err != nil {
		desp = errDesp("登录失败", err)
		pushMessage(false, desp)
		log.Fatal("【时光相册】" + desp)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	cr, err := bot.Checkin(ctx)
	if err != nil {
		desp = errDesp("签到失败", err)
		pushMessage(false, desp)
		log.Fatal("【时光相册】" + desp)
	}

	desp = fmt.Sprintf("\n你已连续签到%d天\n累计获得空间%s\n明天可再白嫖%s\n\n请继续保持(￣▽￣)",
		cr.Continuity, cr.TotalReward, cr.TomorrowReward)
	pushMessage(true, desp)
	log.Print("【时光相册】" + desp)
}
