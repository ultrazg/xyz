### 分类下的节目列表

根据标签获取分类下的节目列表

#### 请求地址

> /category_podcast_list

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求参数

| 参数           | 必填  | 类型    | 说明                                                     |
| -------------- | ----- | ------- | -------------------------------------------------------- |
| categoryId     | true  | string  | 分类的 id。由 `/category_list` 接口获取                  |
| tab            | true  | string  | 分类的标签。由 `/category_list_tab` 接口获取             |
| omitSubscribed | false | boolean | 是否已订阅。默认值 **false**                             |
| loadMoreKey    | false | number  | 分页查询的条件，由本接口返回，如有该字段则表示存在下一页 |

#### 返回字段

| 返回字段 | 类型   | 说明                                            |
| -------- | ------ | ----------------------------------------------- |
| episode  | object | 单集的详情。可查看 `/episode_detail` 接口的说明 |
| podcast  | object | 节目的详情。可查看 `/podcast_detail` 接口的说明 |




#### 示例

> 地址：https://www.example.com/category_podcast_list

参数

```javascript
{
  "categoryId": "63c76a8924b1622727bd321b",
  "omitSubscribed": false,
  "tab": "ALL",
  // 由本接口返回，传入该字段可查询下一页数据
  "loadMoreKey": 10
}
```

响应

