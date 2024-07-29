### 我的订阅

查询订阅节目，如果传入 uid 则查询该 id 用户的订阅

#### 请求地址

> /subscription

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数 | 必填 | 类型   | 说明          |
| :--- | :--- | :----- | ------------- |
| uid  | false | string | 如果查询已登录用户的订阅，为空即可。如果需要查询其他用户则需要传用户的 uid |
| loadMoreKey | false | object | 分页查询的条件，由接口返回 |

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
|loadMoreKey|object|如果存在下一页，则会返回该对象。将这个对象传入请求参数，便可实现分页查询|
| ...         | ...    | ...                                              |


#### 示例

> 地址：https://www.example.com/subscription

响应
```javascript
{
  code: 200,
  msg: "OK",
  data: {
    data: [
      {
        type: "PODCAST",
        pid: "5f4486d19504bbdb7729ea07",
        title: "话仙桃",
        author: "晃抠你",
        brief: "跟日常生活、居住、城市、文化等",
        description: "「话仙」在闽南话里是闲散聊天的意思...",
        subscriptionCount: 24190,
        image: {
          picUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZnLXRKaE9GcDY1ZlMyUnFOR2IzaGRfN2Q2blEucG5n.png",
          largePicUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZnLXRKaE9GcDY1ZlMyUnFOR2IzaGRfN2Q2blEucG5n.png@large",
          middlePicUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZnLXRKaE9GcDY1ZlMyUnFOR2IzaGRfN2Q2blEucG5n.png@middle",
          smallPicUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZnLXRKaE9GcDY1ZlMyUnFOR2IzaGRfN2Q2blEucG5n.png@small",
          thumbnailUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZnLXRKaE9GcDY1ZlMyUnFOR2IzaGRfN2Q2blEucG5n.png@thumbnail",
        },
        color: {
          original: "#FEDFE1",
          light: "#FC8997",
          dark: "#FC8997",
        },
        hasTopic: true,
        topicLabels: ["ALL"],
        syncMode: "SELF_HOSTING",
        episodeCount: 35,
        latestEpisodePubDate: "2024-04-10T10:49:04.190Z",
        subscriptionStatus: "ON",
        subscriptionPush: true,
        subscriptionPushPriority: "HIGH",
        subscriptionStar: false,
        status: "NORMAL",
        permissions: [
          {
            name: "SHARE",
            status: "PERMITTED",
          },
        ],
        payType: "PAY_EPISODE_PODCAST",
        payEpisodeCount: 1,
        podcasters: [
          {
            type: "USER",
            uid: "5e7c95bca2f2e503cec3c6a3",
            avatar: {
              picture: {
                picUrl: "https://image.xyzcdn.net/FlGBP9GnG3fp0RcuSX5j-3Ac-LGb",
                largePicUrl:
                  "https://image.xyzcdn.net/FlGBP9GnG3fp0RcuSX5j-3Ac-LGb@large",
                middlePicUrl:
                  "https://image.xyzcdn.net/FlGBP9GnG3fp0RcuSX5j-3Ac-LGb@middle",
                smallPicUrl:
                  "https://image.xyzcdn.net/FlGBP9GnG3fp0RcuSX5j-3Ac-LGb@small",
                thumbnailUrl:
                  "https://image.xyzcdn.net/FlGBP9GnG3fp0RcuSX5j-3Ac-LGb@thumbnail",
                format: "jpeg",
                width: 132,
                height: 132,
              },
            },
            nickname: "晃抠你",
            isNicknameSet: false,
            bio: "小红书@晃抠你",
            gender: "FEMALE",
            isCancelled: false,
            readTrackInfo: {},
            ipLoc: "瑞士",
            relation: "STRANGE",
            isBlockedByViewer: false,
          },
        ],
        readTrackInfo: {},
        hasPopularEpisodes: true,
        contacts: [
          {
            type: "wechatOfficialAccounts",
            name: "晃抠你",
            note: "订阅公众号",
          },
          {
            type: "wechat",
            name: "huangconi",
            note: "添加微信",
          },
          {
            type: "email",
            name: "hswconnie@gmail.com",
            note: "商务合作",
          },
        ],
        isCustomized: false,
      },
      // ...
    ],
  },
}
```