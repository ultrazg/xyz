### 正在收听的人数

根据 eid 查询单集当前时间正在收听的人数。可轮询该接口

#### 请求地址

> /episode_live_count

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数 | 必填 | 类型   | 说明    |
| :--- | :--- | :----- | ------- |
| eid  | true | string | 单集 id |

#### 返回字段

| 返回字段           | 类型   | 说明           |
| :----------------- | :----- | :------------- |
| audiencesCountText | string | 正在收听的人数 |


#### 示例

> 地址：https://www.example.com/episode_live_count

参数

```javascript
{
  "eid": "662b664a8a089719b7f6bbd3"
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: {
      audiencesCountText: "3",
    },
  },
  msg: "OK",
}
```