# Time Server 

# 编译 & 运行

```shell
go build
```

# 效果

1. 运行 timeserver

```
2017/10/21 20:16:47 Time Server is listen: localhost:8000
```

2. 使用nc 和 Time Server 建立 TCP 连接 

```shell
nc localhost 8000
```

3. Time Server 会打印出客户端的地址,并返回当前时间,然后断开 TCP 连接

4. 当 Time Server 接收到中断信号、结束信号、KILL信号, 会关闭 TCP 连接,并打印出接收到的操作系统信号
