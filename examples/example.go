package main

import (
	hive "github.com/araddon/hive"
	"log"
)

func init() {
	hive.MakePool("lio25:10000")
}

func main() {

	// checkout a connection
	conn, err := hive.GetHiveConn()
	if err == nil {

		_, _ = conn.Client.Execute("CREATE TABLE rrr(a STRING, b INT, c DOUBLE);")
		er, err := conn.Client.Execute("SELECT * FROM logevent_stg LIMIT 20")
		if er == nil && err == nil {
			for {
				row, _, _ := conn.Client.FetchOne()
				if len(row) > 0 {
					log.Println("row ", row)
				} else {
					return
				}
			}
		} else {
			log.Println(er, err)
		}
	}
	if conn != nil {
		// make sure to check connection back into pool
		conn.Checkin()
	}
}
