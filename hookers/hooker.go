package hookers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type LfsHooker struct {
	engine *xorm.Engine
}
