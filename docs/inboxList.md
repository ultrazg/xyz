### 订阅更新列表

查询已订阅的节目是否有更新。每次查询 20 条数据，可分页

#### 请求地址

> /inbox_list

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

> 地址：https://www.example.com/inbox_list

参数

```javascript
{
  // 传入 loadMoreKey 这个参数即可实现分页查询
  "loadMoreKey": {
    "pubDate": "2024-04-03T09:45:00.000Z",
    "id": "660d287a3723a5a6ce7fa929"
  }
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: [
      {
        clapCount: 492,
        commentCount: 736,
        description: "📖本期主播：\n* 谷大白话（@谷大白话）...",
        duration: 7687,
        eid: "665068274efbc0c3dccd06d4",
        enclosure: {
          url: "https://media.xyzcdn.net/ljVbllR5CY8xo3_szqJopo20b9C9.m4a",
        },
        favoriteCount: 517,
        image: {
          largePicUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z2R0tpY1k1MUpybE1CbjJrb05ZMVRmNk14R1cucG5n.png@large",
          middlePicUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z2R0tpY1k1MUpybE1CbjJrb05ZMVRmNk14R1cucG5n.png@middle",
          picUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z2R0tpY1k1MUpybE1CbjJrb05ZMVRmNk14R1cucG5n.png",
          smallPicUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z2R0tpY1k1MUpybE1CbjJrb05ZMVRmNk14R1cucG5n.png@small",
          thumbnailUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z2R0tpY1k1MUpybE1CbjJrb05ZMVRmNk14R1cucG5n.png@thumbnail",
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
          id: "ljVbllR5CY8xo3_szqJopo20b9C9.m4a",
          mimeType: "audio/mp4",
          size: 124441059,
          source: {
            mode: "PUBLIC",
            url: "https://media.xyzcdn.net/ljVbllR5CY8xo3_szqJopo20b9C9.m4a",
          },
        },
        mediaKey: "ljVbllR5CY8xo3_szqJopo20b9C9.m4a",
        payType: "FREE",
        permissions: [
          {
            name: "SHARE",
            status: "PERMITTED",
          },
          // ...
        ],
        pid: "61791d921989541784257779",
        playCount: 92458,
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
          episodeCount: 139,
          hasPopularEpisodes: true,
          image: {
            largePicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png@large",
            middlePicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png@middle",
            picUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png",
            smallPicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png@small",
            thumbnailUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png@thumbnail",
          },
          isCustomized: false,
          latestEpisodePubDate: "2024-05-24T16:00:00.000Z",
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
          subscriptionCount: 155877,
          subscriptionPush: false,
          subscriptionPushPriority: "HIGH",
          subscriptionStar: false,
          subscriptionStatus: "ON",
          syncMode: "SELF_HOSTING",
          title: "不开玩笑 Jokes Aside",
          topicLabels: [],
          type: "PODCAST",
        },
        pubDate: "2024-05-24T16:00:00.000Z",
        readTrackInfo: {},
        shownotes: "<p><strong>📖本期主播：...",
        sponsors: [],
        status: "NORMAL",
        title: "137. 我太想进步了！聊聊我们学过的有用无用的东西",
        transcript: {
          mediaId: "ljVbllR5CY8xo3_szqJopo20b9C9.m4a",
        },
        type: "EPISODE",
        wechatShare: {
          style: "LINK_CARD",
        },
      },
      // ...
    ],
    loadMoreKey: {
      id: "66172f8d340b6a1f4fd41348",
      pubDate: "2024-04-11T00:11:07.836Z",
    },
    userStats: {
      subscriptionCount: 5,
    },
  },
  msg: "OK",
}
```