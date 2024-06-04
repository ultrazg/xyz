### 更新星标订阅

更新我的星标订阅

#### 请求地址

> /subscription_update

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数     | 必填 | 类型    | 说明             |
| :------- | :--- | :------ | ---------------- |
| pid      | true | string  | 节目 id          |
| withStar | true | boolean | 是否加入星标订阅 |

#### 返回字段

| 返回字段    | 类型   | 说明                                             |
| :---------- | :----- | :----------------------------------------------- |
| type        | string | 类别。节目、单集等，详情看type对应的类别说明文档 |
| pid         | string | 节目id                                           |
| eid         | string | 单集id                                           |
| title       | string | 标题                                             |
| author      | string | 作者                                             |
| shownotes   | string | 注释                                             |
| description | string | 描述                                             |
| media       | object | 播客音频信息（大小、链接等）                     |
| podcasters  | array  | 播客嘉宾的信息（头像、昵称和uid等）              |
| ...         | ...    | ...                                              |


#### 示例

> 地址：https://www.example.com/subscription_update

参数

```javascript
{
  "pid":"61791d921989541784257779",
  "withStar":true
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: {
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
        {
          name: "不开玩笑JokesAside",
          type: "weibo",
          url: "https://weibo.com/n/不开玩笑JokesAside",
        },
        {
          name: "JokesAside2024",
          note: "加听友群",
          type: "wechat",
        },
        {
          name: "liulu@jieyanwenhua.com",
          note: "商务合作",
          type: "email",
        },
      ],
      description:
        "一个时而开玩笑，时而不开玩笑的播客节目，由猫头鹰喜剧出品...",
      episodeCount: 140,
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
      subscriptionCount: 158167,
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
  msg: "OK",
}
```