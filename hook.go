package main

import (
	"github.com/rs/zerolog/log"
	"lfszerolog/hookers"
)

func HookerMysql() {
	hook := &hookers.LfsHooker{}
	hook, _ = hook.NewMysqlHooker("root:123456@(127.0.0.1:3306)/logs_db?charset=utf8mb4")
	hooded := log.Hook(hook)
	hooded.Info().Msg("hello world")
}
