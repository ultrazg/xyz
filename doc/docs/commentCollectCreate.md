### 收藏评论

收藏单集内某个评论

#### 请求地址

> /comment_collect_create

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数      | 必填 | 类型   | 说明    |
| :-------- | :--- | :----- | ------- |
| commentId | true | string | 评论 id |

#### 返回字段

| 返回字段 | 类型   | 说明       |
| :------- | :----- | :--------- |
| toast    | string | Toast 内容 |


#### 示例

> 地址：https://www.example.com/comment_collect_create

参数

```javascript
{
  "commentId": "66229567e14d49b402c9e7de"
}
```

响应

```javascript
{
    "code": 200,
    "data": {
        "toast": "已收藏"
    },
    "msg": "OK"
}
```