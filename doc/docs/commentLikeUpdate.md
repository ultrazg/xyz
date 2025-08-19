### 点赞/取消点赞评论

点赞或取消点赞评论

#### 请求地址

> /comment_like_update

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数  | 必填 | 类型    | 说明                                                           |
| :---- | :--- | :------ | -------------------------------------------------------------- |
| id    | true | string  | 评论 id                                                        |
| liked | true | boolean | **true** 为点赞，**false** 为取消点赞                          |
| type  | true | string  | 点赞评论的类型。**COMMENT** 为评论，**PICK** 为「TA 们的喜欢」 |

#### 返回字段

无

#### 示例

> 地址：https://www.example.com/comment_like_update

参数

```javascript
{
  "id": "66229567e14d49b402c9e7de",
  "liked": true，
  "type": "COMMENT"
}
```

响应

```javascript
{
    "code": 200,
    "data": {},
    "msg": "OK"
}
```
