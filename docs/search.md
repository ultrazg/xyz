### 搜索

搜索内容
> 目前暂不支持指定类别搜索以及分页

#### 请求地址

> /search

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数    | 必填 | 类型   | 说明                       |
| :------ | :--- | :----- | -------------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数    | 必填 | 类型   | 说明                       |
| :------ | :--- | :----- | -------------------------- |
| keyword | true | string | 要搜索的内容，支持模糊查询 |

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
|...|...|...|


#### 示例

> 地址：https://www.example.com/search

参数

```javascript
{
  "keyword": "小米"
}
```

响应

``` javascript
{
  code: 200,
  msg: "OK",
  data: {
    data: [
      {
        type: "HEADER",
        title: "节目",
        link: "cosmos://page.cos/search?tab=podcast",
        readTrackInfo: {
          searchId: "1713588225610",
        },
      },
      {
        type: "PODCAST",
        pid: "64b770b0412b0153ea3c1da9",
        title: "小米的散碎时光铺",
        author: "刘婷匠逸庭",
        description: "",
        subscriptionCount: 3,
        image: {
          picUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L2xzOC1qdGVsUktOcHBKamgtZkFuX2djRjhXUG8uanBn.jpg",
          largePicUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L2xzOC1qdGVsUktOcHBKamgtZkFuX2djRjhXUG8uanBn.jpg@large",
          middlePicUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L2xzOC1qdGVsUktOcHBKamgtZkFuX2djRjhXUG8uanBn.jpg@middle",
          smallPicUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L2xzOC1qdGVsUktOcHBKamgtZkFuX2djRjhXUG8uanBn.jpg@small",
          thumbnailUrl:
            "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L2xzOC1qdGVsUktOcHBKamgtZkFuX2djRjhXUG8uanBn.jpg@thumbnail",
        },
        color: {
          original: "#CDCAC0",
          light: "#E3C14A",
          dark: "#E3C14A",
        },
        topicLabels: [],
        syncMode: "SELF_HOSTING",
        episodeCount: 5,
        latestEpisodePubDate: "2024-02-29T02:52:45.168Z",
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
            uid: "6346c0a4edce67104a98ce33",
            avatar: {
              picture: {
                picUrl:
                  "https://image.xyzcdn.net/FqUrrUGD1YIaeVqrpn9MV0yPz_iY.jpg",
                largePicUrl:
                  "https://image.xyzcdn.net/FqUrrUGD1YIaeVqrpn9MV0yPz_iY.jpg@large",
                middlePicUrl:
                  "https://image.xyzcdn.net/FqUrrUGD1YIaeVqrpn9MV0yPz_iY.jpg@middle",
                smallPicUrl:
                  "https://image.xyzcdn.net/FqUrrUGD1YIaeVqrpn9MV0yPz_iY.jpg@small",
                thumbnailUrl:
                  "https://image.xyzcdn.net/FqUrrUGD1YIaeVqrpn9MV0yPz_iY.jpg@thumbnail",
                format: "jpeg",
                width: 1371,
                height: 1371,
              },
            },
            nickname: "小米球儿",
            isNicknameSet: true,
            bio: "",
            gender: "FEMALE",
            isCancelled: false,
            readTrackInfo: {
              searchId: "1713588225610",
            },
            ipLoc: "北京",
            relation: "STRANGE",
            isBlockedByViewer: false,
          },
        ],
        readTrackInfo: {
          searchId: "1713588225610",
          recallPolicy: "POD_SubscriptionTimeDecay",
          recallScore: "1100.5139",
        },
        hasPopularEpisodes: false,
        contacts: [],
        isCustomized: false,
        truncatedDescription: "",
      },
      // ...
      {
        type: "FOOTER",
        title: "查看全部",
        link: "cosmos://page.cos/search?tab=podcast",
        readTrackInfo: {
          searchId: "1713588225610",
        },
      },
      {
        type: "HEADER",
        title: "单集",
        link: "cosmos://page.cos/search?tab=episode",
        readTrackInfo: {
          searchId: "1713588225610",
        },
      },
      {
        type: "EPISODE",
        eid: "661153124f66d1c1da58cbae",
        pid: "636669d51064cb55f31505fc",
        title: "小米SU7营销复盘：你所知道的为什么都是错的-Vol 46",
        shownotes:
          "<p>本期节目关注风口上的小米汽车，主播借助在营销、产品上的经验解答</p>...",
        description: "本期节目关注风口上的小米汽车...",
        enclosure: {
          url: "https://media.xyzcdn.net/lpKUQpQv73JHU8bXuk3G953r0Yyx.m4a",
        },
        isPrivateMedia: false,
        mediaKey: "lpKUQpQv73JHU8bXuk3G953r0Yyx.m4a",
        media: {
          id: "lpKUQpQv73JHU8bXuk3G953r0Yyx.m4a",
          size: 84639099,
          mimeType: "audio/mp4",
          source: {
            mode: "PUBLIC",
            url: "https://media.xyzcdn.net/lpKUQpQv73JHU8bXuk3G953r0Yyx.m4a",
          },
        },
        clapCount: 173,
        commentCount: 307,
        playCount: 34115,
        favoriteCount: 366,
        pubDate: "2024-04-06T13:50:09.931Z",
        status: "NORMAL",
        duration: 5233,
        podcast: {
          type: "PODCAST",
          pid: "636669d51064cb55f31505fc",
          title: "脑放电波",
          author: "托马斯_T49b",
          brief: "关注科技前沿、品牌营销和个人成长",
          description: "脑放电波是一档关注科技前沿...",
          subscriptionCount: 22795,
          image: {
            picUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z2RTZMYmpjWGh0NVQ3cHJoVTU5U1FOcW5UZEIucG5n.png",
            largePicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z2RTZMYmpjWGh0NVQ3cHJoVTU5U1FOcW5UZEIucG5n.png@large",
            middlePicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z2RTZMYmpjWGh0NVQ3cHJoVTU5U1FOcW5UZEIucG5n.png@middle",
            smallPicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z2RTZMYmpjWGh0NVQ3cHJoVTU5U1FOcW5UZEIucG5n.png@small",
            thumbnailUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z2RTZMYmpjWGh0NVQ3cHJoVTU5U1FOcW5UZEIucG5n.png@thumbnail",
          },
          color: {
            original: "#0604F8",
            light: "#2452CF",
            dark: "#2452CF",
          },
          hasTopic: true,
          topicLabels: ["ALL"],
          syncMode: "SELF_HOSTING",
          episodeCount: 46,
          latestEpisodePubDate: "2024-04-06T13:50:09.931Z",
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
              uid: "62860db1edce67104ad0a3f4",
              avatar: {
                picture: {
                  picUrl:
                    "https://image.xyzcdn.net/FlPt9N-TQWK3zdDykg84RP47vROG.jpg",
                  largePicUrl:
                    "https://image.xyzcdn.net/FlPt9N-TQWK3zdDykg84RP47vROG.jpg@large",
                  middlePicUrl:
                    "https://image.xyzcdn.net/FlPt9N-TQWK3zdDykg84RP47vROG.jpg@middle",
                  smallPicUrl:
                    "https://image.xyzcdn.net/FlPt9N-TQWK3zdDykg84RP47vROG.jpg@small",
                  thumbnailUrl:
                    "https://image.xyzcdn.net/FlPt9N-TQWK3zdDykg84RP47vROG.jpg@thumbnail",
                  format: "jpeg",
                  width: 1500,
                  height: 1500,
                },
              },
              nickname: "托马斯白",
              isNicknameSet: true,
              bio: "",
              gender: "MALE",
              isCancelled: false,
              readTrackInfo: {
                searchId: "1713588225610",
              },
              ipLoc: "广东",
              relation: "STRANGE",
              isBlockedByViewer: false,
            },
            {
              type: "USER",
              uid: "635cfc85edce67104a05bd11",
              avatar: {
                picture: {
                  picUrl:
                    "https://image.xyzcdn.net/FhoiAyauNu8_nXa_naLbuA_OBAEO.jpg",
                  largePicUrl:
                    "https://image.xyzcdn.net/FhoiAyauNu8_nXa_naLbuA_OBAEO.jpg@large",
                  middlePicUrl:
                    "https://image.xyzcdn.net/FhoiAyauNu8_nXa_naLbuA_OBAEO.jpg@middle",
                  smallPicUrl:
                    "https://image.xyzcdn.net/FhoiAyauNu8_nXa_naLbuA_OBAEO.jpg@small",
                  thumbnailUrl:
                    "https://image.xyzcdn.net/FhoiAyauNu8_nXa_naLbuA_OBAEO.jpg@thumbnail",
                  format: "jpeg",
                  width: 3000,
                  height: 3000,
                },
              },
              nickname: "田达蒙",
              isNicknameSet: true,
              bio: "",
              gender: "MALE",
              isCancelled: false,
              readTrackInfo: {
                searchId: "1713588225610",
              },
              ipLoc: "北京",
              relation: "STRANGE",
              isBlockedByViewer: false,
            },
            {
              type: "USER",
              uid: "6023ccc6e0f5e723bb740342",
              avatar: {
                picture: {
                  picUrl:
                    "https://image.xyzcdn.net/Fs_8f5QMm_XL1YLZptINYf30ih7d.jpg",
                  largePicUrl:
                    "https://image.xyzcdn.net/Fs_8f5QMm_XL1YLZptINYf30ih7d.jpg@large",
                  middlePicUrl:
                    "https://image.xyzcdn.net/Fs_8f5QMm_XL1YLZptINYf30ih7d.jpg@middle",
                  smallPicUrl:
                    "https://image.xyzcdn.net/Fs_8f5QMm_XL1YLZptINYf30ih7d.jpg@small",
                  thumbnailUrl:
                    "https://image.xyzcdn.net/Fs_8f5QMm_XL1YLZptINYf30ih7d.jpg@thumbnail",
                  format: "jpeg",
                  width: 3000,
                  height: 3000,
                },
              },
              nickname: "Nixon_Hu",
              isNicknameSet: true,
              bio: "联系我 hxk2312",
              isCancelled: false,
              readTrackInfo: {
                searchId: "1713588225610",
              },
              ipLoc: "山东",
              relation: "STRANGE",
              isBlockedByViewer: false,
            },
          ],
          readTrackInfo: {
            searchId: "1713588225610",
          },
          hasPopularEpisodes: true,
          contacts: [
            {
              type: "wechatOfficialAccounts",
              name: "托马斯白",
              note: "订阅公众号",
            },
            {
              type: "jike",
              name: "托马斯白",
              url: "https://m.okjike.com/users/77C28719-0AE5-4118-8CF6-36886802B208",
            },
            {
              type: "wechat",
              name: "BrainAmp01",
              note: "加听友群",
            },
            {
              type: "email",
              name: "alanneo@me.com",
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
          {
            name: "COMMENT",
            status: "PERMITTED",
          },
          {
            name: "VOICE_COMMENT",
            status: "DENIED",
          },
          {
            name: "COMMENT_PAGE",
            status: "PERMITTED",
          },
          {
            name: "CLAP",
            status: "PERMITTED",
          },
          {
            name: "PICK",
            status: "PERMITTED",
          },
          {
            name: "VOTE",
            status: "PERMITTED",
          },
        ],
        payType: "FREE",
        wechatShare: {
          style: "LINK_CARD",
        },
        readTrackInfo: {
          searchId: "1713588225610",
          recallPolicy: "EPI_EDITOR_SELECTED",
          recallScore: "0.0000",
          searchAdvertisementId: "66154f8c5f18b9dd6e1b7773",
        },
        labels: [
          {
            name: "最受欢迎",
            code: "POPULAR",
          },
        ],
        sponsors: [],
        isCustomized: false,
        ipLoc: "北京",
        truncatedDescription: "本期节目关注风口上的小米汽车...",
      },
      // ...
      {
        type: "FOOTER",
        title: "查看全部",
        link: "cosmos://page.cos/search?tab=episode",
        readTrackInfo: {
          searchId: "1713588225610",
        },
      },
      {
        type: "HEADER",
        title: "用户",
        link: "cosmos://page.cos/search?tab=user",
        readTrackInfo: {
          searchId: "1713588225610",
        },
      },
      {
        type: "SEARCHED_USERS",
        users: [
          {
            type: "USER",
            uid: "5e7caa45a2f2e503cec3e203",
            avatar: {
              picture: {
                picUrl: "https://image.xyzcdn.net/FuUWsV2PMyQBGisMGFuH5iTz_Vpt",
                largePicUrl:
                  "https://image.xyzcdn.net/FuUWsV2PMyQBGisMGFuH5iTz_Vpt@large",
                middlePicUrl:
                  "https://image.xyzcdn.net/FuUWsV2PMyQBGisMGFuH5iTz_Vpt@middle",
                smallPicUrl:
                  "https://image.xyzcdn.net/FuUWsV2PMyQBGisMGFuH5iTz_Vpt@small",
                thumbnailUrl:
                  "https://image.xyzcdn.net/FuUWsV2PMyQBGisMGFuH5iTz_Vpt@thumbnail",
                format: "jpeg",
                width: 132,
                height: 132,
              },
            },
            nickname: "小米",
            isNicknameSet: false,
            isCancelled: false,
            readTrackInfo: {
              searchId: "1713588225610",
              recallPolicy: "USER_Match",
              recallScore: "103900.0000",
            },
            ipLoc: "广东",
            relation: "STRANGE",
            isBlockedByViewer: false,
          },
          {
            type: "USER",
            uid: "5fb9346be0f5e723bb12f862",
            avatar: {
              picture: {
                picUrl: "https://image.xyzcdn.net/Fo4xvk1XtpoktwZbWRpEZb_gzDUO",
                largePicUrl:
                  "https://image.xyzcdn.net/Fo4xvk1XtpoktwZbWRpEZb_gzDUO@large",
                middlePicUrl:
                  "https://image.xyzcdn.net/Fo4xvk1XtpoktwZbWRpEZb_gzDUO@middle",
                smallPicUrl:
                  "https://image.xyzcdn.net/Fo4xvk1XtpoktwZbWRpEZb_gzDUO@small",
                thumbnailUrl:
                  "https://image.xyzcdn.net/Fo4xvk1XtpoktwZbWRpEZb_gzDUO@thumbnail",
                width: 663,
                height: 663,
                format: "png",
              },
            },
            nickname: "小米锅巴",
            isNicknameSet: false,
            gender: "MALE",
            isCancelled: false,
            readTrackInfo: {
              searchId: "1713588225610",
              recallPolicy: "USER_Match",
              recallScore: "10900.0000",
            },
            relation: "STRANGE",
            isBlockedByViewer: false,
          },
          {
            type: "USER",
            uid: "61c9d4df2cbd7c017640614b",
            avatar: {
              picture: {
                picUrl: "https://image.xyzcdn.net/Foggx8RQ-rTGZTYEGc6ahqO7P3IU",
                largePicUrl:
                  "https://image.xyzcdn.net/Foggx8RQ-rTGZTYEGc6ahqO7P3IU@large",
                middlePicUrl:
                  "https://image.xyzcdn.net/Foggx8RQ-rTGZTYEGc6ahqO7P3IU@middle",
                smallPicUrl:
                  "https://image.xyzcdn.net/Foggx8RQ-rTGZTYEGc6ahqO7P3IU@small",
                thumbnailUrl:
                  "https://image.xyzcdn.net/Foggx8RQ-rTGZTYEGc6ahqO7P3IU@thumbnail",
                format: "jpeg",
                width: 132,
                height: 132,
              },
            },
            nickname: "小米专卖店",
            isNicknameSet: true,
            gender: "FEMALE",
            isCancelled: false,
            readTrackInfo: {
              searchId: "1713588225610",
              recallPolicy: "USER_Match",
              recallScore: "10900.0000",
            },
            relation: "STRANGE",
            isBlockedByViewer: false,
          },
        ],
      },
    ],
    highlightWord: {
      words: ["小米"],
      singleMaxHighlightTime: 3,
    },
  },
}
```