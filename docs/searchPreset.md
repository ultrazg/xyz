### 「你可能想搜的内容」

在搜索框里显示你可能想搜的内容，可以通过 `/search` 接口并指定 `type = ALL` 来使用

#### 请求地址

> /search_preset

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

| 返回字段 | 类型   | 说明       |
| :------- | :----- | :--------- |
| text     | string | 显示的文本 |
| ...      | ...    | ...        |


#### 示例

> 地址：https://www.example.com/search_preset

响应

``` javascript
{
  code: 200,
  data: {
    data: [
      {
        link: "cosmos://page.cos/search?tab=all&keyword=%E6%88%9B%E7%BA%B3",
        resident: false,
        text: "戛纳",
        type: "FIXED",
      },
      // ...
    ],
  },
  msg: "OK",
}
```