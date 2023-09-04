package db

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"log"
	"time"
	"web-socket-show/conf"
)

// ConfigInit 配置文件初始化
func ConfigInit() {
	if err := conf.ParserConfig("/conf/tsconfig.json"); err != nil {
		fmt.Printf("Error: %#v\n", err)
		//panic(err)
	}

}

type Msg struct {
	Id     int64 `json:"id"`
	Time   time.Time
	Select string
	Send   string
}

func main() {
	//ConfigInit()
	//test
	m1 := Msg{Id: int64(21), Select: "copy", Send: "hello"}
	//insert
	err := InsertOper(m1)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Printf("insert succeeded!")

	//select
	//msgs, err := SelectOper("`select_op`='readWrite'")
	//msgs, err := SelectOper()
	msgs, err := SelectOper("`id`=21")
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Printf("find %d numbers.\n", len(msgs))
	for _, m := range msgs {
		fmt.Printf("Id: %d, Create_time: %s, Select_op: %s, Send_value: %s\n",
			m.Id, m.Time.String(), m.Select, m.Send)
	}

}

func InsertOper(msg Msg) error {
	conn, err := connect()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ctx := context.Background()

	query := "INSERT  into wss.ws_show_t1(id,select_op,send_value) values (?,?,?)"
	//stmt, err := conn
	err = conn.Exec(ctx, query, msg.Id, msg.Select, msg.Send)

	//batch 批量处理插入
	//batch, err := conn.PrepareBatch(ctx, "INSERT INTO wss.ws_show_t1")
	//if err != nil {
	//	return err
	//}
	//// we can insert JSON as either a string, struct or map
	////col2Data := `{"id": 11 , "select_op": "copy", "send_value": "pasword"}`
	//col3Data := map[string]interface{}{
	//	"id":         msg.Id,
	//	"select_op":  msg.Select,
	//	"send_value": msg.Send,
	//}
	//col1Data := Msg{
	//	Id:     msg.Id,
	//	Select: msg.Select,
	//	Send:   msg.Send,
	//}
	//// both named and unnamed can be added with slices
	//if err := batch.Append(col1Data, col3Data); err != nil {
	//	//fmt.Println("Append ERR")
	//	return err
	//}
	//
	//if err = batch.Send(); err != nil {
	//	return err
	//}

	return err

}

//ft: id=2  or  select_op="readWrite"  arg: ft string
func SelectOper(ft string) ([]Msg, error) {
	conn, err := connect()
	if err != nil {
		panic((err))
	}

	query := fmt.Sprintf("SELECT * FROM wss.ws_show_t1 WHERE %s", ft)
	//query := fmt.Sprintf("SELECT * FROM wss.ws_show_t1 ")
	//fmt.Println(query)
	ctx := context.Background()
	rows, err := conn.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var msgs []Msg
	for rows.Next() {
		var msg Msg
		err = rows.Scan(&msg.Id, &msg.Time, &msg.Select, &msg.Send)
		if err != nil {
			return nil, err
		}
		fmt.Printf(msg.Time.String())
		msgs = append(msgs, msg)
	}
	return msgs, rows.Err()

}

//  clickhouse.Open  <--- driver.Conn,
func connect() (driver.Conn, error) {
	ConfigInit()
	//conn := clickhouse.OpenDB(&clickhouse.Options{
	//	Addr: conf.Cfg.Ckconfig.Host,
	//	Auth: clickhouse.Auth{
	//		Database: conf.Cfg.Ckconfig.DatabaseName,
	//		Username: conf.Cfg.Ckconfig.Username,
	//		Password: conf.Cfg.Ckconfig.Password,
	//	},
	//	Settings: clickhouse.Settings{
	//		"max_execution_time": 60,
	//	},
	//	DialTimeout: 5 * time.Second,
	//	Compression: &clickhouse.Compression{
	//		clickhouse.CompressionZSTD,
	//		conf.Cfg.Ckconfig.CompressionLevel,
	//	},
	//	Debug: false,
	//})
	//
	//fmt.Println(conf.Cfg.Ckconfig.Host)
	//if err := conn.Ping(); err != nil {
	//	if exception, ok := err.(*clickhouse.Exception); ok {
	//		fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
	//	}
	//	return nil, err
	//}
	//return conn
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: conf.Cfg.Ckconfig.Host,
			Auth: clickhouse.Auth{
				Database: conf.Cfg.Ckconfig.DatabaseName,
				Username: conf.Cfg.Ckconfig.Username,
				Password: conf.Cfg.Ckconfig.Password,
			},
			ClientInfo: clickhouse.ClientInfo{
				Products: []struct {
					Name    string
					Version string
				}{
					{Name: "an-example-go-client", Version: "0.1"},
				},
			},

			Debugf: func(format string, v ...interface{}) {
				fmt.Printf(format, v)
			},
			//TLS: &tls.Config{
			//	InsecureSkipVerify: true,
			//},
		})
	)
	if err != nil {
		return nil, err
	}
	//// 测试连接是否可用
	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	return conn, nil
}
