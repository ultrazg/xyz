### 节目内「最受欢迎」单集列表

获取节目内「最受欢迎」的单集

#### 请求地址

> /episode_list_by_filter

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数 | 必填 | 类型   | 说明   |
| :--- | :--- | :----- | ------ |
| pid  | true | string | 节目id |

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

> 地址：https://www.example.com/episode_list_by_filter

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
        clapCount: 712,
        commentCount: 1717,
        description: "近年来，中国生育率急剧下降，...",
        duration: 2829,
        eid: "65af93e12b3170a65eb3faa4",
        enclosure: {
          url: "https://media.xyzcdn.net/lsn88lFJmwfo_JvxFswumTkE9YsJ.m4a",
        },
        favoriteCount: 1159,
        image: {
          largePicUrl: "https://bts-image.xyzcdn.net/...",
          middlePicUrl: "https://bts-image.xyzcdn.net/...",
          picUrl: "https://bts-image.xyzcdn.net/...",
          smallPicUrl: "https://bts-image.xyzcdn.net/...",
          thumbnailUrl: "https://bts-image.xyzcdn.net/...",
        },
        ipLoc: "上海",
        isCustomized: false,
        isFavorited: false,
        isFinished: false,
        isPicked: false,
        isPlayed: false,
        isPrivateMedia: false,
        labels: [
          {
            code: "POPULAR",
            name: "最受欢迎",
          },
        ],
        media: {
          id: "lsn88lFJmwfo_JvxFswumTkE9YsJ.m4a",
          mimeType: "audio/mp4",
          size: 45767972,
          source: {
            mode: "PUBLIC",
            url: "https://media.xyzcdn.net/lsn88lFJmwfo_JvxFswumTkE9YsJ.m4a",
          },
        },
        mediaKey: "lsn88lFJmwfo_JvxFswumTkE9YsJ.m4a",
        payType: "FREE",
        permissions: [
          {
            name: "SHARE",
            status: "PERMITTED",
          },
          // ...
        ],
        pid: "5e4ee557418a84a0466737b7",
        playCount: 139280,
        podcast: {
          author: "JustPod",
          brief: "基于经验视角提供内容的文化类播客",
          color: {
            dark: "#AF4D1F",
            light: "#AF4D1F",
            original: "#7A6B65",
          },
          contacts: [
            {
              name: "忽左忽右leftright",
              note: "订阅公众号",
              type: "wechatOfficialAccounts",
            },
            //  ...
          ],
          description:
            "「忽左忽右」是一档文化沙龙类播客节目，...",
          episodeCount: 465,
          hasPopularEpisodes: true,
          hasTopic: true,
          image: {
            format: "png",
            height: 1400,
            largePicUrl: "https://image.xyzcdn.net/...",
            middlePicUrl: "https://image.xyzcdn.net/...",
            picUrl: "https://image.xyzcdn.net/...",
            smallPicUrl: "https://image.xyzcdn.net/...",
            thumbnailUrl: "https://image.xyzcdn.net/...",
            width: 1400,
          },
          isCustomized: false,
          latestEpisodePubDate: "2024-12-24T10:07:47.136Z",
          payEpisodeCount: 34,
          payType: "PAY_EPISODE_PODCAST",
          permissions: [
            {
              name: "SHARE",
              status: "PERMITTED",
            },
            // ...
          ],
          pid: "5e4ee557418a84a0466737b7",
          podcasters: [
            {
              avatar: {
                picture: {
                  format: "jpeg",
                  height: 2000,
                  largePicUrl: "https://image.xyzcdn.net/...",
                  middlePicUrl: "https://image.xyzcdn.net/...",
                  picUrl: "https://image.xyzcdn.net/...",
                  smallPicUrl: "https://image.xyzcdn.net/...",
                  thumbnailUrl: "https://image.xyzcdn.net/...",
                  width: 2000,
                },
              },
              bio: "中国领先的播客与数字音频企业，探索中文播客的更多可能。",
              ipLoc: "上海",
              isBlockedByViewer: false,
              isCancelled: false,
              isNicknameSet: false,
              nickname: "JustPod",
              readTrackInfo: {},
              relation: "STRANGE",
              type: "USER",
              uid: "603773dde0f5e723bbae35ae",
            },
          ],
          readTrackInfo: {},
          showZhuiguangIcon: false,
          status: "NORMAL",
          subscriptionCount: 797078,
          subscriptionPush: false,
          subscriptionPushPriority: "HIGH",
          subscriptionStar: false,
          subscriptionStatus: "OFF",
          syncMode: "SELF_HOSTING",
          title: "忽左忽右",
          topicLabels: ["ALL"],
          type: "PODCAST",
        },
        pubDate: "2024-01-23T10:19:19.000Z",
        readTrackInfo: {},
        shownotes:
          '<div data-page-id="UeYEds2RPoaEwixPfBicqqmCnGh" data-docx-has-block-data="false">\n<div class=" old-record-id-Ea9wdBBR8oHqldxquzycusJvnje">近年来，中国生育率急剧下降，...',
        sponsors: [],
        status: "NORMAL",
        title: "299 最新人口数据公布，梁建章怎么看？",
        transcript: {
          mediaId: "lsn88lFJmwfo_JvxFswumTkE9YsJ.m4a",
        },
        type: "EPISODE",
        wechatShare: {
          style: "LINK_CARD",
        },
      },
    ],
  },
  msg: "OK",
}
```