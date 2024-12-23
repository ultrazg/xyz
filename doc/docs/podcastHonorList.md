### 获取节目荣誉墙

获取节目所获得的荣誉

#### 请求地址

> /podcast_honor_list

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
| pid  | true | string | 节目 id |

#### 返回字段

| 返回字段      | 类型   | 说明     |
| :------------ | :----- | :------- |
| campaignTitle | string | 活动标题 |
| id            | string | id       |
| title         | string | 标题     |


#### 示例

> 地址：https://www.example.com/podcast_honor_list

参数

```javascript
{
  "pid": "61791d921989541784257779"
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: [
      {
        campaignTitle: "小宇宙播客大赏",
        id: "65b082520937a0a8385ffa9b",
        title: "2023年度热门播客",
        url: "cosmos://page.cos/...",
      },
      // ...
    ],
  },
  msg: "OK",
}
```