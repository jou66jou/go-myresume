package sqltojson

import (
	"database/sql"
	"fmt"

	"github.com/bdwilliams/go-jsonify/jsonify"
	_ "github.com/go-sql-driver/mysql"
)

func test() {
	con, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", "test", "test", "127.0.0.1", "ruokdraw"))
	if err != nil {
		panic(err.Error())
	}

	rows, err := con.Query("select * from lottery_draw_rule;")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	defer con.Close()

	fmt.Println(jsonify.Jsonify(rows))
}
