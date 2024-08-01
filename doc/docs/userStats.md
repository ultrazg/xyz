### 用户统计数据

根据 uid 查询用户的粉丝数、关注数、订阅数和总收听时长

#### 请求地址

> /user_stats

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
| uid  | true | string | 用户 uid |

#### 返回字段

| 返回字段           | 类型   | 说明       |
| :----------------- | :----- | :--------- |
| followerCount      | number | 粉丝数     |
| followingCount     | number | 关注数     |
| subscriptionCount  | number | 订阅数     |
| totalPlayedSeconds | number | 播放总时长 |


#### 示例

> 地址：https://www.example.com/user_stats

参数

```javascript
{
  "uid":"UID"
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: {
      followerCount: 0,
      followingCount: 1,
      subscriptionCount: 5,
      totalPlayedSeconds: 505379,
    },
  },
  msg: "OK",
}
```