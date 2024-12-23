### 查询节目公告

根据 pid 查询节目的公告

#### 请求地址

> /podcast_bulletin

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数 | 必填 | 类型   | 说明    |
| :--- | :--- | :----- | ------- |
| pid  | true | string | 节目 id |

#### 返回字段

| 返回字段           | 类型    | 说明                                                      |
| :----------------- | :------ | :-------------------------------------------------------- |
| content            | string  | 公告内容                                                  |
| createdAt          | string  | 公告发布时间                                              |
| offline            | boolean | 是否下线                                                  |
| podcast            | object  | 公告所属的节目                                            |
| reactions          | array   | 公告内点赞详情（点赞数、点赞 emoji 图标的 base64 格式等） |
| uniqueVisitorCount | number  | 公告被查看数                                              |
| ...                | ...     | ...                                                       |


#### 示例

> 地址：https://www.example.com/podcast_bulletin

参数

```javascript
{
  "pid": "61791d921989541784257779"
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: {
      content:
        "各位听友，我们将于10月23日0点进行一期特别加更，欢迎大家届时踊跃收听哦❤️",
      createdAt: "2024-10-21T14:11:31.875Z",
      id: "67165f530d2f24f289c4ad43",
      offline: false,
      pictures: [],
      podcast: {
        author: "不开玩笑小助手",
        brief: "由猫头鹰喜剧出品",
        color: {
          dark: "#0B0B0B",
          light: "#ACACAC",
          original: "#1A1A1A",
        },
        contacts: [
          {
            name: "猫头鹰喜剧",
            note: "订阅公众号",
            type: "wechatOfficialAccounts",
          },
          // ...
        ],
        description: "一档时而开玩笑，时而不开玩笑的播客节目，...",
        episodeCount: 170,
        hasPopularEpisodes: true,
        image: {
          format: "png",
          height: 4600,
          largePicUrl: "https://image.xyzcdn.net/...",
          middlePicUrl: "https://image.xyzcdn.net/...",
          picUrl: "https://image.xyzcdn.net/...",
          smallPicUrl: "https://image.xyzcdn.net/...",
          thumbnailUrl: "https://image.xyzcdn.net/...",
          width: 4600,
        },
        isCustomized: false,
        latestEpisodePubDate: "2024-12-20T16:00:00.000Z",
        payEpisodeCount: 0,
        payType: "FREE",
        permissions: [
          {
            name: "SHARE",
            status: "PERMITTED",
          },
          // ...
        ],
        pid: "61791d921989541784257779",
        podcasters: [
          {
            avatar: {
              picture: {
                format: "jpeg",
                height: 3000,
                largePicUrl:
                  "https://image.xyzcdn.net/FnQ-E7VcqLbzqplvdVPGrQRGHmxC.jpg@large",
                middlePicUrl:
                  "https://image.xyzcdn.net/FnQ-E7VcqLbzqplvdVPGrQRGHmxC.jpg@middle",
                picUrl:
                  "https://image.xyzcdn.net/FnQ-E7VcqLbzqplvdVPGrQRGHmxC.jpg",
                smallPicUrl:
                  "https://image.xyzcdn.net/FnQ-E7VcqLbzqplvdVPGrQRGHmxC.jpg@small",
                thumbnailUrl:
                  "https://image.xyzcdn.net/FnQ-E7VcqLbzqplvdVPGrQRGHmxC.jpg@thumbnail",
                width: 3000,
              },
            },
            bio: "不开玩笑 Jokes Aside 由猫头鹰喜剧出品。",
            ipLoc: "上海",
            isBlockedByViewer: false,
            isCancelled: false,
            isNicknameSet: false,
            nickname: "不开玩笑小助手",
            readTrackInfo: {},
            relation: "STRANGE",
            type: "USER",
            uid: "61518a53e0f5e723bb520729",
          },
          // ...
        ],
        readTrackInfo: {},
        showZhuiguangIcon: false,
        status: "NORMAL",
        subscriptionCount: 263197,
        subscriptionPush: true,
        subscriptionPushPriority: "HIGH",
        subscriptionStar: false,
        subscriptionStatus: "ON",
        syncMode: "SELF_HOSTING",
        title: "不开玩笑 Jokes Aside",
        topicLabels: [],
        type: "PODCAST",
      },
      reactions: [
        {
          count: 695,
          name: "thumbsup",
          reacted: false,
          url: "https://media.xyzcdn.net/FsYQT0zX4gBI_E9l60mQGkEjSum0.json",
        },
        // ...
      ],
      uniqueVisitorCount: 169173,
    },
  },
  msg: "OK",
}
```