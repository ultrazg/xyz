### 查询用户信息

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

| 返回字段       | 类型   | 说明                                                   |
| :------------- | :----- | :----------------------------------------------------- |
| type           | string | 类别。节目、单集等，详情看 type 对应的类别说明文档     |
| uid            | string | 用户的 uid                                             |
| avatar         | object | 用户的头像信息                                         |
| nickname       | string | 昵称                                                   |
| relation       | string | 是否关注此用户。值为 **FOLLOWING** 即为 **已关注**     |
| bio            | string | 个签                                                   |
| certifications | array  | 认证信息。**kind** 表示认证类别                        |
| gender         | string | 性别                                                   |
| ipLoc          | string | IP 属地                                                |
| ...            | ...    | ...                                                    |


#### 示例

> 地址：https://www.example.com/get_profile

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
      avatar: {
        picture: {
          format: "jpeg",
          height: 960,
          largePicUrl:
            "https://image.xyzcdn.net/FtmJ4F4CkGUez2kbrvpst_CnemoN.jpg@large",
          // ...
          width: 960,
        },
      },
      bio: "浪漫至死不渝   （wb同名）",
      certifications: [
        {
          kind: "PODCASTER",
          shows: [
            {
              title: "嗯呢嗯呢呗",
            },
          ],
        },
      ],
      gender: "FEMALE",
      ipLoc: "广东",
      isBlockedByViewer: false,
      isCancelled: false,
      isInvited: true,
      isNicknameSet: true,
      nickname: "-鳄鱼听了都要做噩梦_",
      readTrackInfo: {},
      relation: "STRANGE",
      type: "USER",
      uid: "61d2702c2cbd7c01765ca5d9",
    },
  },
  msg: "OK",
}
```