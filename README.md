<h1 align="center">
  <img src="./logo.png" alt="xyz" width="100">
  <br>xyz<br>
</h1>
<h4 align="center">小宇宙FM API</h4>


## 环境

Go 1.22.0 <img alt="Static Badge" src="https://img.shields.io/badge/Go-1.22.0-blue.svg">




## 安装

```shell
$ git clone git@github.com:ultrazg/xyz.git
$ cd xyz/cmd
$ go mod tidy
```



## 运行

```shell
$ go run .
```

服务端启动默认端口为 `8080`，若想使用其他端口，可执行以下命令：

```shell
$ go run . -p 3000
```

服务启动时打开文档：

```shell
$ go run . -d
```

> 接口地址：http://localhost:{{port}}/login
>
> 文档地址：http://localhost:{{port}}/doc



## 构建

项目内提供对应平台的 `build.sh` 文件，按需执行即可



## 功能（更新中）

1. 发送验证码
2. 短信登录
3. 刷新token
4. 搜索
5. 「你可能想搜的内容」
6. 获取我的信息
7. 获取节目、单集等内容
8. 获取我的订阅
9. 查询节目列表
10. 获取播客音频链接
11. 查询单集详情
12. 查询节目详情
13. 相关节目推荐
14. 查询我的贴纸
15. 展示我的贴纸墙
16. 查询单集播放进度
17. 查询单集评论
18. 查询评论回复



## License

[The MIT License](https://github.com/ultrazg/xyz/blob/dev/LICENSE)
