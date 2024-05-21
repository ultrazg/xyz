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

- [x] 发送验证码
- [x] 短信登录
- [x] 刷新token
- [x] 搜索
- [x] 「你可能想搜的内容」
- [x] 获取我的信息
- [x] 获取节目、单集等内容
- [x] 获取「我的订阅」
- [x] 查询节目列表
- [x] 获取播客音频链接
- [x] 查询单集详情
- [x] 查询节目详情
- [x] 相关节目推荐
- [x] 查询「我的贴纸」
- [x] 展示「我的贴纸墙」
- [x] 查询单集播放进度
- [x] 查询单集评论
- [x] 查询评论回复
- [x] 获取榜单、精选节目、推荐等
- [x] 正在收听的人数
- [x] 精彩时间点
- [ ] 创建精彩时间点
- [ ] 订阅列表更新
- [ ] 获取分类、分类筛选条件以及查询分类内容
- [ ] 星标订阅管理
- [ ] 收藏单集、评论
- [ ] 查询「我的收藏」
- [ ] 收听历史
- [ ] 「我的通知」



## License

[The MIT License](https://github.com/ultrazg/xyz/blob/dev/LICENSE)
