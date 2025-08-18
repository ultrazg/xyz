### 创建评论

回复他人评论或者给单集评论

#### 请求地址

> /comment_create

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数             | 必填  | 类型   | 说明                                                         |
| :--------------- | :---- | :----- | ------------------------------------------------------------ |
| text             | true  | string | 评论内容                                                     |
| id               | true  | string | 单集的 id                                                    |
| type             | true  | string | 类型，填 **EPISODE** 即可                                    |
| replyToCommentId | false | string | 要回复的评论的 id。如果传入此参数，则是回复他人的评论，为空则是为单集添加评论 |

#### 返回字段

| 返回字段 | 类型   | 说明       |
| :------- | :----- | :--------- |
| data.author | object | 当前账号的信息 |
| data.badges | object | 当前账号取得的徽章 |
| data.collected | boolean | 是否收藏该评论 |
| data.createdAt | string | 发表时间 |
| data.id | string | 评论的 id |
| data.likeCount | number | 评论的点赞数 |
| data.liked | boolean | 是否点赞该评论 |
| data.owner | object | 评论所属的单集信息 |
| data.pid | string | 评论所属的节目 |
| data.pinned | boolean | 是否被置顶 |
| data.replyCount | number | 评论的回复数 |
| data.text | string | 评论的内容 |
| toast | string | 提示文案 |


#### 示例

> 地址：https://www.example.com/comment_create

参数

```javascript
{
  "type": "EPISODE",
  "text": "Hello World!",
  "id": "681dc93cb7c8a9962c3fb0e6",
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: {
      author: {
        avatar: {
          picture: {
            format: "jpeg",
            height: 1003,
            largePicUrl: "https://www.example.com",
            // ...
            width: 1003,
          },
        },
        bio: "example",
        birthYear: 1998,
        gender: "MALE",
        industry: "互联网/IT",
        ipLoc: "广东",
        isBlockedByViewer: false,
        isCancelled: false,
        isNicknameSet: true,
        jikeUserInfo: {
          nickname: "YOUR-NICKNAME",
        },
        nickname: "YOUR-NICKNAME",
        phoneNumber: {
          areaCode: "+86",
          mobilePhoneNumber: "131****1111",
        },
        phoneNumberNeeded: false,
        readTrackInfo: {},
        type: "USER",
        uid: "YOUR-UID",
        wechatUserInfo: {
          nickName: "YOUR-NICKNAME",
        },
      },
      authorAssociation: "NONE",
      badges: [
        {
          icon: {
            format: "png",
            height: 42,
            picUrl:
              "https://image.xyzcdn.net/xyz/badge-icon/100h_plus_0930.png",
            width: 90,
          },
          tip: "你收听该节目达100小时",
        },
      ],
      collected: false,
      createdAt: "2025-08-18T16:10:57.114Z",
      entities: [],
      id: "68a35091d38bacaa7ed7ca8e",
      isAuthorMuted: false,
      level: 1,
      likeCount: 0,
      liked: false,
      owner: {
        id: "681dc93cb7c8a9962c3fb0e6",
        type: "EPISODE",
      },
      permissions: [
        {
          name: "SHARE",
          status: "PERMITTED",
        },
        // ...
      ],
      pid: "61791d921989541784257779",
      pinned: false,
      replyCount: 0,
      status: "NORMAL",
      text: "😄",
      type: "COMMENT",
    },
    toast: "已评论",
  },
  msg: "OK",
}
```