``` javascript
{
  code: 200,
  data: {
    data: [
      {
        episode: {
          clapCount: 0,
          commentCount: 1,
          description: "过去三年中中国资本市场大火的赛道少不了慢病管理...",
          duration: 2358,
          eid: "66573b0aefabaae3a15ded93",
          enclosure: {
            url: "https://jt.ximalaya.com//GKwRINsKLPnpASNjQALaNIVy.m4a?channel=rss&album_id=56031677&track_id=731799131&uid=382519014&jt=https://audio.xmcdn.com/storages/c286-audiofreehighqps/B7/8C/GKwRINsKLPnpASNjQALaNIVy.m4a",
          },
          favoriteCount: 1,
          image: {
            largePicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy9mN2Q3LWF1ZGlvZnJlZWhpZ2hxcHMvRUQvNTQvR01Db09SNEtMUGdoQUFWRV9nTGFNcU91LmpwZWc=.jpeg@large",
            middlePicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy9mN2Q3LWF1ZGlvZnJlZWhpZ2hxcHMvRUQvNTQvR01Db09SNEtMUGdoQUFWRV9nTGFNcU91LmpwZWc=.jpeg@middle",
            picUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy9mN2Q3LWF1ZGlvZnJlZWhpZ2hxcHMvRUQvNTQvR01Db09SNEtMUGdoQUFWRV9nTGFNcU91LmpwZWc=.jpeg",
            smallPicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy9mN2Q3LWF1ZGlvZnJlZWhpZ2hxcHMvRUQvNTQvR01Db09SNEtMUGdoQUFWRV9nTGFNcU91LmpwZWc=.jpeg@small",
            thumbnailUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy9mN2Q3LWF1ZGlvZnJlZWhpZ2hxcHMvRUQvNTQvR01Db09SNEtMUGdoQUFWRV9nTGFNcU91LmpwZWc=.jpeg@thumbnail",
          },
          ipLoc: "北京",
          isCustomized: false,
          isFavorited: false,
          isFinished: false,
          isPicked: false,
          isPlayed: false,
          isPrivateMedia: false,
          labels: [],
          media: {
            id: "//GKwRINsKLPnpASNjQALaNIVy.m4a?channel=rss&album_id=56031677&track_id=731799131&uid=382519014&jt=https%3A%2F%2Faudio.xmcdn.com%2Fstorages%2Fc286-audiofreehighqps%2FB7%2F8C%2FGKwRINsKLPnpASNjQALaNIVy.m4a",
            mimeType: "audio/x-m4a",
            size: 37748655,
            source: {
              mode: "PUBLIC",
              url: "https://jt.ximalaya.com//GKwRINsKLPnpASNjQALaNIVy.m4a?channel=rss&album_id=56031677&track_id=731799131&uid=382519014&jt=https://audio.xmcdn.com/storages/c286-audiofreehighqps/B7/8C/GKwRINsKLPnpASNjQALaNIVy.m4a",
            },
          },
          mediaKey:
            "//GKwRINsKLPnpASNjQALaNIVy.m4a?channel=rss&album_id=56031677&track_id=731799131&uid=382519014&jt=https%3A%2F%2Faudio.xmcdn.com%2Fstorages%2Fc286-audiofreehighqps%2FB7%2F8C%2FGKwRINsKLPnpASNjQALaNIVy.m4a",
          payType: "FREE",
          permissions: [
            {
              name: "SHARE",
              status: "PERMITTED",
            },
            // ...
          ],
          pid: "61dc011a1f7c9738c060c944",
          playCount: 4,
          podcast: {
            author: "_朗姆酒蛋糕",
            brief: "浅显易懂地解释经济和商业知识",
            color: {
              dark: "#CE851C",
              light: "#CE851C",
              original: "#CC8629",
            },
            contacts: [
              {
                name: "S_Jimmy",
                note: "加听友群",
                type: "wechat",
              },
              {
                name: "kidpodcast824@gmail.com",
                note: "听众反馈",
                type: "email",
              },
            ],
            description: "我们希望用声音，...",
            episodeCount: 50,
            hasPopularEpisodes: true,
            image: {
              largePicUrl:
                "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy83OTllLWF1ZGlvZnJlZWhpZ2hxcHMvQkEvMkEvR0t3UklKSUYtMnl4QUFGYlp3RWtrQ1NsLmpwZWc=.jpeg@large",
              middlePicUrl:
                "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy83OTllLWF1ZGlvZnJlZWhpZ2hxcHMvQkEvMkEvR0t3UklKSUYtMnl4QUFGYlp3RWtrQ1NsLmpwZWc=.jpeg@middle",
              picUrl:
                "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy83OTllLWF1ZGlvZnJlZWhpZ2hxcHMvQkEvMkEvR0t3UklKSUYtMnl4QUFGYlp3RWtrQ1NsLmpwZWc=.jpeg",
              smallPicUrl:
                "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy83OTllLWF1ZGlvZnJlZWhpZ2hxcHMvQkEvMkEvR0t3UklKSUYtMnl4QUFGYlp3RWtrQ1NsLmpwZWc=.jpeg@small",
              thumbnailUrl:
                "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy83OTllLWF1ZGlvZnJlZWhpZ2hxcHMvQkEvMkEvR0t3UklKSUYtMnl4QUFGYlp3RWtrQ1NsLmpwZWc=.jpeg@thumbnail",
            },
            isCustomized: false,
            latestEpisodePubDate: "2024-05-29T14:07:51.000Z",
            payEpisodeCount: 0,
            payType: "FREE",
            permissions: [
              {
                name: "SHARE",
                status: "PERMITTED",
              },
            ],
            pid: "61dc011a1f7c9738c060c944",
            podcasters: [
              {
                avatar: {
                  picture: {
                    format: "jpeg",
                    height: 132,
                    largePicUrl:
                      "https://image.xyzcdn.net/Fi4IKhYGlss9IpqkYg9-9Xa5eYV9@large",
                    middlePicUrl:
                      "https://image.xyzcdn.net/Fi4IKhYGlss9IpqkYg9-9Xa5eYV9@middle",
                    picUrl:
                      "https://image.xyzcdn.net/Fi4IKhYGlss9IpqkYg9-9Xa5eYV9",
                    smallPicUrl:
                      "https://image.xyzcdn.net/Fi4IKhYGlss9IpqkYg9-9Xa5eYV9@small",
                    thumbnailUrl:
                      "https://image.xyzcdn.net/Fi4IKhYGlss9IpqkYg9-9Xa5eYV9@thumbnail",
                    width: 132,
                  },
                },
                bio: "投资银行｜科技｜创业",
                gender: "MALE",
                ipLoc: "北京",
                isBlockedByViewer: false,
                isCancelled: false,
                isNicknameSet: true,
                nickname: "_刀锋",
                readTrackInfo: {},
                relation: "STRANGE",
                type: "USER",
                uid: "61dbfbf32cbd7c0176dca312",
              },
              // ...
            ],
            readTrackInfo: {},
            status: "NORMAL",
            subscriptionCount: 8979,
            subscriptionPush: false,
            subscriptionPushPriority: "HIGH",
            subscriptionStar: false,
            subscriptionStatus: "OFF",
            syncMode: "RSS",
            title: "Rum Cake",
            topicLabels: [],
            type: "PODCAST",
          },
          pubDate: "2024-05-29T14:07:51.000Z",
          readTrackInfo: {},
          shownotes: '<p style="color:#333333;...',
          sponsors: [],
          status: "NORMAL",
          title: "EP50：糖尿病的ABCs",
          transcript: {
            mediaId:
              "61dc011a1f7c9738c060c944/lk2GiABQdVlMYPY2tTJqYd-TEP0E.mp4a",
          },
          type: "EPISODE",
          wechatShare: {
            style: "LINK_CARD",
          },
        },
        podcast: {
          author: "_朗姆酒蛋糕",
          brief: "浅显易懂地解释经济和商业知识",
          color: {
            dark: "#CE851C",
            light: "#CE851C",
            original: "#CC8629",
          },
          contacts: [
            {
              name: "S_Jimmy",
              note: "加听友群",
              type: "wechat",
            },
            {
              name: "kidpodcast824@gmail.com",
              note: "听众反馈",
              type: "email",
            },
          ],
          description:
            "我们希望用声音，让每一个人都能获取属于自己的深度商业知识，...",
          episodeCount: 50,
          hasPopularEpisodes: true,
          image: {
            largePicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy83OTllLWF1ZGlvZnJlZWhpZ2hxcHMvQkEvMkEvR0t3UklKSUYtMnl4QUFGYlp3RWtrQ1NsLmpwZWc=.jpeg@large",
            middlePicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy83OTllLWF1ZGlvZnJlZWhpZ2hxcHMvQkEvMkEvR0t3UklKSUYtMnl4QUFGYlp3RWtrQ1NsLmpwZWc=.jpeg@middle",
            picUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy83OTllLWF1ZGlvZnJlZWhpZ2hxcHMvQkEvMkEvR0t3UklKSUYtMnl4QUFGYlp3RWtrQ1NsLmpwZWc=.jpeg",
            smallPicUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy83OTllLWF1ZGlvZnJlZWhpZ2hxcHMvQkEvMkEvR0t3UklKSUYtMnl4QUFGYlp3RWtrQ1NsLmpwZWc=.jpeg@small",
            thumbnailUrl:
              "https://bts-image.xyzcdn.net/aHR0cHM6Ly9mZGZzLnhtY2RuLmNvbS9zdG9yYWdlcy83OTllLWF1ZGlvZnJlZWhpZ2hxcHMvQkEvMkEvR0t3UklKSUYtMnl4QUFGYlp3RWtrQ1NsLmpwZWc=.jpeg@thumbnail",
          },
          isCustomized: false,
          latestEpisodePubDate: "2024-05-29T14:07:51.000Z",
          payEpisodeCount: 0,
          payType: "FREE",
          permissions: [
            {
              name: "SHARE",
              status: "PERMITTED",
            },
          ],
          pid: "61dc011a1f7c9738c060c944",
          podcasters: [
            {
              avatar: {
                picture: {
                  format: "jpeg",
                  height: 132,
                  largePicUrl:
                    "https://image.xyzcdn.net/Fi4IKhYGlss9IpqkYg9-9Xa5eYV9@large",
                  middlePicUrl:
                    "https://image.xyzcdn.net/Fi4IKhYGlss9IpqkYg9-9Xa5eYV9@middle",
                  picUrl:
                    "https://image.xyzcdn.net/Fi4IKhYGlss9IpqkYg9-9Xa5eYV9",
                  smallPicUrl:
                    "https://image.xyzcdn.net/Fi4IKhYGlss9IpqkYg9-9Xa5eYV9@small",
                  thumbnailUrl:
                    "https://image.xyzcdn.net/Fi4IKhYGlss9IpqkYg9-9Xa5eYV9@thumbnail",
                  width: 132,
                },
              },
              bio: "投资银行｜科技｜创业",
              gender: "MALE",
              ipLoc: "北京",
              isBlockedByViewer: false,
              isCancelled: false,
              isNicknameSet: true,
              nickname: "_刀锋",
              readTrackInfo: {},
              relation: "STRANGE",
              type: "USER",
              uid: "61dbfbf32cbd7c0176dca312",
            },
            // ...
          ],
          readTrackInfo: {},
          status: "NORMAL",
          subscriptionCount: 8979,
          subscriptionPush: false,
          subscriptionPushPriority: "HIGH",
          subscriptionStar: false,
          subscriptionStatus: "OFF",
          syncMode: "RSS",
          title: "Rum Cake",
          topicLabels: [],
          type: "PODCAST",
        },
      },
      //  ...
    ],
    loadMoreKey: 10,
  },
  msg: "OK",
}
```