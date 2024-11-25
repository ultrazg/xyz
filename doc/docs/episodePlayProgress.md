### 查询单集播放进度

根据 eid 查询该单集播放的进度

#### 请求地址

> /episode_play_progress

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数 | 必填 | 类型  | 说明                                |
| :--- | :--- | :---- | ----------------------------------- |
| eids | true | array | 单集 id。是一个数组，可传入多个 eid |

#### 返回字段

| 返回字段 | 类型   | 说明           |
| :------- | :----- | :------------- |
| progress | number | 已播放的秒数   |
| pid      | string | 节目 id        |
| eid      | string | 单集 id        |
| playedAt | string | 上次播放的时间 |

#### 示例

> 地址：https://www.example.com/episode_play_progress

参数

```javascript
{
    "eids": [
      	// 可传入多个
        "64db2d493fa4090b744c313c"
    ]
}
```

响应

```javascript
{
  data: [
    {
      eid: "64db2d493fa4090b744c313c",
      pid: "61791d921989541784257779",
      playedAt: "2024-05-16T08:44:04.351Z",
      progress: 2678,
    },
  ],
}
```
