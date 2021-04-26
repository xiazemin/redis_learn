https://github.com/go-redis/redis

https://github.com/luin/ioredis


brew install redis

brew services start redis

brew services restart redis

% go run exp1/main.go
key value
key2 does not exist


% go run exp2/main.go 
PONG <nil>


redigo
https://segmentfault.com/a/1190000017879129

go-redis

% go run exp5/sentinel.go 
hello world
client: Redis<FailoverClient db:2>

% go run exp6/single.go
PONG <nil>
feekey examples
feekey does not exist

% go run exp7/cluster.go
PONG <nil>
pool state init state: &{1 2 0 2 2 0}
pool state final state: &{1999 4 0 4 4 0}

% go run exp9/rw.go
flush err: <nil>
[127.0.0.1 8000]
Redis<127.0.0.1:8000 db:0> 8000


https://pkg.go.dev/github.com/go-redis/redis/v8#example-NewClusterClient--ManualSetup

1，NewClient returns a client to the Redis Server specified by Options.
2，NewFailoverClient returns a Redis client that uses Redis Sentinel for automatic failover. It's safe for concurrent use by multiple goroutines.
3，NewClusterClient returns a Redis Cluster client as described in http://redis.io/topics/cluster-spec.
4，NewFailoverClusterClient returns a client that supports routing read-only commands to a slave node.
5，NewSentinelClient SentinelClient is a client for a Redis Sentinel.
6，NewUniversalClient returns a new multi client. The type of the returned client depends on the following conditions:
6.1. If the MasterName option is specified, a sentinel-backed FailoverClient is returned. 6.2. if the number of Addrs is two or more, a ClusterClient is returned. 3. Otherwise, a single-node Client is returned.


go-redis 有一个很大的不足就是：在 sentinel 部署模式下，它默认总是获取主库连接，因此在高并发尤其是读多写少的场景下并不适用。

cluster 模式也有两个明显的不足：
1，cluster 不能保证数据的强一致性。
2，cluster 不支持处理多个 key。
为了保证多个命令的原子性，我们通常会使用 lua 脚本来实现。但 redis 要求单个 lua 脚本的所有 key 在同一个槽位。

一个 sentinel 集群可以监控多个 master-slave 集群，每个 master-slave 集群有一个主节点和多个从节点，一旦 sentinel 监控到某个 master 宕机，会自动从它的多个 slave 节点中选出最合适的一个作为新的 master 节点。

https://www.jianshu.com/p/a26754b8131f
https://help.aliyun.com/document_detail/65001.html
https://zq99299.github.io/note-book/cache-pdp/redis/029.html

https://cloudnative.to/blog/redis-cluster-with-istio/#redis-%E8%AF%BB%E5%86%99%E5%88%86%E7%A6%BB


redis普通主从模式
通过持久化功能，Redis保证了即使在服务器重启的情况下也不会损失（或少量损失）数据，因为持久化会把内存中数据保存到硬盘上，重启会从硬盘上加载数据。 。但是由于数据是存储在一台服务器上的，如果这台服务器出现硬盘故障等问题，也会导致数据丢失。为了避免单点故障，通常的做法是将数据库复制多个副本以部署在不同的服务器上，这样即使有一台服务器出现故障，其他服务器依然可以继续提供服务。为此， Redis 提供了复制（replication）功能，可以实现当一台数据库中的数据更新后，自动将更新的数据同步到其他数据库上。

https://www.huaweicloud.com/articles/38a076804762433beaa573e3b8dd3aa4.html
https://gitee.com/websterlu/redisx