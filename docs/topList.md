### 获取榜单

获取首页的最热榜、锋芒榜和新星榜数据

#### 请求地址

> /top_list

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数     | 必填 | 类型   | 说明                                                         |
| :------- | :--- | :----- | ------------------------------------------------------------ |
| category | true | string | **HOT** 为查询**最热榜**，**ROCK** 为查询锋芒榜，**NEW** 为查询**新星榜** |

#### 返回字段

| 返回字段    | 类型   | 说明                                             |
| :---------- | :----- | :----------------------------------------------- |
| background  | string | 背景图片                                         |
| information | string | 榜单说明                                         |
| items       | array  | 榜单内容                                         |
| publishDate | string | 更新时间                                         |
| rulesUrl    | string | 榜单规则链接                                     |
| title       | string | 榜单标题                                         |
| type        | string | 类别。节目、单集等，详情看type对应的类别说明文档 |


#### 示例

> 地址：https://www.example.com/top_list

参数

```javascript
{
  "cateory": "HOT",
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: {
      background: "https://image.xyzcdn.net/FgDukmIERjFaXd8fS4omYj2QqPRN.png",
      category: "HOT_EPISODES_IN_24_HOURS",
      id: "66a62a41b5f8c30c4c372605",
      information:
        "这里展示的是最受欢迎的内容，每日8点根据节目单集的数据计算得出。",
      items: [
        {
          item: {
            clapCount: 211,
            commentCount: 353,
            description: "聊聊抬杠\n【主播】\n...",
            duration: 4090,
            eid: "66a517f0c8295bd902e14885",
            enclosure: {
              url: "https://jt.ximalaya.com//GKwRIMAKeY3PAXxxFQL237HY.m4a?channel=rss&album_id=43044571&track_id=744097846&uid=262232071&jt=https://audio.xmcdn.com/storages/8edd-audiofreehighqps/80/F0/GKwRIMAKeY3PAXxxFQL237HY.m4a",
            },
            favoriteCount: 136,
            image: {
              largePicUrl:
                "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy9kMjA4LWF1ZGlvZnJlZWhpZ2hxcHMvMjYvREQvQ01Db09TUUVEOTBDQUFOSm9RQ01kUWFSLmpwZWc=.jpeg@large",
              // ...
            },
            ipLoc: "上海",
            isCustomized: false,
            isFavorited: false,
            isFinished: false,
            isPicked: false,
            isPlayed: false,
            isPrivateMedia: false,
            labels: [],
            media: {
              backupSource: {
                mode: "PUBLIC",
                url: "https://media.xyzcdn.net/5f912d9783c34e85ddd7e23d/luXdaeqc-NLDezGe0sZjzSElC49j.mp4a",
              },
              id: "//GKwRIMAKeY3PAXxxFQL237HY.m4a?channel=rss&album_id=43044571&track_id=744097846&uid=262232071&jt=https%3A%2F%2Faudio.xmcdn.com%2Fstorages%2F8edd-audiofreehighqps%2F80%2FF0%2FGKwRIMAKeY3PAXxxFQL237HY.m4a",
              mimeType: "audio/x-m4a",
              size: 163956120,
              source: {
                mode: "PUBLIC",
                url: "https://jt.ximalaya.com//GKwRIMAKeY3PAXxxFQL237HY.m4a?channel=rss&album_id=43044571&track_id=744097846&uid=262232071&jt=https://audio.xmcdn.com/storages/8edd-audiofreehighqps/80/F0/GKwRIMAKeY3PAXxxFQL237HY.m4a",
              },
            },
            mediaKey:
              "//GKwRIMAKeY3PAXxxFQL237HY.m4a?channel=rss&album_id=43044571&track_id=744097846&uid=262232071&jt=https%3A%2F%2Faudio.xmcdn.com%2Fstorages%2F8edd-audiofreehighqps%2F80%2FF0%2FGKwRIMAKeY3PAXxxFQL237HY.m4a",
            payType: "FREE",
            permissions: [
              {
                name: "SHARE",
                status: "PERMITTED",
              },
              // ...
            ],
            pid: "5f912d9783c34e85ddd7e23d",
            playCount: 27187,
            podcast: {
              author: "嘻谈录",
              color: {
                dark: "#3150CF",
                light: "#3150CF",
                original: "#1A29B4",
              },
              contacts: [
                {
                  name: "嘻谈录",
                  note: "订阅公众号",
                  type: "wechatOfficialAccounts",
                },
                // ...
              ],
              description: "嘻谈录是最正能量的喜剧播客之一，...",
              episodeCount: 162,
              hasPopularEpisodes: true,
              hasTopic: true,
              image: {
                largePicUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy8yZDRlLWF1ZGlvZnJlZWhpZ2hxcHMvQTMvNkYvQ01Db09TSUVGVWl1QUFoRmVRQ09Jc2lELmpwZWc=.jpeg@large",
                // ...
              },
              isCustomized: false,
              latestEpisodePubDate: "2024-07-27T15:50:00.000Z",
              payEpisodeCount: 0,
              payType: "FREE",
              permissions: [
                {
                  name: "SHARE",
                  status: "PERMITTED",
                },
              ],
              pid: "5f912d9783c34e85ddd7e23d",
              podcasters: [
                {
                  avatar: {
                    picture: {
                      format: "jpeg",
                      height: 3000,
                      largePicUrl:
                        "https://image.xyzcdn.net/Fo-3h-lqiR2K-bLnhRJ9lh0EeFP8.jpg@large",
                      //  ...
                      width: 3000,
                    },
                  },
                  bio: "播客制作人/微博@枣仁安神zzz /即刻@嘻谈录小张/还在上学",
                  gender: "MALE",
                  ipLoc: "澳大利亚",
                  isBlockedByViewer: false,
                  isCancelled: false,
                  isNicknameSet: true,
                  nickname: "嘻谈录小张",
                  readTrackInfo: {},
                  relation: "STRANGE",
                  type: "USER",
                  uid: "60093bd0e0f5e723bb44ef52",
                },
                // ...
              ],
              readTrackInfo: {},
              status: "NORMAL",
              subscriptionCount: 132118,
              subscriptionPush: false,
              subscriptionPushPriority: "HIGH",
              subscriptionStar: false,
              subscriptionStatus: "OFF",
              syncMode: "RSS",
              title: "嘻谈录",
              topicLabels: ["ALL"],
              type: "PODCAST",
            },
            pubDate: "2024-07-27T15:50:00.000Z",
            readTrackInfo: {},
            shownotes: '<p style="color:#333333;...',
            sponsors: [],
            status: "NORMAL",
            title: "133.我的家庭没了，我的抬杠变强了",
            transcript: {
              mediaId:
                "5f912d9783c34e85ddd7e23d/luXdaeqc-NLDezGe0sZjzSElC49j.mp4a",
            },
            type: "EPISODE",
            wechatShare: {
              style: "LINK_CARD",
            },
          },
        },
        // ...
      ],
      publishDate: "2024-07-28T16:00:00.000Z",
      rulesUrl:
        "https://feedback.xiaoyuzhoufm.com/v6worU4NnWyL/qa/6315bddf116442001162b2f8",
      targetType: "EPISODE",
      title: "最热榜",
      type: "TOP_LIST",
    },
  },
  msg: "OK",
}
```