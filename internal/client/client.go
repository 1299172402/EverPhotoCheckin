package client

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/winterssy/EverPhotoCheckin/internal/model"
	"github.com/winterssy/ghttp"
)

const (
	_apiEntry   = "https://api.everphoto.cn"
	_apiAuth    = _apiEntry + "/auth"
	_apiCheckin = _apiEntry + "/users/self/checkin/v2"
)

var (
	_headers = ghttp.Headers{
		"user-agent":  "EverPhoto/2.8.0",
		// "user-agent":  "EverPhoto/2.8.0 (Android;28000;16s;28;tengxun_33_1)",
		"x-device-id": "02:00:00:00:00:00",
		"application": "tc.everphoto",
		"x-locked":    "1",
	}
)

type Bot struct {
	client *ghttp.Client
}

const _salt = "tc.everphoto."

func salt(value string) string {
	hash := md5.Sum([]byte(_salt + value))
	return hex.EncodeToString(hash[:])
}

func New(mobile, password string) (*Bot, error) {
	client := ghttp.New()
	client.RegisterBeforeRequestCallbacks(
		ghttp.WithHeaders(_headers),
		ghttp.EnableRetry(),
	)

	resp, err := client.Post(_apiAuth, ghttp.WithForm(ghttp.Form{
		"mobile":   mobile,
		"password": salt(password),
	}))
	if err != nil {
		return nil, err
	}

	ar := new(model.AuthResponse)
	if err = resp.JSON(ar); err != nil {
		return nil, err
	}

	if ar.Code != 0 {
		return nil, errors.New(ar.Message)
	}

	client.RegisterBeforeRequestCallbacks(ghttp.WithBearerToken(ar.Data.Token))
	return &Bot{client: client}, nil
}

func NewWithToken(token string) *Bot {
	client := ghttp.New()
	client.RegisterBeforeRequestCallbacks(
		ghttp.WithHeaders(_headers),
		ghttp.WithBearerToken(token),
		ghttp.EnableRetry(),
	)
	return &Bot{client: client}
}

func (bot *Bot) Checkin(ctx context.Context) (*model.CheckinResult, error) {
	resp, err := bot.client.Post(_apiCheckin,
		ghttp.WithContext(ctx),
	)
	if err != nil {
		return nil, err
	}

	cr := new(model.CheckinResponse)
	if err = resp.JSON(cr); err != nil {
		return nil, err
	}

	if cr.Code != 0 {
		return nil, errors.New(cr.Message)
	}

	return cr.Data, nil
}
