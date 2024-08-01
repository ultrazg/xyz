### 查询已收藏的评论

查询已收藏的评论

#### 请求地址

> /comment_collect_list

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

| 返回字段    | 类型    | 说明                 |
| :---------- | :------ | :------------------- |
| author      | object  | 发布该评论的用户信息 |
| collected   | boolean | 是否收藏             |
| collectedAt | string  | 收藏时间             |
| createdAt   | string  | 该评论发布时间       |
| episode     | object  | 评论所在的单集信息   |
| likeCount   | number  | 评论点赞数           |
| liked       | boolean | 是否已点赞该评论     |
| pinned      | boolean | 是否置顶评论         |
| ...         | ...     | ...                  |


#### 示例

> 地址：https://www.example.com/comment_collect_list

响应

```javascript
{
  code: 200,
  data: {
    data: [
      {
        author: {
          avatar: {
            picture: {
              format: "jpeg",
              height: 132,
              largePicUrl:
                "https://image.xyzcdn.net/FsE468B1uAENVGRQseVerfezpafM@large",
              // ...
              width: 132,
            },
          },
          bio: "",
          gender: "FEMALE",
          ipLoc: "上海",
          isBlockedByViewer: false,
          isCancelled: false,
          isNicknameSet: true,
          nickname: "谷谷鸟鸟",
          readTrackInfo: {},
          relation: "STRANGE",
          type: "USER",
          uid: "658d919dedce67104abda382",
        },
        authorAssociation: "NONE",
        badges: [],
        collected: true,
        collectedAt: "2024-06-04T16:40:34.433Z",
        createdAt: "2024-04-19T16:01:43.671Z",
        entities: [],
        episode: {
          clapCount: 1305,
          commentCount: 1199,
          description: "☺️本期主播：\n* 史炎（@史炎nacl）...",
          duration: 6413,
          eid: "662251f28b343a8e9aa4ffc7",
          enclosure: {
            url: "https://media.xyzcdn.net/lix_S74v-PGpZlACpprN11KlRr3u.m4a",
          },
          favoriteCount: 790,
          image: {
            largePicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z1NDlCdzVvVzJJb3VtVHdaZC1aVktrSXZZdFoucG5n.png@large",
            // ...
          },
          ipLoc: "上海",
          isCustomized: false,
          isFavorited: true,
          isFinished: true,
          isPicked: false,
          isPlayed: true,
          isPrivateMedia: false,
          labels: [],
          media: {
            id: "lix_S74v-PGpZlACpprN11KlRr3u.m4a",
            mimeType: "audio/mp4",
            size: 103821609,
            source: {
              mode: "PUBLIC",
              url: "https://media.xyzcdn.net/lix_S74v-PGpZlACpprN11KlRr3u.m4a",
            },
          },
          mediaKey: "lix_S74v-PGpZlACpprN11KlRr3u.m4a",
          payType: "FREE",
          permissions: [
            {
              name: "SHARE",
              status: "PERMITTED",
            },
            // ...
          ],
          pid: "61791d921989541784257779",
          playCount: 154165,
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
            description:
              "一个时而开玩笑，时而不开玩笑的播客节目，由猫头鹰喜剧出品。\n\n公众号“猫头鹰喜剧”回复“听友群”，小助手会把你拉进群聊~",
            episodeCount: 140,
            hasPopularEpisodes: true,
            image: {
              largePicUrl:
                "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png@large",
              //  ...
            },
            isCustomized: false,
            latestEpisodePubDate: "2024-05-31T16:00:00.000Z",
            payEpisodeCount: 0,
            payType: "FREE",
            permissions: [
              {
                name: "SHARE",
                status: "PERMITTED",
              },
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
                    // ...
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
            status: "NORMAL",
            subscriptionCount: 158495,
            subscriptionPush: false,
            subscriptionPushPriority: "HIGH",
            subscriptionStar: true,
            subscriptionStatus: "ON",
            syncMode: "SELF_HOSTING",
            title: "不开玩笑 Jokes Aside",
            topicLabels: [],
            type: "PODCAST",
          },
          pubDate: "2024-04-19T16:00:00.000Z",
          readTrackInfo: {},
          shownotes:
            "<p><strong>☺️本期主播：</strong></p>\n<ul>\n  <li>史炎...",
          sponsors: [],
          status: "NORMAL",
          title: "132. GGGG卑卑卑卑：聊聊自信和自卑",
          transcript: {
            mediaId: "lix_S74v-PGpZlACpprN11KlRr3u.m4a",
          },
          type: "EPISODE",
          wechatShare: {
            style: "LINK_CARD",
          },
        },
        id: "66229567e14d49b402c9e7de",
        ipLoc: "上海",
        isAuthorMuted: false,
        level: 1,
        likeCount: 501,
        liked: false,
        owner: {
          id: "662251f28b343a8e9aa4ffc7",
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
        status: "NORMAL",
        text: "豆豆 你就是干这个的！",
        type: "COMMENT",
      },
    ],
    loadMoreKey: "2024-06-04T16:40:34.433Z",
  },
  msg: "OK",
}
```