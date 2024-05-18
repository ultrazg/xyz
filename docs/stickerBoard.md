### 查询我的贴纸墙

查询已在贴纸墙中已张贴的贴纸

#### 请求地址

> /sticker_board

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数 | 必填 | 类型   | 说明             |
| :--- | :--- | :----- | ---------------- |
| uid  | true | string | 已登录用户的 uid |

#### 返回字段

| 返回字段 | 类型   | 说明                       |
| :------- | :----- | :------------------------- |
| stickers | array  | 展示的贴纸                 |
| sticker  | object | 贴纸描述、链接、尺寸信息等 |
| rotation | number | 贴纸旋转角度               |
| issuer   | string | 贴纸发行人                 |
| name     | string | 贴纸名称                   |
| ownedAt  | time   | 获取时间                   |
| scale    | number | 贴纸缩放等级               |
| x、y     | number | 贴纸坐标（相较于屏幕）     |
| ...      | ...    | ...                        |


#### 示例

> 地址：https://www.example.com/sticker_board

参数

```javascript
{
  "uid": "YOUR-UID"
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: {
      stickers: [
        {
          rotation: 0.4586824968005772,
          sticker: {
            description: "在国际空间站绕地球飞行66圈，别晕船了",
            id: "abe0a5d5-4adb-48d9-9ebc-9eaa8019fcba",
            image: {
              format: "png",
              height: 570,
              largePicUrl:
                "https://image.xyzcdn.net/stickers/statics/100_hours.png@large",
              middlePicUrl:
                "https://image.xyzcdn.net/stickers/statics/100_hours.png@middle",
              picUrl: "https://image.xyzcdn.net/stickers/statics/100_hours.png",
              smallPicUrl:
                "https://image.xyzcdn.net/stickers/statics/100_hours.png@small",
              thumbnailUrl:
                "https://image.xyzcdn.net/stickers/statics/100_hours.png@thumbnail",
              width: 570,
            },
            issuer: "小宇宙",
            name: "总收听100小时",
            number: "LH-1963425",
            ownedAt: "2024-02-02T16:47:38.537Z",
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
          x: -77,
          y: 144,
        },
        // ...
      ],
    },
    isPrivate: false,
  },
  msg: "OK",
}
```