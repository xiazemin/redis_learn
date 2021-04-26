
https://zhuanlan.zhihu.com/p/59172042

https://www.jianshu.com/p/a32542ce4c0b


 /usr/local/etc/redis.conf

mkdir -p redis/cluster/7000
mkdir -p redis/cluster/7001
mkdir -p redis/cluster/7002
mkdir -p redis/cluster/7003
mkdir -p redis/cluster/7004
mkdir -p redis/cluster/7005

cp redis.conf redis/cluster/7000

vi  redis/cluster/7000/redis.conf

port 7000                                     # Redis 节点的端口号
cluster-enabled yes                           # 实例以集群模式运行
cluster-config-file nodes-7000.conf           # 节点配置文件路径
cluster-node-timeout 5000                     # 节点间通信的超时时间
appendonly yes                                # 数据持久化


cp redis/cluster/7000/redis.conf redis/cluster/7001/redis.conf
cp redis/cluster/7000/redis.conf redis/cluster/7002/redis.conf
cp redis/cluster/7000/redis.conf redis/cluster/7003/redis.conf
cp redis/cluster/7000/redis.conf redis/cluster/7004/redis.conf
cp redis/cluster/7000/redis.conf redis/cluster/7005/redis.conf

redis/cluster/7001/redis.conf
vi  redis/cluster/7002/redis.conf

redis-server redis/cluster/7000/redis.conf &
redis-server redis/cluster/7001/redis.conf &
redis-server redis/cluster/7002/redis.conf &
redis-server redis/cluster/7003/redis.conf &
redis-server redis/cluster/7004/redis.conf &
redis-server redis/cluster/7005/redis.conf &

 % ps -ef |grep redis
  501  4461     1   0  9:19下午 ??         0:01.20 /opt/homebrew/opt/redis/bin/redis-server 127.0.0.1:6379
  501  6286   464   0 10:22上午 ttys001    0:00.25 redis-server 127.0.0.1:7000 [cluster]
  501  6287   464   0 10:22上午 ttys001    0:00.25 redis-server 127.0.0.1:7001 [cluster]
  501  6288   464   0 10:22上午 ttys001    0:00.25 redis-server 127.0.0.1:7002 [cluster]
  501  6493   464   0 10:27上午 ttys001    0:00.01 redis-server 127.0.0.1:7003 [cluster]
  501  6494   464   0 10:27上午 ttys001    0:00.01 redis-server 127.0.0.1:7004 [cluster]
  501  6495   464   0 10:27上午 ttys001    0:00.01 redis-server 127.0.0.1:7005 [cluster]

http://download.redis.io/redis-stable/src/redis-trib.rb

chmod +x redis-trib.rb

# 无需指定哪个节点为 master，哪个节点为 slave，因为 redis 内部算法已经帮我们实现了
# 使用 –replicas 1 创建集群，即每个 master 带一个 slave
./redis-trib.rb create --replicas 1 127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 127.0.0.1:7004 127.0.0.1:7005
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin20/rbconfig.rb:229: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
WARNING: redis-trib.rb is not longer available!
You should use redis-cli instead.

All commands and features belonging to redis-trib.rb have been moved
to redis-cli.

https://stackoverflow.com/questions/52880996/redis-trib-rb-is-not-longer-available-but-redis-cli-cluster-create-throws-unre/52887311

redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 127.0.0.1:7004 127.0.0.1:7005 --cluster-replicas 1

>>> Performing hash slots allocation on 6 nodes...
Master[0] -> Slots 0 - 5460
Master[1] -> Slots 5461 - 10922
Master[2] -> Slots 10923 - 16383
Adding replica 127.0.0.1:7004 to 127.0.0.1:7000
Adding replica 127.0.0.1:7005 to 127.0.0.1:7001
Adding replica 127.0.0.1:7003 to 127.0.0.1:7002
>>> Trying to optimize slaves allocation for anti-affinity
[WARNING] Some slaves are in the same host as their master
M: 8289fb1c92d985732e2bbd0713733a94541d7c36 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
M: 1b13903926f2f60dcafd2b4e8e0de2835d7be9ed 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
M: 89682e0e2d89672f20ee5e1bb0af68c6c9c86da4 127.0.0.1:7002
   slots:[10923-16383] (5461 slots) master
S: 0dce1bd487fd9353bfa347eb202243c79e8cad6a 127.0.0.1:7003
   replicates 8289fb1c92d985732e2bbd0713733a94541d7c36
S: b07af22f38b3bed4cb17d98b82f9e0d62789265c 127.0.0.1:7004
   replicates 1b13903926f2f60dcafd2b4e8e0de2835d7be9ed
S: b9d1569cc7b5f1bbc4294be3e3e5f7f4cd67bfbc 127.0.0.1:7005
   replicates 89682e0e2d89672f20ee5e1bb0af68c6c9c86da4
