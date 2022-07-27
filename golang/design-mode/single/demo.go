package main

import (
	"sync"
	"time"
)

type SSH struct {
	User       string
	Password   string
	PkFile     string
	PkPassword string
	Timeout    *time.Duration
}

var ssh *SSH
var sshOnce sync.Once

func GetSshIns() *SSH {
	sshOnce.Do(func() {
		ssh = &SSH{
			User:     "root",
			Password: "123456",
		}
	})

	return ssh
}
