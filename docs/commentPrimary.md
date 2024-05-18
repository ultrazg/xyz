### 查询单集的评论

根据 eid 查询单集的评论

#### 请求地址

> /comment_primary

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数  | 必填 | 类型   | 说明                                                         |
| :---- | :--- | :----- | ------------------------------------------------------------ |
| id    | true | string | 单集 eid                                                     |
| order | true | string | 排序条件。**全部评论（HOT）**、**最新评论（TIME）**、**时点评论（TIMESTAMP）** |

#### 返回字段

| 返回字段         | 类型    | 说明                                                         |
| :--------------- | :------ | :----------------------------------------------------------- |
| type             | string  | 类别                                                         |
| id               | string  | 该评论的 id                                                  |
| owner            | object  | 该评论所属的单集的信息                                       |
| author           | object  | 该评论发布者的信息                                           |
| text             | string  | 评论正文                                                     |
| likeCount        | number  | 评论点赞数                                                   |
| liked            | boolean | 是否已点赞该评论                                             |
| collected        | boolean | 是否已收藏该评论                                             |
| createdAt        | string  | 评论发布时间                                                 |
| pinned           | boolean | 是否置顶评论                                                 |
| threadReplyCount | number  | 评论回复数                                                   |
| replies          | array   | 展示该评论的部分回复内容                                     |
| totalCount       | number  | 总评论数                                                     |
| loadMoreKey      | object  | 如果存在下一页，则会返回该对象。将这个对象传入请求参数，便可实现分页查询 |
| ...              | ...     | ...                                                          |


#### 示例

> 地址：https://www.example.com/comment_primary

参数

