### 全部分类

获取全部分类

#### 请求地址

> /category_list

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求参数

无

#### 返回字段

| 返回字段 | 类型   | 说明             |
| -------- | ------ | ---------------- |
| emoji    | emoji  | 分类的 emoji     |
| icon     | object | 分类的 icon 相关 |
| id       | string | 分类的 id        |
| name     | string | 分类的名称       |




#### 示例

> 地址：https://www.example.com/category_list

参数

无

响应

``` javascript
{
  code: 200,
  data: {
    data: [
      {
        emoji: "🏷️",
        icon: {
          format: "png",
          height: 171,
          largePicUrl:
            "https://image.xyzcdn.net/Ftb67SRne_SSLVRuAzHHzILdadha.png@large",
          middlePicUrl:
            "https://image.xyzcdn.net/Ftb67SRne_SSLVRuAzHHzILdadha.png@middle",
          picUrl: "https://image.xyzcdn.net/Ftb67SRne_SSLVRuAzHHzILdadha.png",
          smallPicUrl:
            "https://image.xyzcdn.net/Ftb67SRne_SSLVRuAzHHzILdadha.png@small",
          thumbnailUrl:
            "https://image.xyzcdn.net/Ftb67SRne_SSLVRuAzHHzILdadha.png@thumbnail",
          width: 171,
        },
        id: "63c76a8924b1622727bd321b",
        name: "商业",
      },
      // ...
    ],
  },
  msg: "OK",
}
```