### 将用户移出黑名单

取消拉黑用户

#### 请求地址

> /blocked_user_remove

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
| ---- | ---- | ------ | -------- |
| uid  | true | string | 用户 uid |



#### 返回字段

| 返回字段 | 类型   | 说明     |
| -------- | ------ | -------- |
| toast    | string | 说明文案 |




#### 示例

> 地址：https://www.example.com/blocked_user_remove

参数

```javascript
{
  "uid": UID
}
```

响应

``` javascript
{
  code: 200,
  data: {
    toast: "已移出黑名单",
  },
  msg: "OK",
}
```