
// inisetting 基本config設定套用，其餘可參考https://cloud.tencent.com/developer/article/1066126
package inisetting

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string
)

func init() {
	var err error
	Cfg, err = ini.Load("inisetting/conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'inisetting/conf/app.ini': %v", err)
	}

}
