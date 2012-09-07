GoLang thrift-hive server client
===========================================

The hive thrift for go interface.  The default thrift server it targets is the cdh3u3 (although teh cdh3u1, u2 appear to be the same, check the branches.)

Install default cdh3 server version::

    go get github.com/araddon/thrift4go/lib/go/thrift
    go get github.com/araddon/hive


Usage
===========

See examples/example.go::
    
    import "github.com/araddon/hive"
    
    func main() {
      
      // checkout a connection
      conn, err := hive.GetHiveConn()
      if err == nil {

        //_, _ = conn.Client.Execute("CREATE TABLE rrr(a STRING, b INT, c DOUBLE);")
        er, err := conn.Client.Execute("SELECT * FROM logevent")
        if er == nil && err == nil {
          for {
            row, _, _ := conn.Client.FetchOne()
            if len(row) > 0 {
              log.Println("row ", row)
            } else {
              break
            }
          }
        }
      }
      if conn != nil {
        // make sure to check connection back into pool
        conn.Checkin()
      }
    }




To create the thrift generated files from scratch
----------------------------------------------------
requires extensive hand editing afterwords to get to work

run these steps::

    thrift --gen go -out . fb303.thrift 
    cd fb303
    go install
    cd ..
    thrift --gen go -out . hive_metastore.thrift 
    cd hive_metastore
    go install
    cd ..
    thrift --gen go -out . queryplan.thrift 
    cd queryplan
    go install
    cd ..
    thrift --gen go -out . hive_service.thrift 
    mv hive_service/ttypes.go .
    mv hive_service/ThriftHive.go .
    go install


Issues with the go thrift compiler
--------------------------------------
need to submit issues...

	* for "include" thrift files, the objects within those includes are not referenced correctly, ie:
	    Was:
			GetQueryPlan() (retval25 *QueryPlan, ex *HiveServerException, err error)
        should be:
            GetQueryPlan() (retval25 *queryplan.QueryPlan, ex *HiveServerException, err error)



Install different version::
    
    # in a gopath src directory
    mkdir -p github.com/araddon
    cd github.com/araddon
    git clone git://github.com/araddon/hive.git
    cd hive
    git checkout cdh3u2  # etc, see list of branches

    # it is made up of 5 sub-packages 
    cd thriftlib
    cd fb303
    go install
    cd ../hive_metastore
    go install
    cd ../queryplan
    go install
    cd ..
    go install
    cd ..
    go install
