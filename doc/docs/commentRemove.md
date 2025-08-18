### 删除评论

删除已发表的评论

#### 请求地址

> /comment_remove

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数             | 必填  | 类型   | 说明                                                         |
| :--------------- | :---- | :----- | ------------------------------------------------------------ |
| commentId             | true  | string | 评论的 id                                               |

#### 返回字段

| 返回字段   | 类型    | 说明               |
| :--------- | :------ | :----------------- |
| toast     | string  | 提示文案     |


#### 示例

> 地址：https://www.example.com/comment_remove

参数

```javascript
{
  "commentId": "68a35091d38bacaa7ed7ca8e"
}
```

响应

```javascript
{
  code: 200,
  data: {
    toast: "已删除",
  },
  msg: "OK",
}
```