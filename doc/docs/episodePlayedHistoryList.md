### 收听历史

查询收听历史

#### 请求地址

> /episode_played_history_list

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
| loadMoreKey | false | object | 分页查询的条件，由接口返回 |

#### 返回字段

| 返回字段    | 类型   | 说明                                                         |
| :---------- | :----- | :----------------------------------------------------------- |
| type        | string | 类别。节目、单集等，详情看type对应的类别说明文档             |
| pid         | string | 节目id                                                       |
| eid         | string | 单集id                                                       |
| title       | string | 标题                                                         |
| author      | string | 作者                                                         |
| shownotes   | string | 注释                                                         |
| description | string | 描述                                                         |
| media       | object | 播客音频信息（大小、链接等）                                 |
| podcasters  | array  | 播客嘉宾的信息（头像、昵称和uid等）                          |
| loadMoreKey | object | 如果存在下一页，则会返回该对象。将这个对象传入请求参数，便可实现分页查询 |
| ...         | ...    | ...                                                          |


#### 示例

> 地址：https://www.example.com/episode_played_history_list

参数

```javascript
{
   // 传入 loadMoreKey 这个参数即可实现分页查询
  "loadMoreKey": "2023-10-27T16:00:00.000Z"
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: [
      {
        episode: {
          clapCount: 924,
          commentCount: 596,
          description: "🤬本期主播：\n* 史炎（@史炎nacl）...",
          duration: 5144,
          eid: "651f519619dde7bf6a08a3b4",
          enclosure: {
            url: "https://media.xyzcdn.net/lh34krqj3Mn8pfDCkxfvOsYSxNjF.m4a",
          },
          favoriteCount: 287,
          image: {
            largePicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZyRHJSeWFsUU8wclhBeGN6RmdXcXQ5TlNLcy0ucG5n.png@large",
            // ...
          },
          ipLoc: "上海",
          isCustomized: false,
          isFavorited: false,
          isFinished: true,
          isPicked: false,
          isPlayed: true,
          isPrivateMedia: false,
          labels: [],
          media: {
            id: "lh34krqj3Mn8pfDCkxfvOsYSxNjF.m4a",
            mimeType: "audio/mp4",
            size: 83279902,
            source: {
              mode: "PUBLIC",
              url: "https://media.xyzcdn.net/lh34krqj3Mn8pfDCkxfvOsYSxNjF.m4a",
            },
          },
          mediaKey: "lh34krqj3Mn8pfDCkxfvOsYSxNjF.m4a",
          payType: "FREE",
          permissions: [
            {
              name: "SHARE",
              status: "PERMITTED",
            },
            // ...
          ],
          pid: "61791d921989541784257779",
          playCount: 119738,
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
              //  ...
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
            status: "NORMAL",
            subscriptionCount: 158526,
            subscriptionPush: false,
            subscriptionPushPriority: "HIGH",
            subscriptionStar: true,
            subscriptionStatus: "ON",
            syncMode: "SELF_HOSTING",
            title: "不开玩笑 Jokes Aside",
            topicLabels: [],
            type: "PODCAST",
          },
          pubDate: "2023-10-06T16:00:00.000Z",
          readTrackInfo: {},
          shownotes:
            "<p><strong>🤬本期主播：</strong></p>\n<ul>\n  <li>史炎...",
          sponsors: [],
          status: "NORMAL",
          title: "103. 老板你够了！奇葩领导品鉴会！",
          transcript: {
            mediaId: "lh34krqj3Mn8pfDCkxfvOsYSxNjF.m4a",
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
                height: 991,
                largePicUrl:
                  "https://image.xyzcdn.net/FmxkwnGdmovKHerCnGUNOygh4ZuR.jpg@large",
                // ...
                width: 991,
              },
            },
            bio: "要银河，不要迎合！",
            gender: "FEMALE",
            ipLoc: "山东",
            isBlockedByViewer: false,
            isCancelled: false,
            isNicknameSet: true,
            nickname: "淡季动物园",
            readTrackInfo: {},
            relation: "STRANGE",
            type: "USER",
            uid: "63ef0efaedce67104aa1883b",
          },
          {
            avatar: {
              picture: {
                format: "jpeg",
                height: 132,
                largePicUrl:
                  "https://image.xyzcdn.net/FtRNGMFZoCn7WRtbuy5ye1yuDmn2@large",
                // ...
                width: 132,
              },
            },
            ipLoc: "北京",
            isBlockedByViewer: false,
            isCancelled: false,
            isNicknameSet: false,
            nickname: "白云去来",
            readTrackInfo: {},
            relation: "STRANGE",
            type: "USER",
            uid: "6022028fe0f5e723bb024972",
          },
          {
            avatar: {
              picture: {
                format: "jpeg",
                height: 132,
                largePicUrl:
                  "https://image.xyzcdn.net/FgSmnc0_PXG0cENx8k2yT3P5Oubg@large",
                //  ...
                width: 132,
              },
            },
            gender: "FEMALE",
            ipLoc: "北京",
            isBlockedByViewer: false,
            isCancelled: false,
            isNicknameSet: true,
            nickname: "沈天天_Y9MW",
            readTrackInfo: {},
            relation: "STRANGE",
            type: "USER",
            uid: "648091aeedce67104ab7fd0d",
          },
        ],
      },
      // ...
    ],
    loadMoreKey: "2023-12-07T23:53:47.084Z",
  },
  msg: "OK",
}
```