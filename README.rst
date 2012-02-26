GoLang thrift-hive server client
===========================================

The hive thrift for go interface.  The default thrift server it targets is the cdh3u3, if you want a different version don't use go get.

Install default cdh3u3 server version:
    
    go get github.com/araddon/hive

Install different version:
    
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


Usage
===========

See examples/example.go



To create the thrift generated files from scratch
----------------------------------------------------
requires extensive hand editing afterwords to get to work

run these steps:

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

