% redis-cli -p 7001 cluster nodes
8289fb1c92d985732e2bbd0713733a94541d7c36 127.0.0.1:7000@17000 master - 0 1619420360730 1 connected 0-5460
b9d1569cc7b5f1bbc4294be3e3e5f7f4cd67bfbc 127.0.0.1:7005@17005 slave 89682e0e2d89672f20ee5e1bb0af68c6c9c86da4 0 1619420361757 3 connected
0dce1bd487fd9353bfa347eb202243c79e8cad6a 127.0.0.1:7003@17003 slave 8289fb1c92d985732e2bbd0713733a94541d7c36 0 1619420359000 1 connected
1b13903926f2f60dcafd2b4e8e0de2835d7be9ed 127.0.0.1:7001@17001 myself,master - 0 1619420360000 2 connected 5461-10922
89682e0e2d89672f20ee5e1bb0af68c6c9c86da4 127.0.0.1:7002@17002 master - 0 1619420360000 3 connected 10923-16383
b07af22f38b3bed4cb17d98b82f9e0d62789265c 127.0.0.1:7004@17004 slave 1b13903926f2f60dcafd2b4e8e0de2835d7be9ed 0 1619420359000 2 connected


 % redis-cli -c -p 7004
127.0.0.1:7004> get test
-> Redirected to slot [6918] located at 127.0.0.1:7001
(nil)

% redis-cli -c -p 7001
127.0.0.1:7001> set test 1
OK


127.0.0.1:7001> get test
"1"

如果key 不存在，会重定向到master


 % redis-cli -c -p 7004
127.0.0.1:7004> get test
-> Redirected to slot [6918] located at 127.0.0.1:7001
"1"
127.0.0.1:7001>


By default, slave redirects the client to its master, since data on slave might be stale, i.e. the write to master might NOT be synced to slave.

However, sometimes we don't care about the staleness, and want to scale read operations by reading (possible stale) data from slave.

In order to achieve that, you can send the READONLY command to slave. Then any ready-only operations on that connection will be served by the slave, and won't be redirected to its master.

If you want to turn off the READONLY mode, you can send the READWRITE command to tell the slave to redirect read requests to master.

NOTE:

no matter whether the slave is in READONLY or READWRITE mode, you CANNOT write to slave, i.e. write operations are always redirected to master.

UPDATE: slave-serve-stale-data and slave-read-only configurations have nothing to do with the READONLY and READWRITE commands.

slave-serve-stale-data controls whether the slave should redirect the request to master or just returns an error, when it loses connection with master, or the replication is still in progress.

slave-read-only controls whether you can write to slave. However, these writes won't be propagated to master and other slaves, and will be removed after resync with the master.

https://stackoverflow.com/questions/49061445/when-redis-get-key-in-slave-why-redirect-to-master


https://redis.io/topics/cluster-spec#scaling-reads-using-slave-Nodes


127.0.0.1:7004> CONFIG SET slave-read-only yes
OK


127.0.0.1:7004> get test
-> Redirected to slot [6918] located at 127.0.0.1:7001
"1"

https://gnuhpc.gitbooks.io/redis-all-about/content/HAClusterArchPractice/ms/readonly.html


% redis-cli -c -p 7004
127.0.0.1:7004> set test 2
-> Redirected to slot [6918] located at 127.0.0.1:7001
OK


127.0.0.1:7001> config get slave-read-only
1) "slave-read-only"
2) "yes"

从节点默认不让读取，如果读取从节点，将会重定向到主节点。使用readonly命令，允许从节点提供读服务，如


redis-cli -c -p 7004
127.0.0.1:7004> readonly
OK
127.0.0.1:7004> get test
"2"
127.0.0.1:7004> get test1
-> Redirected to slot [4768] located at 127.0.0.1:7000
(nil)


https://www.jianshu.com/p/14aaa1c1cab6


readwrite

       取消（重置）readonly命令的设置，恢复salve节点默认状态

https://redis.io/topics/cluster-spec#scaling-reads-using-slave-nodes

http://www.redis.cn/commands/readonly.html



2. 开启对Cluster中Slave Node的访问
在一个负载比较高的Redis Cluster中，如果允许对slave节点进行读操作将极大的提高集群的吞吐能力。

开启对Slave 节点的访问，受以下3个参数的影响

type ClusterOptions struct {
    // Enables read-only commands on slave nodes.
    ReadOnly bool
    // Allows routing read-only commands to the closest master or slave node.
    // It automatically enables ReadOnly.
    RouteByLatency bool
    // Allows routing read-only commands to the random master or slave node.
    // It automatically enables ReadOnly.
    RouteRandomly bool
    ... 
}


如果ReadOnly = true，只选择Slave Node
如果ReadOnly = true 且 RouteByLatency = true 将从slot对应的Master Node 和 Slave Node选择，选择策略为: 选择PING 延迟最低的节点
如果ReadOnly = true 且 RouteRandomly = true 将从slot对应的Master Node 和 Slave Node选择，选择策略为:随机选择

http://vearne.cc/archives/1113

