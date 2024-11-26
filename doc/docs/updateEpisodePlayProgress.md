### 更新单集播放进度

根据 eid 和 pid 更新该单集播放的进度

#### 请求地址

> /episode_play_progress_update

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数          | 必填 | 类型   | 说明           |
| :------------ | :--- | :----- | -------------- |
| data          | true | array  | 请求体         |
| data.eid      | true | string | 单集 id        |
| data.pid      | true | string | 节目 id        |
| data.progress | true | number | 已播放的秒数   |
| data.playedAt | true | string | 上次播放的时间 |

#### 返回字段

无

#### 示例

> 地址：https://www.example.com/episode_play_progress_update

参数

```javascript
{
    "data": [
      // 可传多个
      {
          "eid": "64db2d493fa4090b744c313c",
          "pid": "61791d921989541784257779",
          "progress": 2678,
          "playedAt": "2024-05-16T08:44:04.351Z"
      }
    ]
}
```

响应

```javascript
{
}
```
