### 查询「我」关注的人

根据 uid 查询「我」关注的人

#### 请求地址

> /following_list

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数 | 必填 | 类型   | 说明       |
| :--- | :--- | :----- | ---------- |
| uid  | true | string | 用户的 uid |

#### 返回字段

| 返回字段 | 类型   | 说明                                         |
| :------- | :----- | :------------------------------------------- |
| avatar   | object | 用户头像相关信息                             |
| gender   | string | 性别                                         |
| ipLoc    | string | IP 属地                                      |
| nickname | string | 用户昵称                                     |
| relation | string | 「我」是否已关注，**FOLLOWING** 为**已关注** |
| uid      | string | 用户 uid                                     |


#### 示例

> 地址：https://www.example.com/following_list

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
        avatar: {
          picture: {
            format: "jpeg",
            height: 1200,
            largePicUrl:
              "https://image.xyzcdn.net/Fv-cQDvJstRNycAWJewSN-_D8U9w.jpg@large",
            // ...
            width: 1200,
          },
        },
        gender: "FEMALE",
        ipLoc: "河北",
        isBlockedByViewer: false,
        isCancelled: false,
        isNicknameSet: true,
        nickname: "晚上吃桃子",
        readTrackInfo: {},
        relation: "FOLLOWING",
        type: "USER",
        uid: "UID",
      },
    ],
  },
  msg: "OK",
}
```

