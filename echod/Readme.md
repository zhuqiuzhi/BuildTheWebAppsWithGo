# Echo Server 

# 编译 & 运行

```shell
go build
```

# 效果

1. 运行 echod 

```
2017/10/22 17:32:21 Echo Server is listen: localhost:8000
```

2. 使用nc 和 echo Server 建立 TCP 连接，并输入字符串

```shell
nc localhost 8000
hello?
	hello?
Is there anybody there?
	Is there anybody there?
```

3. echo server 接收到字符串后，会等待 1s,然后返回 Tab 和 客户端传过来的字符串 
