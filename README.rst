To Generate the hive client
===========================================


    
    # cd to folder.  check out the branch you want
    git checkout 0.7.1-cdh3u3

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
    cd hive_service
    go install
