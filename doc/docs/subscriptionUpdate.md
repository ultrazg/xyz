### 更新订阅

订阅/取消订阅节目

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

| 参数 | 必填 | 类型   | 说明                              |
| :--- | :--- | :----- | --------------------------------- |
| pid  | true | string | 节目 id                           |
| mode | true | string | 订阅为 **ON**，取消订阅为 **OFF** |

#### 返回字段

| 返回字段           | 类型    | 说明                                                |
| :----------------- | :------ | :-------------------------------------------------- |
| type               | string  | 类别。节目、单集等，详情看type对应的类别说明文档    |
| pid                | string  | 节目id                                              |
| eid                | string  | 单集id                                              |
| title              | string  | 标题                                                |
| author             | string  | 作者                                                |
| shownotes          | string  | 注释                                                |
| description        | string  | 描述                                                |
| media              | object  | 播客音频信息（大小、链接等）                        |
| podcasters         | array   | 播客嘉宾的信息（头像、昵称和uid等）                 |
| subscriptionStatus | string  | 订阅状态。**ON** 为**已订阅**，**OFF** 为**未订阅** |
| subscriptionStar   | boolean | 是否为星标订阅                                      |
| ...                | ...     | ...                                                 |


#### 示例

> 地址：https://www.example.com/subscription_update

参数

```javascript
{
  "pid": "61791d921989541784257779",
  "mode": "ON"
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
        // ...
      ],
      description:
        "一档时而开玩笑，时而不开玩笑的播客节目，...",
      episodeCount: 170,
      hasPopularEpisodes: true,
      image: {
        largePicUrl:
          "https://bts-image.xyzcdn.net/...",
        middlePicUrl:
          "https://bts-image.xyzcdn.net/...",
        picUrl:
          "https://bts-image.xyzcdn.net/...",
        smallPicUrl:
          "https://bts-image.xyzcdn.net/...",
        thumbnailUrl:
          "https://bts-image.xyzcdn.net/...",
      },
      isCustomized: false,
      latestEpisodePubDate: "2024-12-20T16:00:00.000Z",
      payEpisodeCount: 0,
      payType: "FREE",
      permissions: [
        {
          name: "SHARE",
          status: "PERMITTED",
        },
        // ...
      ],
      pid: "61791d921989541784257779",
      podcasters: [
        {
          avatar: {
            picture: {
              format: "jpeg",
              height: 3000,
              largePicUrl:
                "https://image.xyzcdn.net/...",
              middlePicUrl:
                "https://image.xyzcdn.net/...",
              picUrl:
                "https://image.xyzcdn.net/...",
              smallPicUrl:
                "https://image.xyzcdn.net/...",
              thumbnailUrl:
                "https://image.xyzcdn.net/...",
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
      showZhuiguangIcon: false,
      status: "NORMAL",
      subscriptionCount: 263184,
      subscriptionPush: true,
      subscriptionPushPriority: "HIGH",
      subscriptionStar: false,
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