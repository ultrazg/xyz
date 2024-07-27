### 精彩时间点

查询单集播放时间线中的精彩时间点，每一段代表当前播放时间段的精彩标记数

#### 图示

<img src="episode_clap.jpeg" alt="episode_clap">

#### 请求地址

> /episode_clap

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数     | 必填 | 类型   | 说明     |
| :------- | :--- | :----- | -------- |
| eid      | true | string | 单集 id  |
| duration | true | number | 单集时长 |

#### 返回字段

| 返回字段     | 类型  | 说明                                                         |
| :----------- | :---- | :----------------------------------------------------------- |
| episodeClaps | array | 精彩时间点的段数，数组里有多少个对象则对应时间线中有多少个时间点（如图示中的圆角方块）。**count** 字段为标记数量 |
| myClaps      | array | 「我」标记的时间点，数组里有多少个对象则表示「我」标记的数量。**index** 字段为时间点（如图示中的圆角方块）的**下标**。以下面的示例为例：「我」在下标为 **10** 的时间点（如图示中的圆角方块）中标记了 **2** 次 |


#### 示例

> 地址：https://www.example.com/episode_clap

参数

```javascript
{
  "duration": 7284,
  "eid": "6634b5c603bcdd73a9480a2f",
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: {
      episodeClaps: [
        {
          count: 40,
        },
        {
          count: 15,
        },
        // ...
      ],
      myClaps: [
        {
          index: 10,
        },
        {
          index: 10,
        },
      ],
    },
  },
  msg: "OK",
}
```