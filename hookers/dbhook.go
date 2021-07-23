package hookers

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/rs/zerolog"
	"time"
)

func (lfs *LfsHooker) NewMysqlHooker(mySqlStr string) (*LfsHooker, error) {
	engine, err := xorm.NewEngine("mysql", mySqlStr)
	if err != nil {
		return lfs, err
	}
	return &LfsHooker{engine: engine}, nil
}

type LogInfo struct {
	Id         int64     `xorm:"not null pk BIGINT(20)"`
	LogContent string    `xorm:"default 'NULL' comment('日志内容') TEXT"`
	LogLevel   string    `xorm:"default 'NULL' comment('日志等级') VARCHAR(255)"`
	LogLevelId int       `xorm:"default NULL comment('日志等级id') INT(11)"`
	CreateTime time.Time `xorm:"default 'NULL' comment('创建时间') DATETIME"`
}

func (lfs *LfsHooker) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level != zerolog.NoLevel {
		_logInfo := &LogInfo{
			LogContent: msg,
			LogLevel:   level.String(),
			LogLevelId: int(level),
			CreateTime: time.Now(),
		}
		_, err := lfs.engine.Insert(_logInfo)
		if err != nil {
			e.Str("insert database status", fmt.Sprintf("fail:%v", err))
		} else {
			e.Str("insert database status", "success")
		}
	}
}
