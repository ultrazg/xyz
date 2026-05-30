### 查询付费单集的音频链接

根据付费单集的 eid 查询该单集的音频链接，需要在**小宇宙 APP**购买该单集

#### 请求地址

> /private_media_get

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数        | 必填  | 类型   | 说明                                                  |
| :---------- | :---- | :----- | ----------------------------------------------------- |
| eid         | true  | string | 单集 id

#### 返回字段

| 返回字段             | 类型   | 说明                                     |
| :------------------- | :----- | :--------------------------------------- |
|url|string|该付费单集的音频链接|


#### 示例

> 地址：https://www.example.com/private_media_get

参数

```javascript
{
  "eid": "6a1513fb8b5c94538eb70c98",
}
```

响应
```javascript
{
    "code": 200,
    "data": {
        "data": {
            "key": "63...a2/li...3q.m4a",
            "ttl": 14400,
            "url": "https://pmedia.xyzcdn.net/63...a2/li...3q.m4a?auth_key=17...1a"
        }
    },
    "msg": "OK"
}
```