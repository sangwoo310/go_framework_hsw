package main

import "go_framework_hsw/telegram"

type HanFramwork interface {

}

type Provider struct {
	telegram telegram.TelegramBot
}

func New() (provider *Provider, err error) {
	p := &Provider{}

	return p,err
}