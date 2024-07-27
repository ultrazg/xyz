### 首页-刷新「大家都在听」推荐

首页「大家都在听」推荐刷新

#### 请求地址

> /refresh_episode_commend

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

无

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

> 地址：https://www.example.com/refresh_episode_recommend

参数

```javascript
无
```

响应

```javascript
{
  data: {
    title: "大家都在听",
    moduleType: "HER_BACKUP_MULTI_RECALL",
    targetType: "EPISODE",
    target: [
      {
        episode: {
          type: "EPISODE",
          eid: "669f388b8fcadceb9045802d",
          pid: "61caf0508bb4cd867fcabe15",
          title: "恋爱经济学：宁做算盘精，不做冤大头？",
          shownotes: "<p>本期我们来聊聊敏感话题，爱情和钱的关系。</p>...",
          description: "本期我们来聊聊敏感话题，爱情和钱的关系。...",
          enclosure: {
            url: "https://media.xyzcdn.net/lg3WVFIKhp1ag3cPJFeLJay8d1s9.m4a",
          },
          isPrivateMedia: false,
          mediaKey: "lg3WVFIKhp1ag3cPJFeLJay8d1s9.m4a",
          media: {
            id: "lg3WVFIKhp1ag3cPJFeLJay8d1s9.m4a",
            size: 50952770,
            mimeType: "audio/mp4",
            source: {
              mode: "PUBLIC",
              url: "https://media.xyzcdn.net/lg3WVFIKhp1ag3cPJFeLJay8d1s9.m4a",
            },
          },
          clapCount: 239,
          commentCount: 698,
          playCount: 15205,
          favoriteCount: 172,
          pubDate: "2024-07-23T16:00:00.000Z",
          status: "NORMAL",
          duration: 3849,
          podcast: {
            type: "PODCAST",
            pid: "61caf0508bb4cd867fcabe15",
            title: "M字闲聊",
            author: "庄尼走路",
            brief: "闲聊关于爱情，城市，音乐或八卦",
            description:
              "闲聊关于爱情，城市，音乐或八卦，希望可以陪你上班，也伴你睡着",
            subscriptionCount: 40528,
            image: {
              picUrl:
                "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L2xrR0xDZmYzS0thM2c2UE5WUkJoVnlYUlVFeUEuanBn.jpg",
              largePicUrl:
                "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L2xrR0xDZmYzS0thM2c2UE5WUkJoVnlYUlVFeUEuanBn.jpg@large",
              middlePicUrl:
                "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L2xrR0xDZmYzS0thM2c2UE5WUkJoVnlYUlVFeUEuanBn.jpg@middle",
              smallPicUrl:
                "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L2xrR0xDZmYzS0thM2c2UE5WUkJoVnlYUlVFeUEuanBn.jpg@small",
              thumbnailUrl:
                "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L2xrR0xDZmYzS0thM2c2UE5WUkJoVnlYUlVFeUEuanBn.jpg@thumbnail",
            },
            color: {
              original: "#5BD435",
              light: "#67D249",
              dark: "#67D249",
            },
            topicLabels: [],
            syncMode: "SELF_HOSTING",
            episodeCount: 44,
            latestEpisodePubDate: "2024-07-23T16:00:00.000Z",
            subscriptionStatus: "OFF",
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
                uid: "5ff4363ae0f5e723bbb08668",
                avatar: {
                  picture: {
                    picUrl:
                      "https://image.xyzcdn.net/FhjNXF5b2Ez0NR7kD-he-PvqNb1-.jpg",
                    largePicUrl:
                      "https://image.xyzcdn.net/FhjNXF5b2Ez0NR7kD-he-PvqNb1-.jpg@large",
                    middlePicUrl:
                      "https://image.xyzcdn.net/FhjNXF5b2Ez0NR7kD-he-PvqNb1-.jpg@middle",
                    smallPicUrl:
                      "https://image.xyzcdn.net/FhjNXF5b2Ez0NR7kD-he-PvqNb1-.jpg@small",
                    thumbnailUrl:
                      "https://image.xyzcdn.net/FhjNXF5b2Ez0NR7kD-he-PvqNb1-.jpg@thumbnail",
                    format: "jpeg",
                    width: 3000,
                    height: 3000,
                  },
                },
                nickname: "庄尼走路",
                isNicknameSet: true,
                bio: "烟霞闲骨格，泉石野生涯",
                gender: "MALE",
                isCancelled: false,
                readTrackInfo: {},
                ipLoc: "山东",
                relation: "STRANGE",
                isBlockedByViewer: false,
              },
              //  ...
            ],
            readTrackInfo: {},
            hasPopularEpisodes: true,
            contacts: [
              {
                type: "weibo",
                name: "M字闲聊播客",
                url: "https://weibo.com/n/M字闲聊播客",
              },
              {
                type: "email",
                name: "mzonepod@163.com",
                note: "听众反馈",
              },
            ],
            isCustomized: false,
          },
          isPlayed: false,
          isFinished: false,
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
          readTrackInfo: {
            recallPolicy: "HORIREC_EPISODE_HOT_PLAYED_USER_COUNT",
            moduleIdx: "7",
            inModuleIdx: "0",
            additionalInfo:
              '{"recModuleType":"HER_BACKUP_MULTI_RECALL","refId":"","ref":""}',
            backendRecDescType: "commentCount",
          },
          labels: [],
          sponsors: [],
          isCustomized: false,
          ipLoc: "浙江",
          transcript: {
            mediaId: "lg3WVFIKhp1ag3cPJFeLJay8d1s9.m4a",
          },
        },
        recommendation: "评论数 99+",
        relatedUsers: [],
        hasNegativeFeedback: false,
      },
      // ...
    ],
    lightBlueEndIndex: 2,
    feedback: {
      canFeedback: true,
      feedbackSource: "DISCOVERY_EPISODE_RECOMMEND",
    },
    readTrackInfo: {
      recModuleType: "HER_BACKUP_MULTI_RECALL",
      dataDate: "20240727",
      recDate: "20240727",
      moduleIdx: "7",
      recSize: "3",
    },
  },
}
```