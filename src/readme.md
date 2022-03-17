### 1、tcp抓包【eth1文件夹 】

```
go get github.com/google/gopacket
```

`win` 需要安装 `npcap`

`linux` 安装 `yum install libpcap libpcap-devel`

`httpdemo`文件夹内的`http`程序运行

文件夹下的`main.go`去监听`http`程序

随意请求，可监听到内容

### 2、iptables手动设置转发docker

随意写一个监听`9090`端口的服务，在容器内运行，不要设置任何`-p`参数

```
# 清空所有已定规则(flush)
$ iptables -F 
# 删除所有用户链
$ iptables -X 
```

```
查看转发
$ iptables -t nat -nvL
```

两句命令给容器手动模拟端口映射

```
# 设置转发
# 9090为宿主机地址 
# 172.17.0.2:9090 代表默认自动生成的容器地址和容器内的端口 (docker inspect * 查看容器地址)
$ iptables -t nat -A PREROUTING -p tcp --dport 9090 -j DNAT --to-destination 172.17.0.2:9090
# 设置链
# -i eth1 -o docker0 表示从 eth1网卡入的流量 出口到 docker0网卡
$ iptables -A FORWARD -i eth1 -o docker0 -j ACCEPT
```

用任意方式访问宿主机`9090`端口，会被代理至容器内

### 3、不同 linux namespace 通信

#### 网络
```
# 列表
$ ip netns list
# 新增
ip netns add [namespace]
# 在对应命名空间执行命令
$ ip netns exec [namespace] [command]
```

例，互通

```
# 创建两个命名空间
$ ip netns add ns1 && ip netns add ns2
# 创建两个互通的veth (自定义名称 veth001 与 veth002，peer代表一对)
$ ip link add name veth001 type veth peer name veth002
# 将两个veth移动到命名空间内
$ ip link set veth001 netns ns1 && ip link set veth002 netns ns2
# 给移动到命名空间后的veth设置ip
$ ip netns exec ns1 ip addr add local 10.12.0.2/24 dev veth001
$ ip netns exec ns2 ip addr add local 10.12.0.3/24 dev veth002
# 启动veth
$ ip netns exec ns1 ip link set veth001 up
$ ip netns exec ns2 ip link set veth002 up
```