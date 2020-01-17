package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	masterDataSourceName = "root:root@/api_base?charset=utf8"
	slave1DataSourceName = "root:root@/api_base?charset=utf8"
	masterEngine         *xorm.Engine
	slave1Engine         *xorm.Engine
	engineGroup          *xorm.EngineGroup
	err                  error
)

func init() {
	if masterEngine, err = xorm.NewEngine("mysql", masterDataSourceName); err != nil {
		fmt.Println("主库连接失败:", err.Error())
	}

	if slave1Engine, err = xorm.NewEngine("mysql", slave1DataSourceName); err != nil {
		fmt.Println("从库连接失败:", err.Error())
	}
	if engineGroup, err = xorm.NewEngineGroup(masterEngine, []*xorm.Engine{slave1Engine}); err != nil {
		fmt.Println("数据库组连接失败:", err.Error())
	}
}

func EngineInstance() *xorm.EngineGroup {
	return engineGroup
}
