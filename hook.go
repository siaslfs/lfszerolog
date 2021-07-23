package main

import (
	"github.com/rs/zerolog/log"
	"lfszerolog/hookers"
)

func main() {
	hook := &hookers.LfsHooker{}
	hook, _ = hook.NewMysqlHooker("root:root@(127.0.0.1:3306)/logs_db?charset=utf8mb4")
	hooded := log.Hook(hook)
	hooded.Info().Str("foo", "bar").Msg("hello world")
	hooded.Info().Fields(map[string]interface{}{
		"xx": "test",
	}).Msg("hello world")
}
