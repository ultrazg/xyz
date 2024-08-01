### 分类下的标签

根据分类 id 查询该分类下的标签内容

#### 请求地址

> /category_list_tab

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求参数

| 参数 | 必填 | 类型 |说明|
| ---- | ---- | ---- |----|
| categoryId | true | string |分类的 id。由 `/category_list` 接口获取|

#### 返回字段

| 返回字段 | 类型   | 说明       |
| -------- | ------ | ---------- |
| id       | string | 标签的 id  |
| name     | string | 标签的名称 |




#### 示例

> 地址：https://www.example.com/category_list_tab

参数

```javascript
{
  "categoryId": "63c76a8924b1622727bd3219"
}
```

响应

``` javascript
{
  code: 200,
  data: {
    data: [
      {
        id: "ALL",
        name: "全部",
      },
      {
        id: "HOT",
        name: "近期热门",
      },
      {
        id: "RECOMMEND",
        name: "为你推荐",
      },
    ],
  },
  msg: "OK",
}
```