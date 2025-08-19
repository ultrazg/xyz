### 新节目广场

查询「新节目广场」的数据

#### 请求地址

> /pilot_discovery_list

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

data 里是单集的集合，参考单集的字段

#### 示例

> 地址：https://www.example.com/pilot_discovery_list

参数

无

响应

```javascript
{
  code: 200,
  data: {
    data: [
      {
        episode: {
          clapCount: 1,
          commentCount: 4,
          description: "我是𝟗𝟖年已婚已育，...",
          duration: 1333,
          eid: "6888eb47edf3fa32d5205517",
          enclosure: {
            url: "https://media.xyzcdn.net/686329381f93173c1b46c01d/lrtp6ERV23qa_7BRBLGlWwyQ7AfB.m4a",
          },
          favoriteCount: 2,
          ipLoc: "北京",
          isCustomized: false,
          isFavorited: false,
          isFinished: false,
          isPicked: false,
          isPlayed: false,
          isPrivateMedia: false,
          labels: [],
          media: {
            id: "686329381f93173c1b46c01d/lrtp6ERV23qa_7BRBLGlWwyQ7AfB.m4a",
            mimeType: "audio/mp4",
            size: 21575636,
            source: {
              mode: "PUBLIC",
              url: "https://media.xyzcdn.net/686329381f93173c1b46c01d/lrtp6ERV23qa_7BRBLGlWwyQ7AfB.m4a",
            },
          },
          mediaKey: "686329381f93173c1b46c01d/lrtp6ERV23qa_7BRBLGlWwyQ7AfB.m4a",
          payType: "FREE",
          permissions: [
            {
              name: "SHARE",
              status: "PERMITTED",
            },
            // ...
          ],
          pid: "686329381f93173c1b46c01d",
          playCount: 345,
          podcast: {
            author: "橙子的心想事橙",
            brief: "深度访谈𝟏𝟎𝟎位优秀女性的搞钱路径与能量",
            color: {
              dark: "#DB7C34",
              light: "#DB7C34",
              original: "#DB7C34",
            },
            contacts: [
              {
                name: "橙子下班搞钱记",
                note: "订阅公众号",
                type: "wechatOfficialAccounts",
              },
              // ...
            ],
            description: "一档既拆解搞钱路径，...",
            episodeCount: 6,
            hasPopularEpisodes: true,
            image: {
              format: "jpeg",
              height: 2460,
              largePicUrl:
                "https://image.xyzcdn.net/Fob8eLfOF9d39EndY2E2RS6SQhml.jpg@large",
              // ...
              width: 2460,
            },
            isCustomized: false,
            latestEpisodePubDate: "2025-08-18T23:00:00.000Z",
            payEpisodeCount: 0,
            payType: "FREE",
            permissions: [
              {
                name: "SHARE",
                status: "PERMITTED",
              },
              {
                name: "AI_SUMMARIZE_EPISODE",
                status: "PERMITTED",
              },
            ],
            pid: "686329381f93173c1b46c01d",
            podcasters: [
              {
                avatar: {
                  picture: {
                    format: "jpeg",
                    height: 3001,
                    largePicUrl:
                      "https://image.xyzcdn.net/Fv3-ec8Z8CJFdVg4PBodRgkDEgfj.jpg@large",
                    // ...
                    width: 3001,
                  },
                },
                bio: "去做，要么成功，要么成长！",
                gender: "FEMALE",
                ipLoc: "北京",
                isBlockedByViewer: false,
                isCancelled: false,
                isNicknameSet: true,
                nickname: "橙子的心想事橙",
                readTrackInfo: {},
                relation: "STRANGE",
                type: "USER",
                uid: "61150a0ee0f5e723bb244fc2",
              },
            ],
            readTrackInfo: {},
            showZhuiguangIcon: false,
            status: "NORMAL",
            subscriptionCount: 542,
            subscriptionPush: false,
            subscriptionPushPriority: "HIGH",
            subscriptionStar: false,
            subscriptionStatus: "OFF",
            syncMode: "SELF_HOSTING",
            title: "谈钱说爱deep talk",
            topicLabels: [],
            type: "PODCAST",
          },
          pubDate: "2025-07-29T23:00:00.000Z",
          readTrackInfo: {
            backendRecDescType: "duration",
            recModuleType: "NEW_PODCAST_EPISODE_SQUARE",
            recallPolicy: "NPSER_ACTIVE_USER_RANDOM",
            score: "0.0439",
          },
          shownotes: "<h2>我是𝟗𝟖年已婚已育，...",
          sponsors: [],
          status: "NORMAL",
          title: "01 认真做一档女性搞钱访谈播客！聊聊搞钱路径与爱的能量",
          transcript: {
            mediaId:
              "686329381f93173c1b46c01d/lrtp6ERV23qa_7BRBLGlWwyQ7AfB.m4a",
          },
          type: "EPISODE",
          wechatShare: {
            style: "LINK_CARD",
          },
        },
        hasNegativeFeedback: false,
        recommendation: "22 分钟",
        relatedUsers: [],
      },
      // ...
    ],
  },
  msg: "OK",
}
```