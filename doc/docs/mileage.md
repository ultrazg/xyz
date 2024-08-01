### 查询收听数据概览

查询「我的个人」收听数据

#### 请求地址

> /mileage_get

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

| 返回字段                   | 类型   | 说明               |
| :------------------------- | :----- | :----------------- |
| totalPlayedSeconds         | number | 总收听时长         |
| lastSevenDayPlayedSeconds  | number | 最近三十天收听记录 |
| lastThirtyDayPlayedSeconds | number | 总收听记录         |
| tagline                    | string | 一行标语           |


#### 示例

> 地址：https://www.example.com/mileage_get

响应

```javascript
{
    "code": 200,
    "data": {
        "data": {
            "lastSevenDayPlayedSeconds": 0,
            "lastThirtyDayPlayedSeconds": 0,
            "tagline": "谷神星自转了15圈",
            "totalPlayedSeconds": 505379
        }
    },
    "msg": "OK"
}
```