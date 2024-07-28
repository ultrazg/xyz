### 查询收听历史

根据 uid 查询「个人主页」下的收听历史列表

#### 请求地址

> /played_list

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

| 返回字段             | 类型    | 说明       |
| :------------------- | :------ | :--------- |
| data                 | object  | 单集的信息 |
| isRecentPlayedHidden | boolean | 是否不可见 |


#### 示例

> 地址：https://www.example.com/played_list

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
        clapCount: 811,
        commentCount: 1081,
        description: "📝本期主播：\n* 小北（@小北指南）...",
        duration: 7386,
        eid: "666bcda3b6a8412729ae252b",
        enclosure: {
          url: "https://media.xyzcdn.net/luq2jmwY6WpgQSlOwXzRq-TWsmYr.m4a",
        },
        favoriteCount: 661,
        image: {
          largePicUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvV2ZVTms0VTVxWEJxUkF5SHNlU1djTTZzazgucG5n.png@large",
          // ...
        },
        ipLoc: "美国",
        isCustomized: false,
        isFavorited: false,
        isFinished: false,
        isPicked: false,
        isPlayed: true,
        isPrivateMedia: false,
        labels: [],
        media: {
          id: "luq2jmwY6WpgQSlOwXzRq-TWsmYr.m4a",
          mimeType: "audio/mp4",
          size: 119577967,
          source: {
            mode: "PUBLIC",
            url: "https://media.xyzcdn.net/luq2jmwY6WpgQSlOwXzRq-TWsmYr.m4a",
          },
        },
        mediaKey: "luq2jmwY6WpgQSlOwXzRq-TWsmYr.m4a",
        payType: "FREE",
        permissions: [
          {
            name: "SHARE",
            status: "PERMITTED",
          },
          // ...
        ],
        pid: "61791d921989541784257779",
        playCount: 138647,
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
          episodeCount: 148,
          hasPopularEpisodes: true,
          image: {
            largePicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png@large",
            // ...
          },
          isCustomized: false,
          latestEpisodePubDate: "2024-07-26T16:00:00.000Z",
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
            //  ...
          ],
          readTrackInfo: {},
          status: "NORMAL",
          subscriptionCount: 179180,
          subscriptionPush: false,
          subscriptionPushPriority: "HIGH",
          subscriptionStar: true,
          subscriptionStatus: "ON",
          syncMode: "SELF_HOSTING",
          title: "不开玩笑 Jokes Aside",
          topicLabels: [],
          type: "PODCAST",
        },
        pubDate: "2024-06-14T16:00:00.000Z",
        readTrackInfo: {},
        shownotes:
          "<p><strong>📝本期主播：</strong></p>\n<ul>\n  <li>小北（...",
        sponsors: [],
        status: "NORMAL",
        title: "140. 考试play：聊聊人生中的考试和考验",
        transcript: {
          mediaId: "luq2jmwY6WpgQSlOwXzRq-TWsmYr.m4a",
        },
        type: "EPISODE",
        wechatShare: {
          style: "LINK_CARD",
        },
      },
      // ...
    ],
    isRecentPlayedHidden: false,
  },
  msg: "OK",
}
```