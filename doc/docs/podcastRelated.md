### 相关节目推荐（可能感兴趣的节目）

根据当前的节目推荐其他相似的节目

#### 请求地址

> /podcast_related

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数 | 必填 | 类型   | 说明    |
| :--- | :--- | :----- | ------- |
| pid  | true | string | 节目 id |

#### 返回字段

| 返回字段          | 类型   | 说明                                               |
| :---------------- | :----- | :------------------------------------------------- |
| type              | string | 类别。节目、单集等，详情看 type 对应的类别说明文档 |
| title             | string | 节目标题                                           |
| subscriptionCount | number | 节目订阅数                                         |
| pid               | string | 节目 id                                            |
| episodeCount      | number | 节目内单集数                                       |
| description       | string | 描述                                               |
| podcasters        | array  | 播客嘉宾的信息（头像、昵称和 uid 等）              |
| image             | object | 节目封面                                           |
| color             | object | 节目主题色                                         |
| ...               | ...    | ...                                                |


#### 示例

> 地址：https://www.example.com/podcast_related

参数

```javascript
{
  "pid": "61791d921989541784257779"
}
```

响应

```javascript
{
  code: 200,
  data: {
    data: [
      {
        podcast: {
          author: "惊讶喜剧",
          color: {
            dark: "#EC8D4C",
            light: "#EC8D4C",
            original: "#DA976C",
          },
          contacts: [],
          description: "欢迎收听由「惊讶喜剧」出品的《正经叭叭》...",
          episodeCount: 153,
          hasPopularEpisodes: true,
          hasTopic: true,
          image: {
            largePicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZySkpTV2RwclpHMElqX2Zqd3JaTUNCVFJ2aEU=@large",
            middlePicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZySkpTV2RwclpHMElqX2Zqd3JaTUNCVFJ2aEU=@middle",
            picUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZySkpTV2RwclpHMElqX2Zqd3JaTUNCVFJ2aEU=",
            smallPicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZySkpTV2RwclpHMElqX2Zqd3JaTUNCVFJ2aEU=@small",
            thumbnailUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZySkpTV2RwclpHMElqX2Zqd3JaTUNCVFJ2aEU=@thumbnail",
          },
          isCustomized: false,
          latestEpisodePubDate: "2024-05-06T15:31:06.042Z",
          payEpisodeCount: 0,
          payType: "FREE",
          permissions: [
            {
              name: "SHARE",
              status: "PERMITTED",
            },
          ],
          pid: "60e43cecc4e7c8188c2f92a4",
          podcasters: [
            {
              avatar: {
                picture: {
                  format: "jpeg",
                  height: 1497,
                  largePicUrl:
                    "https://image.xyzcdn.net/Fhdh1pyrLZwUqxU-n-2b47QD9MqV.jpg@large",
                  middlePicUrl:
                    "https://image.xyzcdn.net/Fhdh1pyrLZwUqxU-n-2b47QD9MqV.jpg@middle",
                  picUrl:
                    "https://image.xyzcdn.net/Fhdh1pyrLZwUqxU-n-2b47QD9MqV.jpg",
                  smallPicUrl:
                    "https://image.xyzcdn.net/Fhdh1pyrLZwUqxU-n-2b47QD9MqV.jpg@small",
                  thumbnailUrl:
                    "https://image.xyzcdn.net/Fhdh1pyrLZwUqxU-n-2b47QD9MqV.jpg@thumbnail",
                  width: 1497,
                },
              },
              bio: "处于支配地位的M",
              gender: "MALE",
              ipLoc: "北京",
              isBlockedByViewer: false,
              isCancelled: false,
              isNicknameSet: false,
              nickname: "大老王王王",
              readTrackInfo: {},
              relation: "STRANGE",
              type: "USER",
              uid: "5f4ece1ce0f5e723bb16ca85",
            },
            //  ...
          ],
          readTrackInfo: {
            recDescType: "EDITOR_COLLECTION",
            recallPolicy: "SIMILAR_PODCAST",
            score: "0.9619",
          },
          status: "NORMAL",
          subscriptionCount: 102896,
          subscriptionPush: false,
          subscriptionPushPriority: "HIGH",
          subscriptionStar: false,
          subscriptionStatus: "OFF",
          syncMode: "SELF_HOSTING",
          title: "正经叭叭",
          topicLabels: ["ALL"],
          type: "PODCAST",
        },
        recommendation: "“惊讶喜剧”出品，现场录制脱口秀",
      },
      // ...
    ],
  },
  msg: "OK",
}
```