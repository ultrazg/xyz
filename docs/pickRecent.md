### 查询用户的喜欢（片段）

查询个人主页「用户的喜欢」片段内容

#### 请求地址

> /pick_list_recent

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

| 返回字段 | 类型   | 说明                 |
| :------- | :----- | :------------------- |
| story    | object | 用户标记喜欢所留下的内容 |
| pickedAt | string | 用户标记的时间       |
| episode  | object | 用户标记喜欢所在的的单集  |
| total    | number | 用户标记的总数       |
| likeCount | number  | 用户标记的点赞数     |
| isLiked   | boolean | 是否点赞             |

#### 示例

> 地址：https://www.example.com/pick_list_recent

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
        id: "66409fe64e87c41035ce0ab3",
        type: "PICK",
        story: {
          text: "喜欢听他俩聊天，总有一天要把小跑和Will老师拐过来录节目😈",
          emotion: "LIKE",
          iconUrl: "https://image.xyzcdn.net/xyz/pick-icon/like.png",
        },
        pickedAt: "2024-05-12T10:54:30.660Z",
        episode: {
          type: "EPISODE",
          eid: "66265840c3e09d8f37b879f2",
          pid: "61c3fafb6b3149acbfac1e29",
          title: "再谈货币逻辑：阿根廷今天“全面美元化”了吗？",
          shownotes: "<p>货币灵魂三问又回来了！</p>...",
          description: "货币灵魂三问又回来了！...",
          image: {
            picUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZnbU5VaW1jU2RVZ2RGLURpZlA2d01tbGw2WXEucG5n.png",
            // ...
          },
          enclosure: {
            url: "https://media.xyzcdn.net/lixwyg7vph43e-8KRx0VHtgVztG9.m4a",
          },
          isPrivateMedia: false,
          mediaKey: "lixwyg7vph43e-8KRx0VHtgVztG9.m4a",
          media: {
            id: "lixwyg7vph43e-8KRx0VHtgVztG9.m4a",
            size: 117442637,
            mimeType: "audio/mp4",
            source: {
              mode: "PUBLIC",
              url: "https://media.xyzcdn.net/lixwyg7vph43e-8KRx0VHtgVztG9.m4a",
            },
          },
          clapCount: 128,
          commentCount: 41,
          playCount: 7744,
          favoriteCount: 95,
          pubDate: "2024-04-22T12:29:50.410Z",
          status: "NORMAL",
          duration: 4855,
          podcast: {
            type: "PODCAST",
            pid: "61c3fafb6b3149acbfac1e29",
            title: "文理两开花",
            author: "肖小跑",
            brief: "试图在混沌中寻找秩序",
            description: "《文理两开花》：...",
            subscriptionCount: 12803,
            image: {
              picUrl:
                "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZsOWlsejJvNUdzTmN2Tlp0d2NOTU5nUnMxMGUucG5n.png",
              // ...
            },
            color: {
              original: "#2E91B8",
              light: "#0AB3E9",
              dark: "#0AB3E9",
            },
            topicLabels: [],
            syncMode: "SELF_HOSTING",
            episodeCount: 62,
            latestEpisodePubDate: "2024-07-14T09:57:00.570Z",
            subscriptionStatus: "OFF",
            subscriptionPush: false,
            subscriptionPushPriority: "HIGH",
            subscriptionStar: false,
            status: "NORMAL",
            permissions: [
              {
                name: "SHARE",
                status: "PERMITTED",
              },
            ],
            payType: "FREE",
            payEpisodeCount: 0,
            podcasters: [
              {
                type: "USER",
                uid: "5ef59773e93418180f1905c0",
                avatar: {
                  picture: {
                    picUrl:
                      "https://image.xyzcdn.net/FrTwOHwAufWpKKdN-625ASRE8Bze",
                    // ...
                    format: "",
                    width: 0,
                    height: 0,
                  },
                },
                nickname: "肖小跑",
                isNicknameSet: false,
                bio: "一路小跑儿",
                isCancelled: false,
                readTrackInfo: {},
                ipLoc: "中国香港",
                relation: "STRANGE",
                isBlockedByViewer: false,
              },
              //  ...
            ],
            readTrackInfo: {},
            hasPopularEpisodes: true,
            contacts: [
              {
                type: "wechatOfficialAccounts",
                name: "文理两开花",
                note: "订阅公众号",
              },
              {
                type: "jike",
                name: "肖小跑",
                url: "https://m.okjike.com/users/d03b9adf-3acb-46c0-8308-85195f22188f",
              },
              {
                type: "wechat",
                name: "Lei_Salin",
                note: "添加微信",
              },
              {
                type: "custom",
                name: "Newsletter",
                url: "https://wenli.substack.com/ ",
              },
            ],
            isCustomized: false,
          },
          isPlayed: false,
          isFinished: false,
          isPicked: false,
          isFavorited: false,
          permissions: [
            {
              name: "SHARE",
              status: "PERMITTED",
            },
            // ...
          ],
          payType: "FREE",
          wechatShare: {
            style: "LINK_CARD",
          },
          readTrackInfo: {},
          labels: [
            {
              name: "最受欢迎",
              code: "POPULAR",
            },
          ],
          sponsors: [],
          isCustomized: false,
          ipLoc: "四川",
          transcript: {
            mediaId: "lixwyg7vph43e-8KRx0VHtgVztG9.m4a",
          },
        },
        user: {
          type: "USER",
          uid: "5ea69adc68a1682f67d0a71f",
          avatar: {
            picture: {
              picUrl:
                "https://image.xyzcdn.net/FsIRR9SioGRQmqGbjKzqVZ_AevU1.jpg",
              // ...
              format: "jpeg",
              width: 500,
              height: 500,
            },
          },
          nickname: "雨白",
          isNicknameSet: false,
          bio: "不知有汉，无论魏晋",
          gender: "FEMALE",
          isCancelled: false,
          readTrackInfo: {},
          ipLoc: "四川",
          relation: "STRANGE",
          isBlockedByViewer: false,
        },
        readTrackInfo: {},
        likeCount: 37,
        commentCount: 0,
        isLiked: false,
      },
      // ...
    ],
    total: 8,
  },
  msg: "OK",
}
```