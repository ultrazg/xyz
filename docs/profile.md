### 查询我的信息

查询已登录账号的个人信息

#### 请求地址

> /profile

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

| 返回字段       | 类型   | 说明                                               |
| :------------- | :----- | :------------------------------------------------- |
| type           | string | 类别。节目、单集等，详情看 type 对应的类别说明文档 |
| uid            | string | 当前登录用户的 uid                                 |
| avatar         | object | 登录用户的头像信息                                 |
| nickname       | string | 昵称                                               |
| wechatUserInfo | object | 绑定的微信信息                                     |
| phoneNumber    | object | 绑定的手机号                                       |
| jikeUserInfo   | object | 绑定的即刻 APP 信息                                |
| ...            | ...    | ...                                                |


#### 示例

> 地址：https://www.example.com/profile

响应

```javascript
{
  "data": {
    "type": "USER",
    "uid": "YOUR-UID",
    "avatar": {
      "picture": {
        "picUrl": "https://www.example.com",
        "largePicUrl": "https://www.example.com",
        "middlePicUrl": "https://www.example.com",
        "smallPicUrl": "https://www.example.com",
        "thumbnailUrl": "https://www.example.com",
        "format": "jpeg",
        "width": 132,
        "height": 132
      }
    },
    "nickname": "YOUR-NICKNAME",
    "isNicknameSet": true,
    "gender": "MALE",
    "isCancelled": false,
    "readTrackInfo": {},
    "ipLoc": "广东",
    "birthYear": 1998,
    "industry": "互联网/IT",
    "wechatUserInfo": {
      "nickName": "YOUR-WECHAT-NICKNAME"
    },
    "phoneNumber": {
      "mobilePhoneNumber": "131****1111",
      "areaCode": "+86"
    },
    "phoneNumberNeeded": false,
    "jikeUserInfo": {
      "nickname": "YOUR-JIKE-NICKNAME"
    },
    "isBlockedByViewer": false,
    "isInvited": true,
    "debug": false,
    "authorship": [],
    "certifications": []
  }
}
```