Can I set the above configuration? (type 'yes' to accept): yes
>>> Nodes configuration updated
>>> Assign a different config epoch to each node
>>> Sending CLUSTER MEET messages to join the cluster
Waiting for the cluster to join
....
>>> Performing Cluster Check (using node 127.0.0.1:7000)
M: 8289fb1c92d985732e2bbd0713733a94541d7c36 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
S: 0dce1bd487fd9353bfa347eb202243c79e8cad6a 127.0.0.1:7003
   slots: (0 slots) slave
   replicates 8289fb1c92d985732e2bbd0713733a94541d7c36
S: b07af22f38b3bed4cb17d98b82f9e0d62789265c 127.0.0.1:7004
   slots: (0 slots) slave
   replicates 1b13903926f2f60dcafd2b4e8e0de2835d7be9ed
M: 1b13903926f2f60dcafd2b4e8e0de2835d7be9ed 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
M: 89682e0e2d89672f20ee5e1bb0af68c6c9c86da4 127.0.0.1:7002
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
S: b9d1569cc7b5f1bbc4294be3e3e5f7f4cd67bfbc 127.0.0.1:7005
   slots: (0 slots) slave
   replicates 89682e0e2d89672f20ee5e1bb0af68c6c9c86da4
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.


检查集群的状态
redis-cli --cluster check 127.0.0.1:7000
127.0.0.1:7000 (8289fb1c...) -> 0 keys | 5461 slots | 1 slaves.
127.0.0.1:7001 (1b139039...) -> 0 keys | 5462 slots | 1 slaves.
127.0.0.1:7002 (89682e0e...) -> 0 keys | 5461 slots | 1 slaves.
[OK] 0 keys in 3 masters.
0.00 keys per slot on average.
>>> Performing Cluster Check (using node 127.0.0.1:7000)
M: 8289fb1c92d985732e2bbd0713733a94541d7c36 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
S: 0dce1bd487fd9353bfa347eb202243c79e8cad6a 127.0.0.1:7003
   slots: (0 slots) slave
   replicates 8289fb1c92d985732e2bbd0713733a94541d7c36
S: b07af22f38b3bed4cb17d98b82f9e0d62789265c 127.0.0.1:7004
   slots: (0 slots) slave
   replicates 1b13903926f2f60dcafd2b4e8e0de2835d7be9ed
M: 1b13903926f2f60dcafd2b4e8e0de2835d7be9ed 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
M: 89682e0e2d89672f20ee5e1bb0af68c6c9c86da4 127.0.0.1:7002
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
S: b9d1569cc7b5f1bbc4294be3e3e5f7f4cd67bfbc 127.0.0.1:7005
   slots: (0 slots) slave
   replicates 89682e0e2d89672f20ee5e1bb0af68c6c9c86da4
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.


查看集群的信息
redis-cli --cluster info 127.0.0.1:7000
127.0.0.1:7000 (8289fb1c...) -> 0 keys | 5461 slots | 1 slaves.
127.0.0.1:7001 (1b139039...) -> 0 keys | 5462 slots | 1 slaves.
127.0.0.1:7002 (89682e0e...) -> 0 keys | 5461 slots | 1 slaves.
[OK] 0 keys in 3 masters.
0.00 keys per slot on average.

登录任意一个节点，执行命令
% redis-cli -c -p 7000
127.0.0.1:7000> set name xiazemin
-> Redirected to slot [5798] located at 127.0.0.1:7001
OK
127.0.0.1:7001> get name
"xiazemin"

分配 slot
redis-cli -p 7000 cluster addslots {0..5461}
不分配的话会自动分配

命令验证：
 redis-cli -p 7000 cluster nodes
0dce1bd487fd9353bfa347eb202243c79e8cad6a 127.0.0.1:7003@17003 slave 8289fb1c92d985732e2bbd0713733a94541d7c36 0 1619404774196 1 connected
b07af22f38b3bed4cb17d98b82f9e0d62789265c 127.0.0.1:7004@17004 slave 1b13903926f2f60dcafd2b4e8e0de2835d7be9ed 0 1619404772147 2 connected
1b13903926f2f60dcafd2b4e8e0de2835d7be9ed 127.0.0.1:7001@17001 master - 0 1619404770083 2 connected 5461-10922
89682e0e2d89672f20ee5e1bb0af68c6c9c86da4 127.0.0.1:7002@17002 master - 0 1619404773175 3 connected 10923-16383
8289fb1c92d985732e2bbd0713733a94541d7c36 127.0.0.1:7000@17000 myself,master - 0 1619404773000 1 connected 0-5460
b9d1569cc7b5f1bbc4294be3e3e5f7f4cd67bfbc 127.0.0.1:7005@17005 slave 89682e0e2d89672f20ee5e1bb0af68c6c9c86da4 0 1619404771115 3 connected

从节点的 cli 命令窗口关联主节点
redis-cli -p 7003 cluster replicate 7000的NodeID

7002的NodeID 其实就是执行 redis-cli -p 7000 cluster nodes 命令出现的那一串 16 进制字符串