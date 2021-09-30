# filecoin-reward-syner
sync reward data from kv db to mongodb

# 设置分片
```
use admin ;
db.runCommand({enablesharding:"filecoin"});

db.runCommand({shardcollection:"filecoin.gas",key:{height:1}});

db.runCommand({shardcollection:"filecoin.transfer",key:{height:1}});
```

# 开始同步

```
./filecoin-reward-syner mongo > runoob.log 2>&1 &
```