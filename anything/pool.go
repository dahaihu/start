package anything

import (
	"context"
	"database/sql"
)

/**
* @Author: 胡大海
* @Date: 2020-03-14 17:58
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */


func GetClient() {
	db, err := sql.Open("mysql", "sakila")
	if err != nil {
		return
	}
	db.Conn(context.Background())
}
