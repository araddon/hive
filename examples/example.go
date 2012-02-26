package main

import (
  "log"
  hive "github.com/araddon/hive"
)

func init() {
  hive.MakePool("192.168.1.17:10000")
}


func main() {
  
  // checkout a connection
  conn, err := hive.GetHiveConn("hive")
  if err == nil{

    //_, _ = conn.Client.Execute("CREATE TABLE rrr(a STRING, b INT, c DOUBLE);")
    er, err := conn.Client.Execute("SELECT * FROM logevent")
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














