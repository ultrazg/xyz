<h1 align="center">
  <img src="./logo.png" alt="xyz" width="100">
  <br>xyz<br>
</h1>
<h4 align="center">小宇宙FM API</h4>

## 免责声明

**⚠️ 本项目仅供学习、研究使用，请遵守国家法律，严禁用于任何非法用途**

## 环境

Go 1.22.0 <img alt="Static Badge" src="https://img.shields.io/badge/Go-1.22.0-blue.svg">

## 安装

```shell
$ git clone git@github.com:ultrazg/xyz.git
$ cd xyz
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
> 文档地址：http://localhost:{{port}}/docs

> 可在 [Releases](https://github.com/ultrazg/xyz/releases) 下载编译好的可执行文件

## 作为模块

```shell
go get github.com/ultrazg/xyz
```

```go
package main

import (
	"fmt"

	"github.com/ultrazg/xyz/service"
)

func main() {
	err := service.Start()
	if err != nil {
		fmt.Println("fail")
	}
}
```

## 构建

项目内提供对应平台的 `build.sh` 文件，按需执行即可

## 功能

- [x] 发送验证码
- [x] 短信登录
- [x] 刷新 token
- [x] 搜索节目、单集和用户
- [x] 「你可能想搜的内容」
- [x] 获取我的信息
- [x] 获取节目、单集等内容
- [x] 获取「我的订阅」
- [x] 订阅/取消订阅节目
- [x] 查询节目列表
- [x] 查询节目内「最受欢迎」的单集列表
- [x] 查询节目公告、荣誉墙、主体等信息
- [x] 获取播客音频链接
- [x] 查询单集详情
- [x] 查询节目详情
- [x] 相关节目推荐
- [x] 查询「我的贴纸」
- [x] 展示「我的贴纸墙」
- [x] 查询/更新单集播放进度
- [x] 查询单集评论
- [x] 查询评论回复
- [x] 获取榜单、精选节目、推荐等
- [x] 正在收听的人数
- [x] 精彩时间点
- [x] 创建精彩时间点
- [x] 订阅列表更新
- [x] 获取分类、分类标签以及查询分类内容
- [x] 星标订阅管理
- [x] 收藏单集、评论
- [x] 查询「我的收藏」
- [x] 收听历史
- [x] 未读消息
- [x] 查询用户信息和用户统计数据
- [x] 刷新「大家都在听」推荐
- [x] 查询收听数据
- [x] 查询「个人主页」收听历史记录
- [x] 查询「用户的喜欢」
- [x] 查询用户创建的播客节目
- [x] 查询首页榜单（最热榜、锋芒榜和新星榜）
- [x] 查询关注与被关注列表
- [x] 点赞/取消点赞评论
- [x] 获取黑名单列表
- [x] 拉黑/取消拉黑用户
- [x] 获取用户偏好设置
- [x] 更新用户偏好设置
- [x] 关注/取关用户
- [ ] ...

## License

[The MIT License](https://github.com/ultrazg/xyz/blob/dev/LICENSE)
