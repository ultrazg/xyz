### 更新收藏单集

更新收藏单集

#### 请求地址

> /favorite_episode_update

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数      | 必填 | 类型    | 说明         |
| :-------- | :--- | :------ | ------------ |
| eid       | true | string  | 单集 id      |
| favorited | true | boolean | 是否加入收藏 |

#### 返回字段

无


#### 示例

> 地址：https://www.example.com/favorite_episode_update

参数

```javascript
{
  "eid":"64db2d493fa4090b744c313c",
  "favorited":true
}
```

响应

```javascript
{
  code: 200,
  data: {},
  msg: "OK",
}
```