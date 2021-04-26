https://www.jianshu.com/p/77ec3fc18508

mkdir -p  redis/sentinel/master
mkdir -p  redis/sentinel/slave

cp redis.conf redis/sentinel/master/redis-8000.conf
cp redis.conf redis/sentinel/slave/redis-8001.conf
cp redis.conf redis/sentinel/slave/redis-8002.conf

vi  redis/sentinel/master/redis-8000.conf
port 8000
daemonize no

vi redis/sentinel/slave/redis-8001.conf
vi redis/sentinel/slave/redis-8002.conf
port 8001
damonnize yes
slaveof 127.0.0.1 8000


redis-server redis/sentinel/master/redis-8000.conf&
% redis-cli -p 8000 ping
PONG

redis-server redis/sentinel/slave/redis-8001.conf
redis-server redis/sentinel/slave/redis-8002.conf

通过 info 命令看看是否成功:
 % redis-cli -p 8000 info replication
# Replication
role:master
connected_slaves:2
slave0:ip=127.0.0.1,port=8001,state=online,offset=42,lag=1
slave1:ip=127.0.0.1,port=8002,state=online,offset=42,lag=1
master_failover_state:no-failover
master_replid:f9e250c407d5af6101f37e9995c91508d99c4671
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:42
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:42


cp redis-sentinel.conf redis/sentinel/redis-sentinel-26379.conf
cp redis-sentinel.conf redis/sentinel/redis-sentinel-26380.conf
cp redis-sentinel.conf redis/sentinel/redis-sentinel-26381.conf

然后修改配置文件：
vi redis/sentinel/redis-sentinel-26379.conf
port 26379
daemonize yes
sentinel monitor mymaster 127.0.0.1 8000 2

关键的是：sentinel monitor mymaster 127.0.0.1 8000 2,这个配置表示该哨兵节点需要监控 8000 这个主节点， 2 代表着判断主节点失败至少需要 2 个 Sentinel 节点同意。

redis-sentinel redis/sentinel/redis-sentinel-26379.conf
redis-sentinel redis/sentinel/redis-sentinel-26380.conf
redis-sentinel redis/sentinel/redis-sentinel-26381.conf

看看哨兵的相关信息：
redis-cli -p 26380 info sentinel
# Sentinel
sentinel_masters:1
sentinel_tilt:0
sentinel_running_scripts:0
sentinel_scripts_queue_length:0
sentinel_simulate_failure_flags:0
master0:name=mymaster,status=ok,address=127.0.0.1:8000,slaves=2,sentinels=3

看到有个 master 节点，名称是 mymaster，地址是我们刚刚配置的 8000， 从节点有 2 个，哨兵有 3 个。

% redis-cli -p 8000 set "hello" "world"
OK
