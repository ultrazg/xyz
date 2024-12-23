### 获取节目主体信息

获取节目的主体等信息

#### 请求地址

> /podcast_get_info

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

| 返回字段 | 类型   | 说明     |
| :------- | :----- | :------- |
| ipLoc    | string | IP属地   |
| subject  | string | 账号主体 |
| ...      | ...    | ...      |


#### 示例

> 地址：https://www.example.com/podcast_get_info

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
    data: {
      ipLoc: "上海",
      subject: "个人",
    },
  },
  msg: "OK",
}
```