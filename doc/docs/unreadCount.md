### 未读消息

查询未读消息的数量，具体内容请到 APP 内查看

#### 请求地址

> /unread_count

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

无

#### 返回字段

| 返回字段    | 类型   | 说明           |
| :---------- | :----- | :------------- |
| unreadCount | number | 未读消息的数量 |


#### 示例

> 地址：https://www.example.com/unread_count

响应

```javascript
{
  code: 200,
  data: {
    data: {
      unreadCount: 0,
    },
  },
  msg: "OK",
}
```