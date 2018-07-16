package repository

import (
	"github.com/go-xorm/xorm"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"flag"
	"os"
	"runtime"
	"path"
	"path/filepath"
	"strings"
	"github.com/go-xorm/core"
)


const (
	MysqlDatasource = "dataSourceName pattern:'username:passwd@tcp(ip:port)/databaseName?charset=utf8'"
	DefaultPattern  = "数据库账号:数据库连接密码@tcp(数据库地址:端口)/数据库实例名称?charset=utf8"
	DEBUG = true
)
var (
	engine *xorm.Engine
	dataSourceName = flag.String("mysql_driver",
		"nsmjpt:mjpt@tcp(112.74.17.193:3306)/grafana?charset=utf8", MysqlDatasource)
)

func usage() {
	log.Fatalf("Usage: --mysql_driver=%s", DefaultPattern)
}




func init(){
	log.Output(2, fmt.Sprintf("repository 包init初始化加载数据库"))
	flag.Parse()
	if *dataSourceName == "" {
		usage()
		os.Exit(1)
	}
	var err error
	if engine, err = mysqlEngine(*dataSourceName); err != nil {
		log.Fatal("加载mysql驱动失败，error："+err.Error())
		os.Exit(2)
	}
	log.Output(2, "成功加载mysql驱动！")

	//设置连接池
	//设置连接池的空闲数大小
	engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	engine.SetMaxOpenConns(150)

	if DEBUG {
		//日志打印SQL
		engine.ShowSQL(true)
		engine.Logger().SetLevel(core.LOG_DEBUG)

		//设置SQL日志文件
		f, err := os.Create("sql.log")
		if err != nil {
			println(err.Error())
			return
		}
		engine.SetLogger(xorm.NewSimpleLogger(f))
		//也支持系统日志syslog
	}
}

func mysqlEngine(dataSource string) (*xorm.Engine, error) {
	//格式:"数据库账号:数据库连接密码@tcp(数据库地址:端口)/数据库实例名称?charset=utf8"
	return xorm.NewEngine("mysql", dataSource)
}

func GetMySQLEngine(ping bool) *xorm.Engine{
	//连接测试
	if ping {
		if err := engine.Ping(); err!=nil{
			log.Output(2, err.Error())
			return nil
		}
	}
	return engine
}

type Customer struct {
	Id       int64       `xorm:"not null pk autoincr INT(11)"`
	Username string    `xorm:"not null VARCHAR(32)"`
	Birthday time.Time `xorm:"DATE"`
	Sex      string    `xorm:"CHAR(1)"`
	Phone    string    `xorm:"VARCHAR(11)"`
	Address  string    `xorm:"VARCHAR(256)"`
}

func GetCurrentFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}
	filename := []string{"config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	fmt.Println("dirname:"+ dirname)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}

func SaveCustomer(customer *Customer)(int64, error){
	return engine.InsertOne(customer)
}

func SaveCustomers(customers []Customer)(int64, error){
	return engine.Insert(&customers)
}

//---------------update ----------------------
func UpdatePhoneNOByCustomerId(customer *Customer)(int64, error){
	return engine.Id(customer.Id).Update(customer)
}

func UpdateCustomerByMap(id int64, props map[string]interface{})(int64, error){
	return engine.Table(new(Customer)).ID(id).Update(props)
}