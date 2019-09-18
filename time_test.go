package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func Test_time(t *testing.T) {
	timestr := `{"ta":"2019-03-04", "tb":"2019-06-06 15:32:23", "tc":"", "td":null, "te":"2019/02/03", "tf":"2020-12-16 06:20:00 +08:00", "th":"2020-04-26 13:13:43 +0800 CST"}`

	type TimeSet struct {
		T1 Time `json:"ta"`
		T2 Time `json:"tb"`
		T3 Time `json:"tc"`
		T4 Time `json:"td"`
		T5 Time `json:"te"`
		T6 Time `json:"tf"`
		T7 Time `json:"tg"`
		T8 Time `json:"th"`
	}
	ts := TimeSet{}
	err := json.Unmarshal([]byte(timestr), &ts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ts)
	if time.Time(ts.T3).IsZero() {
		fmt.Println("this is zero")
	}
}

func Test_time02(t *testing.T) {

	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbUser := "root"
	dbPass := "951225"
	dbName := "test"
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Shanghai")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	log.Println(dsn)
	testOrm, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Println("open db: ", err)
		os.Exit(1)
	}
	defer testOrm.Close()
	tm := time.Now().UTC()
	fmt.Println(tm)
	dt := Time(tm)
	_ = dt
	sql_intert := "insert into new_table values(?,?,?,?)"
	err = testOrm.Debug().Exec(sql_intert, "q", dt, dt, tm).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	type NewTable struct {
		Content  string `gorm:"cloumn:content"`
		TestTime Time   `gorm:"cloumn:test_time"`
		DateFmt  Time   `gorm:"cloumn:date_fmt"`
		DateTime Time   `gorm:"cloumn:date_time"`
	}
	res := []NewTable{}
	querysql := "select * from new_table "
	err = testOrm.Debug().Raw(querysql).Scan(&res).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range res {
		fmt.Println(v)
	}

}

func Test_NTP(t *testing.T) {
	fmt.Println(GetTimeFromNTP())
}
