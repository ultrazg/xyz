### 用户统计数据

根据 uid 查询用户的个人信息

#### 请求地址

> /get_profile

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数 | 必填 | 类型   | 说明     |
| :--- | :--- | :----- | -------- |
| uid  | true | string | 用户 uid |

#### 返回字段

| 返回字段       | 类型   | 说明                                               |
| :------------- | :----- | :------------------------------------------------- |
| type           | string | 类别。节目、单集等，详情看 type 对应的类别说明文档 |
| uid            | string | 用户的 uid                                         |
| avatar         | object | 用户的头像信息                                     |
| nickname       | string | 昵称                                               |
| wechatUserInfo | object | 绑定的微信信息                                     |
| phoneNumber    | object | 绑定的手机号                                       |
| jikeUserInfo   | object | 绑定的即刻 APP 信息                                |
| relation       | string | 是否关注此用户。值为 **FOLLOWING** 即为 **已关注** |
| ...            | ...    | ...                                                |


#### 示例

> 地址：https://www.example.com/user_stats

参数

```javascript
{
  "uid":"UID"
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: {
      authorship: [],
      avatar: {
        picture: {
          format: "jpeg",
          height: 132,
          largePicUrl:
            "https://www.example.com",
          middlePicUrl:
            "https://www.example.com",
          picUrl: "https://www.example.com",
          smallPicUrl:
            "https://www.example.com",
          thumbnailUrl:
            "https://www.example.com",
          width: 132,
        },
      },
      birthYear: 1998,
      certifications: [],
      debug: false,
      gender: "MALE",
      industry: "互联网/IT",
      ipLoc: "广东",
      isBlockedByViewer: false,
      isCancelled: false,
      isInvited: true,
      isNicknameSet: true,
      jikeUserInfo: {
        nickname: "JIKE_NICKNAME",
      },
      nickname: "NICKNAME",
      phoneNumber: {
        areaCode: "+86",
        mobilePhoneNumber: "131****1111",
      },
      phoneNumberNeeded: false,
      readTrackInfo: {},
      type: "USER",
      uid: "UID",
      wechatUserInfo: {
        nickName: "WECHAT_NICKNAME",
      },
    },
  },
  msg: "OK",
}
```