### 关注/取关用户

关注或取关用户

#### 请求地址

> /relation_update

#### 请求方式

> POST

#### 支持格式

> JSON


#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数     | 必填 | 类型   | 说明                                     |
| -------- | ---- | ------ | ---------------------------------------- |
| uid      | true | string | 要关注/取关的用户 uid                    |
| relation | true | string | 关注传 **FOLLOWING**，取关传 **STRANGE** |



#### 返回字段

无




#### 示例

> 地址：https://www.example.com/relation_update

参数

```javascript
{
  "uid": UID
  "relation": "STRANGE"
}
```

响应

``` javascript
{
  code: 200,
  data: {},
  msg: "OK",
}
```