```javascript
{
  "eid": "66467d2c251bd96e6cdcddde",
  "order": "HOT",
  // 传入 loadMoreKey 这个参数即可实现分页查询
  loadMoreKey: {
    direction: "NEXT",
    hotSortScore: 0.9648972534042627,
    id: "66484d8d10cd1345ef983a90",
  },
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: [
      {
        id: "66468a29aa440ecaaa8db490",
        type: "COMMENT",
        owner: {
          id: "66467d2c251bd96e6cdcddde",
          type: "EPISODE",
        },
        author: {
          type: "USER",
          uid: "65f8d456edce67104a48c738",
          avatar: {
            picture: {
              picUrl:
                "https://image.xyzcdn.net/Fmv_obCVr2PFKj-ncjYb463yDDUh.jpg",
              largePicUrl:
                "https://image.xyzcdn.net/Fmv_obCVr2PFKj-ncjYb463yDDUh.jpg@large",
              middlePicUrl:
                "https://image.xyzcdn.net/Fmv_obCVr2PFKj-ncjYb463yDDUh.jpg@middle",
              smallPicUrl:
                "https://image.xyzcdn.net/Fmv_obCVr2PFKj-ncjYb463yDDUh.jpg@small",
              thumbnailUrl:
                "https://image.xyzcdn.net/Fmv_obCVr2PFKj-ncjYb463yDDUh.jpg@thumbnail",
              format: "jpeg",
              width: 1500,
              height: 1500,
            },
          },
          nickname: "PPxiaSKy",
          isNicknameSet: true,
          gender: "MALE",
          isCancelled: false,
          readTrackInfo: {},
          ipLoc: "天津",
          relation: "STRANGE",
          isBlockedByViewer: false,
        },
        authorAssociation: "NONE",
        text: "生产队的驴都不敢这么歇。",
        level: 1,
        likeCount: 29,
        liked: false,
        collected: false,
        createdAt: "2024-05-16T22:35:21.138Z",
        status: "NORMAL",
        permissions: [
          {
            name: "SHARE",
            status: "PERMITTED",
          },
          {
            name: "COMMENT",
            status: "PERMITTED",
          },
          {
            name: "VOICE_COMMENT",
            status: "DENIED",
          },
          {
            name: "COMMENT_PIN_OPERATION",
            status: "DENIED",
          },
          {
            name: "MUTE_COMMENT_AUTHOR",
            status: "DENIED",
          },
          {
            name: "DELETE",
            status: "DENIED",
          },
        ],
        pid: "61dab969ab2c6dcf05d5ffdf",
        pinned: false,
        isAuthorMuted: false,
        entities: [],
        badges: [],
        ipLoc: "天津",
        threadReplyCount: 2,
        replies: [
          {
            id: "66469a168029f4442e95a226",
            type: "COMMENT",
            owner: {
              id: "66467d2c251bd96e6cdcddde",
              type: "EPISODE",
            },
            thread: "66468a29aa440ecaaa8db490",
            author: {
              type: "USER",
              uid: "60a28d1ee0f5e723bba9e254",
              avatar: {
                picture: {
                  picUrl:
                    "https://image.xyzcdn.net/Fo4xvk1XtpoktwZbWRpEZb_gzDUO",
                  largePicUrl:
                    "https://image.xyzcdn.net/Fo4xvk1XtpoktwZbWRpEZb_gzDUO@large",
                  middlePicUrl:
                    "https://image.xyzcdn.net/Fo4xvk1XtpoktwZbWRpEZb_gzDUO@middle",
                  smallPicUrl:
                    "https://image.xyzcdn.net/Fo4xvk1XtpoktwZbWRpEZb_gzDUO@small",
                  thumbnailUrl:
                    "https://image.xyzcdn.net/Fo4xvk1XtpoktwZbWRpEZb_gzDUO@thumbnail",
                  width: 663,
                  height: 663,
                  format: "png",
                },
              },
              nickname: "HD647764v",
              isNicknameSet: false,
              gender: "MALE",
              isCancelled: false,
              readTrackInfo: {},
              ipLoc: "山东",
              relation: "STRANGE",
              isBlockedByViewer: false,
            },
            authorAssociation: "NONE",
            text: "生产队的队都不敢这么歇。",
            level: 2,
            likeCount: 2,
            liked: false,
            collected: false,
            createdAt: "2024-05-16T23:43:18.184Z",
            status: "NORMAL",
            permissions: [
              {
                name: "SHARE",
                status: "PERMITTED",
              },
              {
                name: "COMMENT",
                status: "PERMITTED",
              },
              {
                name: "VOICE_COMMENT",
                status: "DENIED",
              },
              {
                name: "COMMENT_PIN_OPERATION",
                status: "DENIED",
              },
              {
                name: "MUTE_COMMENT_AUTHOR",
                status: "DENIED",
              },
              {
                name: "DELETE",
                status: "DENIED",
              },
            ],
            pid: "61dab969ab2c6dcf05d5ffdf",
            pinned: false,
            isAuthorMuted: false,
            entities: [],
            badges: [],
            ipLoc: "山东",
            replyToComment: {
              id: "66468a29aa440ecaaa8db490",
              type: "COMMENT",
              owner: {
                id: "66467d2c251bd96e6cdcddde",
                type: "EPISODE",
              },
              author: {
                type: "USER",
                uid: "65f8d456edce67104a48c738",
                avatar: {
                  picture: {
                    picUrl:
                      "https://image.xyzcdn.net/Fmv_obCVr2PFKj-ncjYb463yDDUh.jpg",
                    largePicUrl:
                      "https://image.xyzcdn.net/Fmv_obCVr2PFKj-ncjYb463yDDUh.jpg@large",
                    middlePicUrl:
                      "https://image.xyzcdn.net/Fmv_obCVr2PFKj-ncjYb463yDDUh.jpg@middle",
                    smallPicUrl:
                      "https://image.xyzcdn.net/Fmv_obCVr2PFKj-ncjYb463yDDUh.jpg@small",
                    thumbnailUrl:
                      "https://image.xyzcdn.net/Fmv_obCVr2PFKj-ncjYb463yDDUh.jpg@thumbnail",
                    format: "jpeg",
                    width: 1500,
                    height: 1500,
                  },
                },
                nickname: "PPxiaSKy",
                isNicknameSet: true,
                gender: "MALE",
                isCancelled: false,
                readTrackInfo: {},
                ipLoc: "天津",
                relation: "STRANGE",
                isBlockedByViewer: false,
              },
              authorAssociation: "NONE",
              text: "生产队的驴都不敢这么歇。",
              level: 1,
              likeCount: 29,
              liked: false,
              collected: false,
              createdAt: "2024-05-16T22:35:21.138Z",
              status: "NORMAL",
              permissions: [
                {
                  name: "SHARE",
                  status: "PERMITTED",
                },
                {
                  name: "COMMENT",
                  status: "PERMITTED",
                },
                {
                  name: "VOICE_COMMENT",
                  status: "DENIED",
                },
                {
                  name: "COMMENT_PIN_OPERATION",
                  status: "DENIED",
                },
                {
                  name: "MUTE_COMMENT_AUTHOR",
                  status: "DENIED",
                },
                {
                  name: "DELETE",
                  status: "DENIED",
                },
              ],
              pid: "61dab969ab2c6dcf05d5ffdf",
              pinned: false,
              isAuthorMuted: false,
              entities: [],
              badges: [],
              ipLoc: "天津",
            },
          },
          // ...
        ],
      },
      // ...
    ],
    loadNextKey: {
      direction: "NEXT",
      hotSortScore: 0.9648972534042627,
      id: "66484d8d10cd1345ef983a90",
    },
    loadMoreKey: {
      direction: "NEXT",
      hotSortScore: 0.9648972534042627,
      id: "66484d8d10cd1345ef983a90",
    },
    totalCount: 116,
  },
  msg: "OK",
}
```