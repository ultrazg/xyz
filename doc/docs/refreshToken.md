### 刷新token

刷新 token。当接口返回 `401` 时调用此接口以获取有效的 token 信息
#### 请求地址

> /refresh_token

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求参数

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |
| x-jike-refresh-token | true | string | x-jike-refresh-token |

#### 返回字段

| 返回字段             | 类型   | 说明                                     |
| :------------------- | :----- | :--------------------------------------- |
| x-jike-access-token  | string | token信息。后续请求都需要，注意保存      |
| x-jike-refresh-token | string | refresh-token。刷新token时需要，注意保存 |


#### 示例

> 地址：https://www.example.com/refresh_token

参数

```javascript
{
  "x-jike-access-token": "YOUR-ACCESS-TOKEN",
  "x-jike-refresh-token": "YOUR-REFRESH-TOKEN"
}
```

响应

``` javascript
{
  code: 200,
  msg: "OK",
  data: {
    "x-jike-access-token": "YOUR-NEW-ACCESS-TOKEN",
    success: true,
    "x-jike-refresh-token": "YOUR-NEW-REFRESH-TOKEN",
  },
}
```