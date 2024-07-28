### 查询收听排行

查询「我的个人」收听排行榜单

#### 请求地址

> /mileage_list

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数 | 必填  | 类型    | 说明                                                         |
| :--- | :---- | :------ | ------------------------------------------------------------ |
| all  | false | boolean | 参数为 **false** 表示查询过去三十天的记录，为 **true** 表示查询所有记录。默认为 **false** |

#### 返回字段

| 返回字段      | 类型   | 说明       |
| :------------ | :----- | :--------- |
| playedSeconds | number | 总收听时长 |
| podcast       | object | 节目的详情 |


#### 示例

> 地址：https://www.example.com/mileage_list

参数

```javascript
{
  "all": true
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: [
      {
        playedSeconds: 480508,
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
            "一个时而开玩笑，时而不开玩笑的播客节目，由猫头鹰喜剧出品...",
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
          subscriptionCount: 176460,
          subscriptionPush: false,
          subscriptionPushPriority: "HIGH",
          subscriptionStar: true,
          subscriptionStatus: "ON",
          syncMode: "SELF_HOSTING",
          title: "不开玩笑 Jokes Aside",
          topicLabels: [],
          type: "PODCAST",
        },
      },
      // ...
    ],
  },
  msg: "OK",
}
```