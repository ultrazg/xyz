### 标记精彩时间点

给单集播放时间线标记一个精彩时间点

#### 请求地址

> /episode_clap_create

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求参数

| 参数      | 必填 | 类型   | 说明               |
| :-------- | :--- | :----- | ------------------ |
| eid       | true | string | 单集 eid           |
| timestamp | true | number | 当前播放时间（秒） |
| duration  | true | number | 单集总时长         |

#### 返回字段

无


#### 示例

> 地址：https://www.example.com/episode_clap_create

参数

```javascript
{
  "duration": 6413,
  "timestamp": 3342,
  "eid": "662251f28b343a8e9aa4ffc7"
}
```

响应

``` javascript
{
	"code": 200,
	"msg": "OK",
	"data": {}
}
```