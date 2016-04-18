# Schema全局同步变更

## Proxy读写锁
### 库锁或表锁
锁记录在etcd中有一份存根，每一个Proxy节点有一个统一的锁记录管理中心，
用于控制所有锁的状态，并提供gROC的调用，库锁是大于表锁的。若要获得库锁，
要等所有Proxy完成对此库中所有的资源完成，同理表锁也是如此

## masters pool
此池存有所有的在正处理schema变更的master节点。所有proxy 节点都应该watch此池中的master节点。                 
三种case：                 
### case 1: 
update master节点自己更新存活期，忽视            
### case 2: 
delete master节点自己之行完所有的DDL删除，忽视               
### case 3: 
timeout master节点自己崩溃，所有watch节点应该去竞选成为master；
master选举产生之后，从etcd拿取未完成的DDL任务并接着执行                     

## 任务提交master
DDL任务提交至Proxy，Proxy将任务送至此DDL所属的库或表的master Proxy节点                            
### case 1: 
集群中已经有此库或表的master节点，那么gRPC传递DDL语句至master 节点，
若提交至master节点的时候master拒绝接受，那么等待(一次TTL周期)并检查master节点，
若master节点健康则重试，若master节点崩溃，那么进入case 2                   
### case 2: 
集群中没有此库或者表的master节点，那么此Proxy开始竞选master节点, 
若其他节点获得master那么执行和case 1一样的操作；若竞选成功那把此节点加入到masters pool中

## master对DDL进行处理

### 检查未完成的DDL任务
成为master节点之后，从etcd获取当前master所cover的DDL范围内的未完成任务，
如果有未完成的任务，那么提出所有的任务，加入到本机的DDL任务队列中
### gRPC接收DDL任务
gRPC传来的任务，生成执行计划，然后写入到etcd中进行存根，然后加入本机的DDL任务对列中
### DDL任务对列
若队列里面有数据，那么进行上锁：         
- 改变etcd状态进行上锁
- 获取所有Proxy节点，广播进行此表或库的上锁
- 依次拿出对列中的任务，根据执行计划执行，完成一步之后就记录
- 改变etcd状态进行解锁
- 获取所有的Proxy节点，广播进行对此表或者库解锁

## 关于死锁的思考🔒
当一个新的Proxy P进入集群，A正被锁住，P通过etcd store获取A的情况，但此时遇到TCP延迟或者丢包率增加，
在P还得到etcd response之前，A被释放掉，P收到了A Unlock的gRPC调用，由于A在P上是为锁定的，所以忽略了
本次调用。那么，P上A最终的状态是锁定的，由于旧的脏数据引起的。
### 解决方案
* 使用版本控制
问题1：版本信息应该保存多久？
设置一个过期时间，etcd应该设置Response timeout, 时间应为timeout的两倍.
问题2：版本信息应该如何生成？
版本控制由两阶段Trsaction ID + seq id组成. transaction ID是uuid，seq id由Master生成                        
上锁者只能是Master的节点，Master节点在一时间点只能存在一个 
* 询问上锁节点
会造成过多的HTTP请求，上锁节点可能崩溃，过于复杂的逻辑