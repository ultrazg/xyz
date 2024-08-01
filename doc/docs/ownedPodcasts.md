### 查询用户创建的播客节目

根据 uid 查询用户创建的播客节目

#### 请求地址

> /owned_podcasts

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

参考「查询节目详情」接口文档


#### 示例

> 地址：https://www.example.com/owned_podcasts

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
        author: "有知有行",
        brief: "有知有行出品的投资与生活向节目",
        color: {
          dark: "#BA3B2E",
          light: "#BA3B2E",
          original: "#BA3B2E",
        },
        contacts: [
          {
            name: "有知有行服务号",
            note: "订阅公众号",
            type: "wechatOfficialAccounts",
          },
          // ...
        ],
        description: "《知行小酒馆》是有知...",
        episodeCount: 161,
        hasPopularEpisodes: true,
        hasTopic: true,
        image: {
          largePicUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZzbzZaUEhTaTYyZVpKT0xob3JjcXB4OFRFd3YuanBn.jpg@large",
          // ...
        },
        isCustomized: false,
        latestEpisodePubDate: "2024-07-26T12:00:00.000Z",
        payEpisodeCount: 0,
        payType: "FREE",
        permissions: [
          {
            name: "SHARE",
            status: "PERMITTED",
          },
        ],
        pid: "6013f9f58e2f7ee375cf4216",
        podcasters: [
          {
            avatar: {
              picture: {
                format: "jpeg",
                height: 720,
                largePicUrl:
                  "https://image.xyzcdn.net/FgB16RE8wTj6eKMglFjszvG2i3Rl.jpg@large",
                // ...
                width: 720,
              },
            },
            bio: "关注投资，更关注怎样更好地生活",
            gender: "MALE",
            ipLoc: "北京",
            isBlockedByViewer: false,
            isCancelled: false,
            isNicknameSet: false,
            nickname: "有知有行",
            readTrackInfo: {},
            relation: "STRANGE",
            type: "USER",
            uid: "5fa391a5e0f5e723bbd34c78",
          },
          // ...
        ],
        readTrackInfo: {},
        status: "NORMAL",
        subscriptionCount: 695120,
        subscriptionPush: false,
        subscriptionPushPriority: "HIGH",
        subscriptionStar: false,
        subscriptionStatus: "OFF",
        syncMode: "SELF_HOSTING",
        title: "知行小酒馆",
        topicLabels: ["ALL"],
        type: "PODCAST",
      },
      // ...
    ],
  },
  msg: "OK",
}
```