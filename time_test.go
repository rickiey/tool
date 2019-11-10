package tool

import (
	"encoding/json"
	"fmt"

	"testing"
	"time"

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

}

func Test_NTP(t *testing.T) {
	fmt.Println(GetTimeFromNTP())
}
