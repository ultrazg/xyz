### 编辑精选历史

获取「编辑精选」的往日精选历史内容

#### 请求地址

> /editor_pick_list_history

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数        | 必填  | 类型   | 说明                       |
| :---------- | :---- | :----- | -------------------------- |
| loadMoreKey | false | string | 加载更多的参数。由接口返回 |

#### 返回字段

| 返回字段             | 类型    | 说明       |
| :------------------- | :------ | :--------- |
| data                 | array  | 信息集合 |
| data.date                 | string  | 发布时间 |
| data.picks                 | array  | 信息集合 |
| data.picks.comment                 | object  | 评论相关 |
| data.picks.episode                 | object  | 单集相关 |
| loadMoreKey | string | 加载更多的参数，如出现这个字段则表示有下一页 |


#### 示例

> 地址：https://www.example.com/editor_pick_list_history

参数

```javascript
{
  "loadMoreKey": "2025-08-16T04:00:00.000Z"
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: [
      {
        date: "2025-08-20",
        dateIsoStr: "2025-08-20T04:00:00.000Z",
        picks: [
          {
            comment: {
              author: {
                avatar: {
                  picture: {
                    format: "webp",
                    height: 1000,
                    largePicUrl:
                      "https://image.xyzcdn.net/FlLf4grV5jC3-uFgM7jsSJRTC2oy.webp@large",
                    // ...
                    width: 1000,
                  },
                },
                gender: "FEMALE",
                ipLoc: "上海",
                isBlockedByViewer: false,
                isCancelled: false,
                isNicknameSet: false,
                nickname: "拯救世界的jerry",
                readTrackInfo: {},
                relation: "STRANGE",
                type: "USER",
                uid: "6527ad31edce67104ad353dc",
              },
              authorAssociation: "NONE",
              badges: [],
              collected: false,
              createdAt: "2025-08-19T10:20:35.788Z",
              entities: [],
              id: "68a44ff3ee674c13e403e5cd",
              ipLoc: "上海",
              isAuthorMuted: false,
              level: 1,
              likeCount: 2,
              liked: false,
              owner: {
                id: "689ee07ef9040f9dc34ba2fe",
                type: "EPISODE",
              },
              permissions: [
                {
                  name: "SHARE",
                  status: "PERMITTED",
                },
                // ...
              ],
              pid: "66ed1e69dbf9bb012f843e39",
              pinned: false,
              replyCount: 0,
              status: "NORMAL",
              text: "惊喜，峰哥上节目了！他说生存哲学我是信的，尤其是他讲理财时风险管理，真正洞察人性弱点了。",
              type: "COMMENT",
            },
            episode: {
              clapCount: 34,
              commentCount: 53,
              description: "本期嘉宾\n...",
              duration: 3617,
              eid: "689ee07ef9040f9dc34ba2fe",
              enclosure: {
                url: "https://media.xyzcdn.net/66ed1e69dbf9bb012f843e39/lkDqWyCMbK_SfUliA6U4ntageC6R.m4a",
              },
              favoriteCount: 68,
              ipLoc: "辽宁",
              isCustomized: false,
              isFavorited: false,
              isFinished: false,
              isPicked: false,
              isPlayed: false,
              isPrivateMedia: false,
              labels: [],
              media: {
                id: "66ed1e69dbf9bb012f843e39/lkDqWyCMbK_SfUliA6U4ntageC6R.m4a",
                mimeType: "audio/mp4",
                size: 116432593,
                source: {
                  mode: "PUBLIC",
                  url: "https://media.xyzcdn.net/66ed1e69dbf9bb012f843e39/lkDqWyCMbK_SfUliA6U4ntageC6R.m4a",
                },
              },
              mediaKey:
                "66ed1e69dbf9bb012f843e39/lkDqWyCMbK_SfUliA6U4ntageC6R.m4a",
              payType: "FREE",
              permissions: [
                {
                  name: "SHARE",
                  status: "PERMITTED",
                },
                // ...
              ],
              pid: "66ed1e69dbf9bb012f843e39",
              playCount: 5082,
              podcast: {
                author: "Cindyzhx",
                color: {
                  dark: "#FD8BA3",
                  light: "#FD8BA3",
                  original: "#F3ADB9",
                },
                contacts: [
                  {
                    name: "华夏基金",
                    note: "订阅公众号",
                    type: "wechatOfficialAccounts",
                  },
                ],
                description: "大家好，欢迎来到【大方谈钱】，...",
                episodeCount: 47,
                hasPopularEpisodes: true,
                hasTopic: true,
                image: {
                  format: "jpeg",
                  height: 3000,
                  largePicUrl:
                    "https://image.xyzcdn.net/FpZaxL5QwNrfCV_n0l3hsEBRqLWt.jpg@large",
                  // ...
                  width: 3000,
                },
                isCustomized: false,
                latestEpisodePubDate: "2025-08-17T22:30:00.000Z",
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
                pid: "66ed1e69dbf9bb012f843e39",
                podcasters: [
                  {
                    avatar: {
                      picture: {
                        format: "jpeg",
                        height: 132,
                        largePicUrl:
                          "https://image.xyzcdn.net/FqBm643mWaCmpoHdO7VjlzIHFZHq@large",
                        // ...
                        width: 132,
                      },
                    },
                    bio: "",
                    gender: "FEMALE",
                    ipLoc: "北京",
                    isBlockedByViewer: false,
                    isCancelled: false,
                    isNicknameSet: true,
                    nickname: "章衡HX",
                    readTrackInfo: {},
                    relation: "STRANGE",
                    type: "USER",
                    uid: "63b7f223edce67104a506dc8",
                  },
                  // ...
                ],
                readTrackInfo: {},
                showZhuiguangIcon: false,
                status: "NORMAL",
                subscriptionCount: 96615,
                subscriptionPush: false,
                subscriptionPushPriority: "HIGH",
                subscriptionStar: false,
                subscriptionStatus: "OFF",
                syncMode: "SELF_HOSTING",
                title: "大方谈钱",
                topicLabels: ["ALL"],
                type: "PODCAST",
              },
              pubDate: "2025-08-17T22:30:00.000Z",
              readTrackInfo: {},
              shownotes: "<p><strong>本期嘉宾</strong></p>...",
              sponsors: [],
              status: "NORMAL",
              title: "个人风险管理哲学：峰哥亡命天涯的保命智慧",
              transcript: {
                mediaId:
                  "66ed1e69dbf9bb012f843e39/lkDqWyCMbK_SfUliA6U4ntageC6R.m4a",
              },
              type: "EPISODE",
              wechatShare: {
                style: "LINK_CARD",
              },
            },
            recentAudiences: [
              {
                avatar: {
                  picture: {
                    format: "jpeg",
                    height: 132,
                    largePicUrl:
                      "https://image.xyzcdn.net/Fp-2TuK9IuQme31Tb2zZKs5ERNII@large",
                    // ...
                    width: 132,
                  },
                },
                gender: "MALE",
                ipLoc: "浙江",
                isBlockedByViewer: false,
                isCancelled: false,
                isNicknameSet: true,
                nickname: "狗肉汤浇山药地",
                readTrackInfo: {},
                relation: "STRANGE",
                type: "USER",
                uid: "655f2642edce67104aa29f51",
              },
              // ...
            ],
          },
          // ...
        ],
      },
      // ..
    ],
    loadMoreKey: "2025-08-16T04:00:00.000Z",
  },
  msg: "OK",
}
```