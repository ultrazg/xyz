### 查询节目详情

根据 eid 查询单集内容详情

#### 请求地址

> /podcast_detail

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
| ...               | ...    | ...                                                |


#### 示例

> 地址：https://www.example.com/podcast_detail

参数

```javascript
{
  "pid": "61791d921989541784257779"
}
```

响应

```javascript
{
    "code": 200,
    "data": {
        "data": {
            "author": "不开玩笑小助手",
            "brief": "由猫头鹰喜剧出品",
            "color": {
                "dark": "#0B0B0B",
                "light": "#ACACAC",
                "original": "#1A1A1A"
            },
            "contacts": [
                {
                    "name": "猫头鹰喜剧",
                    "note": "订阅公众号",
                    "type": "wechatOfficialAccounts"
                },
                {
                    "name": "不开玩笑JokesAside",
                    "type": "weibo",
                    "url": "https://weibo.com/n/不开玩笑JokesAside"
                },
                {
                    "name": "JokesAside2024",
                    "note": "加听友群",
                    "type": "wechat"
                },
                {
                    "name": "bkwx@jieyanwenhua.com",
                    "note": "商务合作",
                    "type": "email"
                }
            ],
            "description": "一个时而开玩笑，时而不开玩笑的播客节目，由猫头鹰喜剧出品。\n\n公众号“猫头鹰喜剧”回复“听友群”，小助手会把你拉进群聊~",
            "episodeCount": 136,
            "hasPopularEpisodes": true,
            "image": {
                "largePicUrl": "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png@large",
                "middlePicUrl": "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png@middle",
                "picUrl": "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png",
                "smallPicUrl": "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png@small",
                "thumbnailUrl": "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZvYjZKc3BveHl4MEg4VzJzZEtlUTdFNHU5M3oucG5n.png@thumbnail"
            },
            "isCustomized": false,
            "latestEpisodePubDate": "2024-05-03T16:00:00.000Z",
            "payEpisodeCount": 0,
            "payType": "FREE",
            "permissions": [
                {
                    "name": "SHARE",
                    "status": "PERMITTED"
                }
            ],
            "pid": "61791d921989541784257779",
            "podcasters": [
                {
                    "avatar": {
                        "picture": {
                            "format": "jpeg",
                            "height": 3000,
                            "largePicUrl": "https://image.xyzcdn.net/FnQ-E7VcqLbzqplvdVPGrQRGHmxC.jpg@large",
                            "middlePicUrl": "https://image.xyzcdn.net/FnQ-E7VcqLbzqplvdVPGrQRGHmxC.jpg@middle",
                            "picUrl": "https://image.xyzcdn.net/FnQ-E7VcqLbzqplvdVPGrQRGHmxC.jpg",
                            "smallPicUrl": "https://image.xyzcdn.net/FnQ-E7VcqLbzqplvdVPGrQRGHmxC.jpg@small",
                            "thumbnailUrl": "https://image.xyzcdn.net/FnQ-E7VcqLbzqplvdVPGrQRGHmxC.jpg@thumbnail",
                            "width": 3000
                        }
                    },
                    "bio": "不开玩笑 Jokes Aside 由猫头鹰喜剧出品。",
                    "ipLoc": "上海",
                    "isBlockedByViewer": false,
                    "isCancelled": false,
                    "isNicknameSet": false,
                    "nickname": "不开玩笑小助手",
                    "readTrackInfo": {},
                    "relation": "STRANGE",
                    "type": "USER",
                    "uid": "61518a53e0f5e723bb520729"
                },
                {
                    "avatar": {
                        "picture": {
                            "format": "jpeg",
                            "height": 132,
                            "largePicUrl": "https://image.xyzcdn.net/FuzCz6vq3TRfppxYKvZ5iKi_mY9Y@large",
                            "middlePicUrl": "https://image.xyzcdn.net/FuzCz6vq3TRfppxYKvZ5iKi_mY9Y@middle",
                            "picUrl": "https://image.xyzcdn.net/FuzCz6vq3TRfppxYKvZ5iKi_mY9Y",
                            "smallPicUrl": "https://image.xyzcdn.net/FuzCz6vq3TRfppxYKvZ5iKi_mY9Y@small",
                            "thumbnailUrl": "https://image.xyzcdn.net/FuzCz6vq3TRfppxYKvZ5iKi_mY9Y@thumbnail",
                            "width": 132
                        }
                    },
                    "bio": "都在酒里了",
                    "gender": "MALE",
                    "ipLoc": "上海",
                    "isBlockedByViewer": false,
                    "isCancelled": false,
                    "isNicknameSet": false,
                    "nickname": "史炎nacl",
                    "readTrackInfo": {},
                    "relation": "STRANGE",
                    "type": "USER",
                    "uid": "6178e058e0f5e723bbed99ff"
                },
                {
                    "avatar": {
                        "picture": {
                            "format": "jpeg",
                            "height": 1500,
                            "largePicUrl": "https://image.xyzcdn.net/FpOL3DvumRmuC102c1cJQv8T5oC1.jpg@large",
                            "middlePicUrl": "https://image.xyzcdn.net/FpOL3DvumRmuC102c1cJQv8T5oC1.jpg@middle",
                            "picUrl": "https://image.xyzcdn.net/FpOL3DvumRmuC102c1cJQv8T5oC1.jpg",
                            "smallPicUrl": "https://image.xyzcdn.net/FpOL3DvumRmuC102c1cJQv8T5oC1.jpg@small",
                            "thumbnailUrl": "https://image.xyzcdn.net/FpOL3DvumRmuC102c1cJQv8T5oC1.jpg@thumbnail",
                            "width": 1500
                        }
                    },
                    "gender": "MALE",
                    "ipLoc": "陕西",
                    "isBlockedByViewer": false,
                    "isCancelled": false,
                    "isNicknameSet": true,
                    "nickname": "韦若琛",
                    "readTrackInfo": {},
                    "relation": "STRANGE",
                    "type": "USER",
                    "uid": "620ba7e32cbd7c01760224ab"
                },
                {
                    "avatar": {
                        "picture": {
                            "format": "jpeg",
                            "height": 2784,
                            "largePicUrl": "https://image.xyzcdn.net/FuHVidKWCk8NUyemy8UsmtSfHYYa.jpg@large",
                            "middlePicUrl": "https://image.xyzcdn.net/FuHVidKWCk8NUyemy8UsmtSfHYYa.jpg@middle",
                            "picUrl": "https://image.xyzcdn.net/FuHVidKWCk8NUyemy8UsmtSfHYYa.jpg",
                            "smallPicUrl": "https://image.xyzcdn.net/FuHVidKWCk8NUyemy8UsmtSfHYYa.jpg@small",
                            "thumbnailUrl": "https://image.xyzcdn.net/FuHVidKWCk8NUyemy8UsmtSfHYYa.jpg@thumbnail",
                            "width": 2784
                        }
                    },
                    "bio": "就这样儿吧",
                    "gender": "MALE",
                    "ipLoc": "北京",
                    "isBlockedByViewer": false,
                    "isCancelled": false,
                    "isNicknameSet": true,
                    "nickname": "大雄话不多",
                    "readTrackInfo": {},
                    "relation": "STRANGE",
                    "type": "USER",
                    "uid": "620bae352cbd7c01769b9bdd"
                },
                {
                    "avatar": {
                        "picture": {
                            "format": "jpeg",
                            "height": 132,
                            "largePicUrl": "https://image.xyzcdn.net/Fs-i5LktLDU0PHj5iXX19QVHODfb@large",
                            "middlePicUrl": "https://image.xyzcdn.net/Fs-i5LktLDU0PHj5iXX19QVHODfb@middle",
                            "picUrl": "https://image.xyzcdn.net/Fs-i5LktLDU0PHj5iXX19QVHODfb",
                            "smallPicUrl": "https://image.xyzcdn.net/Fs-i5LktLDU0PHj5iXX19QVHODfb@small",
                            "thumbnailUrl": "https://image.xyzcdn.net/Fs-i5LktLDU0PHj5iXX19QVHODfb@thumbnail",
                            "width": 132
                        }
                    },
                    "bio": "",
                    "gender": "MALE",
                    "ipLoc": "北京",
                    "isBlockedByViewer": false,
                    "isCancelled": false,
                    "isNicknameSet": true,
                    "nickname": "Notorious天一",
                    "readTrackInfo": {},
                    "relation": "STRANGE",
                    "type": "USER",
                    "uid": "62500ab0edce67104a40471e"
                },
                {
                    "avatar": {
                        "picture": {
                            "format": "jpeg",
                            "height": 1280,
                            "largePicUrl": "https://image.xyzcdn.net/Frfq79dA8qpT3ZXKanuVt7h4IduS.jpg@large",
                            "middlePicUrl": "https://image.xyzcdn.net/Frfq79dA8qpT3ZXKanuVt7h4IduS.jpg@middle",
                            "picUrl": "https://image.xyzcdn.net/Frfq79dA8qpT3ZXKanuVt7h4IduS.jpg",
                            "smallPicUrl": "https://image.xyzcdn.net/Frfq79dA8qpT3ZXKanuVt7h4IduS.jpg@small",
                            "thumbnailUrl": "https://image.xyzcdn.net/Frfq79dA8qpT3ZXKanuVt7h4IduS.jpg@thumbnail",
                            "width": 1280
                        }
                    },
                    "bio": "一男的。喜剧风格不太常规，购票之前建议先看公众号【梁彦增脱口秀】",
                    "gender": "MALE",
                    "ipLoc": "广东",
                    "isBlockedByViewer": false,
                    "isCancelled": false,
                    "isNicknameSet": true,
                    "nickname": "梁彦增",
                    "readTrackInfo": {},
                    "relation": "STRANGE",
                    "type": "USER",
                    "uid": "63d54919edce67104a0f90e8"
                },
                {
                    "avatar": {
                        "picture": {
                            "format": "jpeg",
                            "height": 132,
                            "largePicUrl": "https://image.xyzcdn.net/Fqqa8nPYP41j9xfojI7Y1eXDVwq3@large",
                            "middlePicUrl": "https://image.xyzcdn.net/Fqqa8nPYP41j9xfojI7Y1eXDVwq3@middle",
                            "picUrl": "https://image.xyzcdn.net/Fqqa8nPYP41j9xfojI7Y1eXDVwq3",
                            "smallPicUrl": "https://image.xyzcdn.net/Fqqa8nPYP41j9xfojI7Y1eXDVwq3@small",
                            "thumbnailUrl": "https://image.xyzcdn.net/Fqqa8nPYP41j9xfojI7Y1eXDVwq3@thumbnail",
                            "width": 132
                        }
                    },
                    "bio": "帅哥。",
                    "gender": "MALE",
                    "ipLoc": "上海",
                    "isBlockedByViewer": false,
                    "isCancelled": false,
                    "isNicknameSet": true,
                    "nickname": "刘仁铖主播版",
                    "readTrackInfo": {},
                    "relation": "STRANGE",
                    "type": "USER",
                    "uid": "6189ea92e0f5e723bbe4c1aa"
                }
            ],
            "readTrackInfo": {},
            "status": "NORMAL",
            "subscriptionCount": 149772,
            "subscriptionPush": false,
            "subscriptionPushPriority": "HIGH",
            "subscriptionStar": false,
            "subscriptionStatus": "OFF",
            "syncMode": "SELF_HOSTING",
            "title": "不开玩笑 Jokes Aside",
            "topicLabels": [],
            "type": "PODCAST"
        }
    },
    "msg": "OK"
}
```