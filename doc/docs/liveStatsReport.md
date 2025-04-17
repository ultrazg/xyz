### 上报播放状态

将当前正在播放的单集状态上报，用于更新单集「正在收听的人数」

#### 请求地址

> /live_stats_report

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数 | 必填 | 类型   | 说明     |
| :--- | :--- | :----- | -------- |
| eid  | true | string | 单集 id  |
| pid  | true | string | 单集 pid |

#### 返回字段

无




#### 示例

> 地址：https://www.example.com/live_stats_report

参数

```javascript
{
  "eid": "61c2f817874a732463bfb04a",
  "pid": "5e280fa7418a84a0461f912b"
}
```

响应

```javascript
{
  code: 200,
  data: {},
  msg: "OK",
}
```