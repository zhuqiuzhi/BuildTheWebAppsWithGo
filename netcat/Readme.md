# netcat

# 编译 & 运行

```shell
go build
```

# 效果

1. 启动一个 echod 服务器

```shell
./echod
2017/10/27 09:59:12 Echo Server is listen: localhost:8000
```

2. 运行 netcat, 输入 Hello，回车，如果要结束对话，按 Ctrl 和 d

```shell
./netcat
Hello
	 HELLO
	 Hello
	 hello
2017/10/27 10:00:32 This is a TCP connect, close write
2017/10/27 10:00:32 Done
```


