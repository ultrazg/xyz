### 查询黑名单列表

获取已拉黑的用户列表

#### 请求地址

> /blocked_user_lists

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

| 返回字段          | 类型    | 说明         |
| ----------------- | ------- | ------------ |
| avatar            | object  | 用户头像信息 |
| nickname          | string  | 用户昵称     |
| isBlockedByViewer | boolean | 是否拉黑     |
| uid               | string  | 用户 uid     |




#### 示例

> 地址：https://www.example.com/blocked_user_lists

参数

无

响应

``` javascript
{
  code: 200,
  data: {
    data: [
      {
        avatar: {
          picture: {
            format: "jpeg",
            height: 1500,
            largePicUrl:
              "https://image.xyzcdn.net/FvMdrce...",
            middlePicUrl:
              "https://image.xyzcdn.net/FvMdrce...",
            picUrl: "https://image.xyzcdn.net/FvMdrce...",
            smallPicUrl:
              "https://image.xyzcdn.net/FvMdrce...",
            thumbnailUrl:
              "https://image.xyzcdn.net/FvMdrce...",
            width: 1500,
          },
        },
        bio: "",
        ipLoc: "北京",
        isBlockedByViewer: true,
        isCancelled: false,
        isNicknameSet: true,
        nickname: NICKNAME,
        readTrackInfo: {},
        relation: "STRANGE",
        type: "USER",
        uid: UID,
      },
    ],
  },
  msg: "OK",
}
```