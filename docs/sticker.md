### 查询已获得的贴纸

根据 uid 查询已获得的贴纸

#### 请求地址

> /sticker

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数 | 必填 | 类型   | 说明          |
| :--- | :--- | :----- | ------------- |
| uid  | true | string | 用户的 uid |

#### 返回字段

| 返回字段    | 类型   | 说明               |
| :---------- | :----- | :----------------- |
| total       | number | 已获得贴纸的数量   |
| description | string | 贴纸描述           |
| image       | object | 贴纸链接、尺寸信息 |
| issuer      | string | 贴纸发行人         |
| name        | string | 贴纸名称           |
| ownedAt     | time   | 获取时间           |
| ...         | ...    | ...                |


#### 示例

> 地址：https://www.example.com/sticker

参数

```javascript
{
  "uid": "UID"
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: [
      {
        description: "希望播客能让你快乐",
        id: "5ad20e98-34c4-49aa-8edb-e568c231b27c",
        image: {
          format: "png",
          height: 570,
          largePicUrl:
            "https://image.xyzcdn.net/stickers/statics/2024.png@large",
          middlePicUrl:
            "https://image.xyzcdn.net/stickers/statics/2024.png@middle",
          picUrl: "https://image.xyzcdn.net/stickers/statics/2024.png",
          smallPicUrl:
            "https://image.xyzcdn.net/stickers/statics/2024.png@small",
          thumbnailUrl:
            "https://image.xyzcdn.net/stickers/statics/2024.png@thumbnail",
          width: 570,
        },
        issuer: "小宇宙",
        name: "2024年快乐",
        number: "AN-246353965",
        ownedAt: "2024-03-13T05:52:22.868Z",
        package: {
          image: {
            format: "png",
            height: 195,
            largePicUrl:
              "https://image.xyzcdn.net/stickers/packages/common.png@large",
            middlePicUrl:
              "https://image.xyzcdn.net/stickers/packages/common.png@middle",
            picUrl: "https://image.xyzcdn.net/stickers/packages/common.png",
            smallPicUrl:
              "https://image.xyzcdn.net/stickers/packages/common.png@small",
            thumbnailUrl:
              "https://image.xyzcdn.net/stickers/packages/common.png@thumbnail",
            width: 573,
          },
        },
        scale: 3,
      },
      // ...
    ],
    isPrivate: false,
    total: 4,
  },
  msg: "OK",
}
```