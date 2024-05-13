### 查询单集列表

根据节目 pid 查询该节目下的单集列表，单次返回 20 条数据

#### 请求地址

> /episode_list

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
| pid         | true  | string | 节目id                     |
| loadMoreKey | false | object | 分页查询的条件，由接口返回 |

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
|loadMoreKey|object|如果存在下一页，则会返回该对象。将这个对象传入请求参数，便可实现分页查询|
|...|...|...|


#### 示例

> 地址：https://www.example.com/episode_list

参数

```javascript
{
  "pid": "61791d921989541784257779",
   // 传入 loadMoreKey 这个参数即可实现分页查询
  "loadMoreKey": {
      "pubDate": "2023-10-27T16:00:00.000Z",
      "id": "653b307c53c1d329114906bc",
      "direction": "NEXT"
  }
}
```

响应
```javascript
{
  code: 200,
  msg: "OK",
  data: {
    data: [
      {
        type: "EPISODE",
        eid: "660fd1834f66d1c1da2a7e89",
        pid: "61791d921989541784257779",
        title: "130. 世界是个草台班子，你我都是戏中人",
        shownotes: "<p><strong>🌿本期主播：</strong></p>...",
        description: "🌿本期主播：...",
        image: {
          picUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z1blpSYll4eUtEeWlDdEhqaXhaR01pX2JOTHQucG5n.png",
          largePicUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z1blpSYll4eUtEeWlDdEhqaXhaR01pX2JOTHQucG5n.png@large",
          middlePicUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z1blpSYll4eUtEeWlDdEhqaXhaR01pX2JOTHQucG5n.png@middle",
          smallPicUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z1blpSYll4eUtEeWlDdEhqaXhaR01pX2JOTHQucG5n.png@small",
          thumbnailUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z1blpSYll4eUtEeWlDdEhqaXhaR01pX2JOTHQucG5n.png@thumbnail",
        },
        enclosure: {
          url: "https://media.xyzcdn.net/lg0nOGrd9aYDRyrb6o7FMcwzw62K.m4a",
        },
        isPrivateMedia: false,
        mediaKey: "/lg0nOGrd9aYDRyrb6o7FMcwzw62K.m4a",
        media: {
          id: "/lg0nOGrd9aYDRyrb6o7FMcwzw62K.m4a",
          size: 82128299,
          mimeType: "audio/mp4",
          source: {
            mode: "PUBLIC",
            url: "https://media.xyzcdn.net/lg0nOGrd9aYDRyrb6o7FMcwzw62K.m4a",
          },
        },
        clapCount: 655,
        commentCount: 888,
        playCount: 113230,
        favoriteCount: 645,
        pubDate: "2024-04-05T16:00:00.000Z",
        status: "NORMAL",
        duration: 5073,
        podcast: {
          type: "PODCAST",
          pid: "61791d921989541784257779",
          title: "不开玩笑 Jokes Aside",
          author: "不开玩笑小助手",
          brief: "由猫头鹰喜剧出品",
          description:
            "一个时而开玩笑，时而不开玩笑的播客节目，由猫头鹰喜剧出品。\n\n公众号“猫头鹰喜剧”回复“听友群”，小助手会把你拉进群聊~",
          subscriptionCount: 139119,
          image: {
            picUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png",
            largePicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png@large",
            middlePicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png@middle",
            smallPicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png@small",
            thumbnailUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png@thumbnail",
          },
          color: {
            original: "#1A1A1A",
            light: "#ACACAC",
            dark: "#0B0B0B",
          },
          topicLabels: [],
          syncMode: "SELF_HOSTING",
          episodeCount: 132,
          latestEpisodePubDate: "2024-04-05T16:00:00.000Z",
          subscriptionStatus: "ON",
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
              uid: "61518a53e0f5e723bb520729",
              avatar: {
                picture: {
                  picUrl:
                    "https://image.xyzcdn.net/FnQ-E7VcqLbzqplvdVPGrQRGHmxC.jpg",
                  largePicUrl:
                    "https://image.xyzcdn.net/FnQ-E7VcqLbzqplvdVPGrQRGHmxC.jpg@large",
                  middlePicUrl:
                    "https://image.xyzcdn.net/FnQ-E7VcqLbzqplvdVPGrQRGHmxC.jpg@middle",
                  smallPicUrl:
                    "https://image.xyzcdn.net/FnQ-E7VcqLbzqplvdVPGrQRGHmxC.jpg@small",
                  thumbnailUrl:
                    "https://image.xyzcdn.net/FnQ-E7VcqLbzqplvdVPGrQRGHmxC.jpg@thumbnail",
                  format: "jpeg",
                  width: 3000,
                  height: 3000,
                },
              },
              nickname: "不开玩笑小助手",
              isNicknameSet: false,
              bio: "不开玩笑 Jokes Aside 由猫头鹰喜剧出品。",
              isCancelled: false,
              readTrackInfo: {},
              ipLoc: "上海",
              relation: "STRANGE",
              isBlockedByViewer: false,
            },
            // ...
          ],
          readTrackInfo: {},
          hasPopularEpisodes: true,
          contacts: [
            {
              type: "wechatOfficialAccounts",
              name: "猫头鹰喜剧",
              note: "订阅公众号",
            },
            {
              type: "weibo",
              name: "不开玩笑JokesAside",
              url: "https://weibo.com/n/不开玩笑JokesAside",
            },
            {
              type: "wechat",
              name: "JokesAside2024",
              note: "加听友群",
            },
            {
              type: "email",
              name: "bkwx@jieyanwenhua.com",
              note: "商务合作",
            },
          ],
          isCustomized: false,
        },
        isPlayed: true,
        isFinished: true,
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
        labels: [],
        sponsors: [],
        isCustomized: false,
        ipLoc: "上海",
      },
      //  ...
    ],
    loadNextKey: {
      pubDate: "2023-12-29T16:00:00.000Z",
      id: "658e6996bf3589e89468264d",
      direction: "NEXT",
    },
    loadMoreKey: {
      pubDate: "2023-12-29T16:00:00.000Z",
      id: "658e6996bf3589e89468264d",
      direction: "NEXT",
    },
    total: 132,
    order: "desc",
  },
}
```