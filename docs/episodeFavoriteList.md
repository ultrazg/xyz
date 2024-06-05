### 查询已收藏的单集

查询已收藏的单集

#### 请求地址

> /favorite_episode_list

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

| 返回字段             | 类型   | 说明                                     |
| :------------------- | :----- | :--------------------------------------- |
| type            | string | 类别。节目、单集等，详情看type对应的类别说明文档  |
|pid|string|节目id|
|eid|string|单集id|
|title|string|标题|
|author|string|作者|
|shownotes|string|注释|
|description|string|描述|
|media|object|播客音频信息（大小、链接等）|
|podcasters|array|播客嘉宾的信息（头像、昵称和uid等）|
|...|...|...|


#### 示例

> 地址：https://www.example.com/favorite_episode_list

响应

```javascript
{
  code: 200,
  data: {
    data: [
      {
        clapCount: 305,
        commentCount: 522,
        description: "🎭本期主播：\n* 史炎（@史炎nacl）...",
        duration: 4976,
        eid: "64db2d493fa4090b744c313c",
        enclosure: {
          url: "https://media.xyzcdn.net/ljhlwIAI7JID7aeBiWcdcClP0NhP.m4a",
        },
        favoriteCount: 172,
        image: {
          largePicUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZoUlNJaEp4YW9QNWNXYVFoOEwtTzNuOVFucEYucG5n.png@large",
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
          id: "ljhlwIAI7JID7aeBiWcdcClP0NhP.m4a",
          mimeType: "audio/mp4",
          size: 80553186,
          source: {
            mode: "PUBLIC",
            url: "https://media.xyzcdn.net/ljhlwIAI7JID7aeBiWcdcClP0NhP.m4a",
          },
        },
        mediaKey: "ljhlwIAI7JID7aeBiWcdcClP0NhP.m4a",
        payType: "FREE",
        permissions: [
          {
            name: "SHARE",
            status: "PERMITTED",
          },
          // ...
        ],
        pid: "61791d921989541784257779",
        playCount: 93732,
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
            // ...
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
        pubDate: "2023-08-15T16:00:00.000Z",
        readTrackInfo: {},
        shownotes:
          '<p><strong>🎭本期主播：</strong></p>\n<ul>\n  <li>史炎（<a href="https://weibo.com/u/1452977132">@史炎nacl</a>）...',
        sponsors: [],
        status: "NORMAL",
        title: "95. 你是不是觉得自己挺幽默的：喜剧演员乱谈幽默",
        transcript: {
          mediaId: "ljhlwIAI7JID7aeBiWcdcClP0NhP.m4a",
        },
        type: "EPISODE",
        wechatShare: {
          style: "LINK_CARD",
        },
      },
      // ...
    ],
  },
  msg: "OK",
}
```