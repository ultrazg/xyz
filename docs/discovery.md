### 首页榜单、精选节目、推荐等

首页推荐、精选、榜单等模块

#### 请求地址

> /discovery

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数        | 必填  | 类型   | 说明                                                         |
| :---------- | :---- | :----- | ------------------------------------------------------------ |
| loadMoreKey | false | string | 类型。**不传**则查询 APP 首页标题为「大家都在听」（注：可能为其他标题，如「XX 的听众也爱听」、「宝藏新节目」等）和「编辑精选」两个板块的内容；参数为 **mediumDiscoveryPictorial** 则查询「最热榜」、「锋芒榜」和「新星榜」的内容；参数为 **discoveryTopic** 则查询「为你精选的节目」顺序往后（包括）4 个板块（到「播客寻宝」结束）的内容；参数为 **pick** 则查询「TA 们的喜欢」和「TA 们开始创作新播客」两个板块的内容 |

#### 返回字段

参考「查询单集详情」接口与「查询节目详情」接口，或返回的字段


#### 示例

> 地址：https://www.example.com/discovery

参数

```javascript
{
  "loadMoreKey": "discoveryTopic"
}
```

响应

```javascript
// 示例为查询 discoveryTopic 时的内容
{
  code: 200,
  data: {
    data: [
      {
        type: "DISCOVERY_COLLECTION",
        data: [
          {
            title: "为你精选的节目",
            moduleType: "HPR_MULTI_RECALL",
            targetType: "PODCAST",
            target: [
              {
                podcast: {
                  type: "PODCAST",
                  pid: "605ade4f9db0b72c34e96dee",
                  title: "木有娱丸",
                  author: "木有娱丸",
                  brief: "娱乐资本论出品，聊文娱圈的一切",
                  description: "“木有娱丸”是由上海...",
                  subscriptionCount: 46133,
                  image: {
                    picUrl:
                      "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZtYVNXTjBIM1M3OTI3cW1sai0xZi1IbUZJUWsuanBn.jpg",
                    largePicUrl:
                      "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZtYVNXTjBIM1M3OTI3cW1sai0xZi1IbUZJUWsuanBn.jpg@large",
                    middlePicUrl:
                      "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZtYVNXTjBIM1M3OTI3cW1sai0xZi1IbUZJUWsuanBn.jpg@middle",
                    smallPicUrl:
                      "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZtYVNXTjBIM1M3OTI3cW1sai0xZi1IbUZJUWsuanBn.jpg@small",
                    thumbnailUrl:
                      "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZtYVNXTjBIM1M3OTI3cW1sai0xZi1IbUZJUWsuanBn.jpg@thumbnail",
                  },
                  color: {
                    original: "#D043C0",
                    light: "#CA4EBB",
                    dark: "#CA4EBB",
                  },
                  topicLabels: [],
                  syncMode: "SELF_HOSTING",
                  episodeCount: 105,
                  latestEpisodePubDate: "2024-05-17T06:01:13.082Z",
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
                      uid: "605ada88e0f5e723bb291a8d",
                      avatar: {
                        picture: {
                          picUrl:
                            "https://image.xyzcdn.net/FhFK3vhH_8YBXf_zFTGgz97hc7Dl.jpg",
                          largePicUrl:
                            "https://image.xyzcdn.net/FhFK3vhH_8YBXf_zFTGgz97hc7Dl.jpg@large",
                          middlePicUrl:
                            "https://image.xyzcdn.net/FhFK3vhH_8YBXf_zFTGgz97hc7Dl.jpg@middle",
                          smallPicUrl:
                            "https://image.xyzcdn.net/FhFK3vhH_8YBXf_zFTGgz97hc7Dl.jpg@small",
                          thumbnailUrl:
                            "https://image.xyzcdn.net/FhFK3vhH_8YBXf_zFTGgz97hc7Dl.jpg@thumbnail",
                          format: "jpeg",
                          width: 500,
                          height: 500,
                        },
                      },
                      nickname: "木有娱丸",
                      isNicknameSet: false,
                      isCancelled: false,
                      readTrackInfo: {},
                      ipLoc: "北京",
                      relation: "STRANGE",
                      isBlockedByViewer: false,
                    },
                  ],
                  readTrackInfo: {
                    recallPolicy: "HPR_ALS",
                    additionalInfo:
                      '{"recModuleType":"HPR_MULTI_RECALL","recModuleLoc":"TOP","recallScore":"0.1141","recDescType":"guess_like"}',
                    recSize: "6",
                  },
                  hasPopularEpisodes: true,
                  contacts: [
                    {
                      type: "wechatOfficialAccounts",
                      name: "娱乐资本论",
                      note: "订阅公众号",
                    },
                    {
                      type: "weibo",
                      name: "娱乐资本论",
                      url: "https://weibo.com/n/娱乐资本论",
                    },
                    {
                      type: "wechat",
                      name: "happywanwan0",
                      note: "加听友群",
                    },
                    {
                      type: "custom",
                      name: "微信客服ylzbl_kexiaofu",
                    },
                  ],
                  isCustomized: false,
                },
                recommendation: "猜你喜欢",
                relatedUsers: [],
                hasNegativeFeedback: false,
              },
              //  ...
            ],
            displayType: "PODCAST_DEFAULT",
            lightBlueEndIndex: 2,
            feedback: {
              canFeedback: true,
              feedbackSource: "DISCOVERY_COLLECTION_V3",
            },
            readTrackInfo: {
              recModuleType: "HPR_MULTI_RECALL",
              recModuleLoc: "TOP",
              dataDate: "20240520",
              recDate: "20240520",
              recSize: "6",
            },
          },
          {
            title: "关于我爱你",
            description: "探索情感世界，共同成长",
            moduleType: "EDITOR_PICKED",
            targetType: "EPISODE",
            target: [
              {
                episode: {
                  type: "EPISODE",
                  eid: "662ba6a784f78c922c4bc27b",
                  pid: "60e2d4c838b5b0e1a54180c4",
                  title:
                    "E76 被爱创飞小半生，归来不做女菩萨：爱情教会了我们点啥？",
                  shownotes:
                    "<p>重磅！</p>\n<p>上一次在「北京装人」这期节目里大获好评的两位嘉宾：...",
                  description:
                    "重磅！\n上一次在「北京装人」这期节目里大获好评的两位嘉宾：...",
                  enclosure: {
                    url: "https://media.xyzcdn.net/lrAX-auIC3teoFhfLWm2uHzG3T-Z.m4a",
                  },
                  isPrivateMedia: false,
                  mediaKey: "lrAX-auIC3teoFhfLWm2uHzG3T-Z.m4a",
                  media: {
                    id: "lrAX-auIC3teoFhfLWm2uHzG3T-Z.m4a",
                    size: 93371553,
                    mimeType: "audio/mp4",
                    source: {
                      mode: "PUBLIC",
                      url: "https://media.xyzcdn.net/lrAX-auIC3teoFhfLWm2uHzG3T-Z.m4a",
                    },
                  },
                  clapCount: 64,
                  commentCount: 166,
                  playCount: 16903,
                  favoriteCount: 110,
                  pubDate: "2024-04-26T13:05:42.627Z",
                  status: "NORMAL",
                  duration: 3860,
                  podcast: {
                    type: "PODCAST",
                    pid: "60e2d4c838b5b0e1a54180c4",
                    title: "离心力比多",
                    author: "拾壹&壁仔",
                    brief: "一对两岸小夫妻的闲聊",
                    description:
                      "这是一档打破两性二元对立，消除立场的遮蔽，真正建立了解、认识彼此的节目。\n\n由两岸小夫妻拾壹和壁仔共同主持，我们不带预设的去触摸真实的世界和人心。\n\n\n📮媒体、商务合作请联系微信：wbyswaiwai，也可来信至：relic_ice@icloud.com\n\n\n微博：快乐壁壁、拾壹-歪波音室\n邮箱：relic_ice@icloud.com",
                    subscriptionCount: 66706,
                    image: {
                      picUrl:
                        "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZqSlpHSTRzcVgxMTU4U01ZZHAtTFlRU0FfQVMucG5n.png",
                      largePicUrl:
                        "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZqSlpHSTRzcVgxMTU4U01ZZHAtTFlRU0FfQVMucG5n.png@large",
                      middlePicUrl:
                        "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZqSlpHSTRzcVgxMTU4U01ZZHAtTFlRU0FfQVMucG5n.png@middle",
                      smallPicUrl:
                        "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZqSlpHSTRzcVgxMTU4U01ZZHAtTFlRU0FfQVMucG5n.png@small",
                      thumbnailUrl:
                        "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZqSlpHSTRzcVgxMTU4U01ZZHAtTFlRU0FfQVMucG5n.png@thumbnail",
                    },
                    color: {
                      original: "#0D0D0D",
                      light: "#ACACAC",
                      dark: "#0B0B0B",
                    },
                    topicLabels: [],
                    syncMode: "SELF_HOSTING",
                    episodeCount: 77,
                    latestEpisodePubDate: "2024-05-15T02:25:25.507Z",
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
                        uid: "5e7c8a5aa2f2e503cec3b64b",
                        avatar: {
                          picture: {
                            picUrl:
                              "https://image.xyzcdn.net/FtqOUFIeqzEZXMB6yga6kwfSeoz3.jpg",
                            largePicUrl:
                              "https://image.xyzcdn.net/FtqOUFIeqzEZXMB6yga6kwfSeoz3.jpg@large",
                            middlePicUrl:
                              "https://image.xyzcdn.net/FtqOUFIeqzEZXMB6yga6kwfSeoz3.jpg@middle",
                            smallPicUrl:
                              "https://image.xyzcdn.net/FtqOUFIeqzEZXMB6yga6kwfSeoz3.jpg@small",
                            thumbnailUrl:
                              "https://image.xyzcdn.net/FtqOUFIeqzEZXMB6yga6kwfSeoz3.jpg@thumbnail",
                            format: "jpeg",
                            width: 3000,
                            height: 3000,
                          },
                        },
                        nickname: "拾壹",
                        isNicknameSet: false,
                        bio: " ",
                        gender: "MALE",
                        isCancelled: false,
                        readTrackInfo: {},
                        ipLoc: "天津",
                        relation: "STRANGE",
                        isBlockedByViewer: false,
                      },
                      {
                        type: "USER",
                        uid: "5ea54a7968a1682f67cf0aa3",
                        avatar: {
                          picture: {
                            picUrl:
                              "https://image.xyzcdn.net/Frt_FFHsceGzrA3d2jM5I8ayMvEL.jpg",
                            largePicUrl:
                              "https://image.xyzcdn.net/Frt_FFHsceGzrA3d2jM5I8ayMvEL.jpg@large",
                            middlePicUrl:
                              "https://image.xyzcdn.net/Frt_FFHsceGzrA3d2jM5I8ayMvEL.jpg@middle",
                            smallPicUrl:
                              "https://image.xyzcdn.net/Frt_FFHsceGzrA3d2jM5I8ayMvEL.jpg@small",
                            thumbnailUrl:
                              "https://image.xyzcdn.net/Frt_FFHsceGzrA3d2jM5I8ayMvEL.jpg@thumbnail",
                            format: "jpeg",
                            width: 1902,
                            height: 1902,
                          },
                        },
                        nickname: "BiZai",
                        isNicknameSet: false,
                        bio: "幼稚又变态丨微博@快乐壁壁",
                        gender: "FEMALE",
                        isCancelled: false,
                        readTrackInfo: {},
                        ipLoc: "北京",
                        relation: "STRANGE",
                        isBlockedByViewer: false,
                      },
                    ],
                    readTrackInfo: {},
                    hasPopularEpisodes: true,
                    contacts: [
                      {
                        type: "wechat",
                        name: "wbyswaiwai",
                        note: "添加微信",
                      },
                      {
                        type: "email",
                        name: "relic_ice@icloud.com",
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
                    recallPolicy: "HORIREC_EDITOR_PICKED",
                    additionalInfo:
                      '{"recModuleType":"EDITOR_PICKED","recModuleLoc":"TOP","collectionId":"6649feabe5fa6fd95f551004","recModuleName":"关于我爱你"}',
                    backendRecDescType: "recommend",
                  },
                  labels: [],
                  sponsors: [],
                  isCustomized: false,
                  ipLoc: "山西",
                  transcript: {
                    mediaId: "lrAX-auIC3teoFhfLWm2uHzG3T-Z.m4a",
                  },
                },
                recommendation: "亲密关系如何重建内心世界",
                relatedUsers: [],
                hasNegativeFeedback: false,
              },
              //  ...
            ],
            displayType: "EPISODE_DEFAULT",
            feedback: {
              canFeedback: false,
              feedbackSource: "DISCOVERY_COLLECTION_V3",
            },
            readTrackInfo: {
              recModuleType: "EDITOR_PICKED",
              recModuleLoc: "TOP",
              dataDate: "20240520",
              recDate: "20240520",
              recSize: "12",
            },
            collectionId: "6649feabe5fa6fd95f551004",
          },
          {
            title: "趣味闲谈类的精彩单集",
            moduleType: "HER_LIKE_CATEGORY",
            targetType: "EPISODE",
            target: [
              {
                episode: {
                  type: "EPISODE",
                  eid: "66322c094b7d3b5d3b97c33f",
                  pid: "650d65cd5c88d241269133a7",
                  title: "35、那一刻，我真的很想分手：恋爱中的失望时刻",
                  shownotes:
                    '<figure><img src="https://image.xyzcdn.net/FmiMaJdrl-y1pim6kjGECkBIm4Cm.png"/></figure>...',
                  description: "每个人都希望拥有甜甜的恋爱...",
                  image: {
                    picUrl:
                      "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z2YUQwbGVHSzBNWXVpQ19uWDBySTNFYTNfVGoucG5n.png",
                    largePicUrl:
                      "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z2YUQwbGVHSzBNWXVpQ19uWDBySTNFYTNfVGoucG5n.png@large",
                    middlePicUrl:
                      "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z2YUQwbGVHSzBNWXVpQ19uWDBySTNFYTNfVGoucG5n.png@middle",
                    smallPicUrl:
                      "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z2YUQwbGVHSzBNWXVpQ19uWDBySTNFYTNfVGoucG5n.png@small",
                    thumbnailUrl:
                      "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0Z2YUQwbGVHSzBNWXVpQ19uWDBySTNFYTNfVGoucG5n.png@thumbnail",
                  },
                  enclosure: {
                    url: "https://media.xyzcdn.net/liSG9VEHmhuQnmgx4F0_nN3H8rlI.m4a",
                  },
                  isPrivateMedia: false,
                  mediaKey: "liSG9VEHmhuQnmgx4F0_nN3H8rlI.m4a",
                  media: {
                    id: "liSG9VEHmhuQnmgx4F0_nN3H8rlI.m4a",
                    size: 104831217,
                    mimeType: "audio/mp4",
                    source: {
                      mode: "PUBLIC",
                      url: "https://media.xyzcdn.net/liSG9VEHmhuQnmgx4F0_nN3H8rlI.m4a",
                    },
                  },
                  clapCount: 55,
                  commentCount: 173,
                  playCount: 26568,
                  favoriteCount: 201,
                  pubDate: "2024-05-02T23:00:00.000Z",
                  status: "NORMAL",
                  duration: 3256,
                  podcast: {
                    type: "PODCAST",
                    pid: "650d65cd5c88d241269133a7",
                    title: "浪里个浪！",
                    author: "拾柒No_17",
                    brief: "三位芒果台员工的无话不谈，由凹凸宇宙出品",
                    description: "☆本节目由播客厂牌——凹凸宇宙出品...",
                    subscriptionCount: 13734,
                    image: {
                      picUrl:
                        "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZoRmY2eXlreGpWVk1XcS1ReFN3Ukl2d2QxRlEucG5n.png",
                      largePicUrl:
                        "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZoRmY2eXlreGpWVk1XcS1ReFN3Ukl2d2QxRlEucG5n.png@large",
                      middlePicUrl:
                        "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZoRmY2eXlreGpWVk1XcS1ReFN3Ukl2d2QxRlEucG5n.png@middle",
                      smallPicUrl:
                        "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZoRmY2eXlreGpWVk1XcS1ReFN3Ukl2d2QxRlEucG5n.png@small",
                      thumbnailUrl:
                        "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZoRmY2eXlreGpWVk1XcS1ReFN3Ukl2d2QxRlEucG5n.png@thumbnail",
                    },
                    color: {
                      original: "#de2198",
                      light: "#D23F95",
                      dark: "#D23F95",
                    },
                    hasTopic: true,
                    topicLabels: ["ALL"],
                    syncMode: "SELF_HOSTING",
                    episodeCount: 37,
                    latestEpisodePubDate: "2024-05-16T23:00:00.000Z",
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
                        uid: "61a2427be0f5e723bbaf7d6d",
                        avatar: {
                          picture: {
                            picUrl:
                              "https://image.xyzcdn.net/FjDm1OsSs7zR1XZ786xu64qmygeQ.jpg",
                            largePicUrl:
                              "https://image.xyzcdn.net/FjDm1OsSs7zR1XZ786xu64qmygeQ.jpg@large",
                            middlePicUrl:
                              "https://image.xyzcdn.net/FjDm1OsSs7zR1XZ786xu64qmygeQ.jpg@middle",
                            smallPicUrl:
                              "https://image.xyzcdn.net/FjDm1OsSs7zR1XZ786xu64qmygeQ.jpg@small",
                            thumbnailUrl:
                              "https://image.xyzcdn.net/FjDm1OsSs7zR1XZ786xu64qmygeQ.jpg@thumbnail",
                            format: "jpeg",
                            width: 1333,
                            height: 1333,
                          },
                        },
                        nickname: "拾柒No_17",
                        isNicknameSet: true,
                        bio: "芒果台11年记者经历，白羊座男生，一个普通人，你不讨厌就好。",
                        gender: "MALE",
                        isCancelled: false,
                        readTrackInfo: {},
                        ipLoc: "湖南",
                        relation: "STRANGE",
                        isBlockedByViewer: false,
                      },
                      {
                        type: "USER",
                        uid: "5fa69d3ee0f5e723bb0913ef",
                        avatar: {
                          picture: {
                            picUrl:
                              "https://image.xyzcdn.net/FlcChb2voYcy2tmXPZ4myVreRxNU.jpg",
                            largePicUrl:
                              "https://image.xyzcdn.net/FlcChb2voYcy2tmXPZ4myVreRxNU.jpg@large",
                            middlePicUrl:
                              "https://image.xyzcdn.net/FlcChb2voYcy2tmXPZ4myVreRxNU.jpg@middle",
                            smallPicUrl:
                              "https://image.xyzcdn.net/FlcChb2voYcy2tmXPZ4myVreRxNU.jpg@small",
                            thumbnailUrl:
                              "https://image.xyzcdn.net/FlcChb2voYcy2tmXPZ4myVreRxNU.jpg@thumbnail",
                            format: "jpeg",
                            width: 3000,
                            height: 3000,
                          },
                        },
                        nickname: "是瑞秋呀",
                        isNicknameSet: true,
                        bio: "声音工作者｜长沙｜斜杠30+姐姐",
                        gender: "FEMALE",
                        isCancelled: false,
                        readTrackInfo: {},
                        ipLoc: "湖南",
                        relation: "STRANGE",
                        isBlockedByViewer: false,
                      },
                      {
                        type: "USER",
                        uid: "6109de3ee0f5e723bbd85538",
                        avatar: {
                          picture: {
                            picUrl:
                              "https://image.xyzcdn.net/Fgo_hO7TH2jBAgrAwHVC_a-So_3j.jpg",
                            largePicUrl:
                              "https://image.xyzcdn.net/Fgo_hO7TH2jBAgrAwHVC_a-So_3j.jpg@large",
                            middlePicUrl:
                              "https://image.xyzcdn.net/Fgo_hO7TH2jBAgrAwHVC_a-So_3j.jpg@middle",
                            smallPicUrl:
                              "https://image.xyzcdn.net/Fgo_hO7TH2jBAgrAwHVC_a-So_3j.jpg@small",
                            thumbnailUrl:
                              "https://image.xyzcdn.net/Fgo_hO7TH2jBAgrAwHVC_a-So_3j.jpg@thumbnail",
                            format: "jpeg",
                            width: 1500,
                            height: 1500,
                          },
                        },
                        nickname: "添樂*小添",
                        isNicknameSet: true,
                        bio: "媒体小哥，声音工作者，《浪里个浪》《瓜儿与少年》《装修内幕》主播",
                        gender: "MALE",
                        isCancelled: false,
                        readTrackInfo: {},
                        ipLoc: "湖南",
                        relation: "STRANGE",
                        isBlockedByViewer: false,
                      },
                      {
                        type: "USER",
                        uid: "66220126edce67104a690eac",
                        avatar: {
                          picture: {
                            picUrl:
                              "https://image.xyzcdn.net/Fmspv-nhcXc1q7_KatBafr5P0DA1.jpg",
                            largePicUrl:
                              "https://image.xyzcdn.net/Fmspv-nhcXc1q7_KatBafr5P0DA1.jpg@large",
                            middlePicUrl:
                              "https://image.xyzcdn.net/Fmspv-nhcXc1q7_KatBafr5P0DA1.jpg@middle",
                            smallPicUrl:
                              "https://image.xyzcdn.net/Fmspv-nhcXc1q7_KatBafr5P0DA1.jpg@small",
                            thumbnailUrl:
                              "https://image.xyzcdn.net/Fmspv-nhcXc1q7_KatBafr5P0DA1.jpg@thumbnail",
                            format: "jpeg",
                            width: 3000,
                            height: 3000,
                          },
                        },
                        nickname: "凹凸宇宙",
                        isNicknameSet: true,
                        bio: "播客厂牌：凹凸宇宙",
                        gender: "FEMALE",
                        isCancelled: false,
                        readTrackInfo: {},
                        ipLoc: "浙江",
                        relation: "STRANGE",
                        isBlockedByViewer: false,
                      },
                    ],
                    readTrackInfo: {},
                    hasPopularEpisodes: true,
                    contacts: [
                      {
                        type: "wechatOfficialAccounts",
                        name: "浪radio",
                        note: "订阅公众号",
                      },
                      {
                        type: "weibo",
                        name: "浪里个浪播客",
                        url: "https://weibo.com/n/浪里个浪播客",
                      },
                      {
                        type: "email",
                        name: "175845371@qq.com",
                        note: "商务合作",
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
                    recallPolicy: "HER_LIKE_CATEGORY_HOT",
                    additionalInfo:
                      '{"recModuleType":"HER_LIKE_CATEGORY","recModuleLoc":"TOP","recallScore":"10.0729","refId":"63c76a8924b1622727bd3222","ref":"趣味闲谈","recDescWord":"20240505","recDescType":"hot_list_date"}',
                    backendRecDescType: "recommend",
                  },
                  labels: [
                    {
                      name: "最受欢迎",
                      code: "POPULAR",
                    },
                  ],
                  sponsors: [],
                  isCustomized: false,
                  ipLoc: "湖南",
                  topicId: "65b1ef4635dd8780ed1366e5",
                  transcript: {
                    mediaId: "liSG9VEHmhuQnmgx4F0_nN3H8rlI.m4a",
                  },
                },
                recommendation: "进入 24 年 5 月最热榜",
                relatedUsers: [],
                hasNegativeFeedback: false,
              },
              // ...
            ],
            lightBlueEndIndex: 5,
            feedback: {
              canFeedback: true,
              feedbackSource: "DISCOVERY_COLLECTION_V3",
            },
            readTrackInfo: {
              recModuleType: "HER_LIKE_CATEGORY",
              recModuleLoc: "TOP",
              dataDate: "20240520",
              recDate: "20240520",
              recSize: "6",
            },
          },
        ],
      },
      {
        type: "NEW_POWER",
        data: {
          id: "6646fd768deded8bf8e42e7e",
          items: [
            {
              type: "PODCAST",
              pid: "646aedeafc773067cc7da018",
              title: "时常帮助",
              author: "马小咩咩咩咩咩",
              brief: "基层肿瘤科医生的癌症科普",
              description: "肿瘤科副主任医师，来聊聊癌症相关的那些事。...",
              subscriptionCount: 897,
              image: {
                picUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZwblgwcFlwN1BZNmduRUFhSVpBUWE1WGo1cVkuanBn.jpg",
                largePicUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZwblgwcFlwN1BZNmduRUFhSVpBUWE1WGo1cVkuanBn.jpg@large",
                middlePicUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZwblgwcFlwN1BZNmduRUFhSVpBUWE1WGo1cVkuanBn.jpg@middle",
                smallPicUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZwblgwcFlwN1BZNmduRUFhSVpBUWE1WGo1cVkuanBn.jpg@small",
                thumbnailUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZwblgwcFlwN1BZNmduRUFhSVpBUWE1WGo1cVkuanBn.jpg@thumbnail",
              },
              color: {
                original: "#DEF9E6",
                light: "#71DF9C",
                dark: "#71DF9C",
              },
              topicLabels: [],
              syncMode: "SELF_HOSTING",
              episodeCount: 5,
              latestEpisodePubDate: "2024-05-15T22:00:00.000Z",
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
                  uid: "6444ef35edce67104a7667a7",
                  avatar: {
                    picture: {
                      picUrl:
                        "https://image.xyzcdn.net/FrZ5JUl8J_O2uuXZN0afwoL1OFaw",
                      largePicUrl:
                        "https://image.xyzcdn.net/FrZ5JUl8J_O2uuXZN0afwoL1OFaw@large",
                      middlePicUrl:
                        "https://image.xyzcdn.net/FrZ5JUl8J_O2uuXZN0afwoL1OFaw@middle",
                      smallPicUrl:
                        "https://image.xyzcdn.net/FrZ5JUl8J_O2uuXZN0afwoL1OFaw@small",
                      thumbnailUrl:
                        "https://image.xyzcdn.net/FrZ5JUl8J_O2uuXZN0afwoL1OFaw@thumbnail",
                      format: "jpeg",
                      width: 132,
                      height: 132,
                    },
                  },
                  nickname: "马小咩咩咩咩咩",
                  isNicknameSet: true,
                  bio: "Keep fit and think big. ",
                  gender: "FEMALE",
                  isCancelled: false,
                  readTrackInfo: {},
                  ipLoc: "安徽",
                  relation: "STRANGE",
                  isBlockedByViewer: false,
                },
              ],
              readTrackInfo: {},
              hasPopularEpisodes: false,
              contacts: [
                {
                  type: "wechatOfficialAccounts",
                  name: "时常帮助",
                  note: "订阅公众号",
                },
                {
                  type: "weibo",
                  name: "玛蔻芮",
                  url: "https://weibo.com/n/玛蔻芮",
                },
              ],
              isCustomized: false,
              recommendation: "基层肿瘤科医生的癌症科普",
              detailedDescription: "《时常帮助》是一档癌症科普播客，...",
            },
            {
              type: "PODCAST",
              pid: "6506e74a283348dcad7b5c37",
              title: "食野之评",
              author: "小野酱",
              brief: "解构这个世界",
              description: "这是一档人文社科和商业观察的对谈节目。...",
              subscriptionCount: 2460,
              image: {
                picUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZwMUdfV1A5NFlVM21tbHpLdjU5S1A5U0F3VFYuanBn.jpg",
                largePicUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZwMUdfV1A5NFlVM21tbHpLdjU5S1A5U0F3VFYuanBn.jpg@large",
                middlePicUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZwMUdfV1A5NFlVM21tbHpLdjU5S1A5U0F3VFYuanBn.jpg@middle",
                smallPicUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZwMUdfV1A5NFlVM21tbHpLdjU5S1A5U0F3VFYuanBn.jpg@small",
                thumbnailUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZwMUdfV1A5NFlVM21tbHpLdjU5S1A5U0F3VFYuanBn.jpg@thumbnail",
              },
              color: {
                original: "#f7eeec",
                light: "#FD8D77",
                dark: "#FD8D77",
              },
              hasTopic: true,
              topicLabels: ["ALL"],
              syncMode: "SELF_HOSTING",
              episodeCount: 35,
              latestEpisodePubDate: "2024-05-13T16:00:00.000Z",
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
                  uid: "601b3b6be0f5e723bb4b7cfb",
                  avatar: {
                    picture: {
                      picUrl:
                        "https://image.xyzcdn.net/FlfxsItCEhutR8OPlDRu4eIgNuf_.jpg",
                      largePicUrl:
                        "https://image.xyzcdn.net/FlfxsItCEhutR8OPlDRu4eIgNuf_.jpg@large",
                      middlePicUrl:
                        "https://image.xyzcdn.net/FlfxsItCEhutR8OPlDRu4eIgNuf_.jpg@middle",
                      smallPicUrl:
                        "https://image.xyzcdn.net/FlfxsItCEhutR8OPlDRu4eIgNuf_.jpg@small",
                      thumbnailUrl:
                        "https://image.xyzcdn.net/FlfxsItCEhutR8OPlDRu4eIgNuf_.jpg@thumbnail",
                      format: "jpeg",
                      width: 3000,
                      height: 3000,
                    },
                  },
                  nickname: "小野酱",
                  isNicknameSet: true,
                  bio: "不具名的作家和投资人",
                  gender: "FEMALE",
                  isCancelled: false,
                  readTrackInfo: {},
                  ipLoc: "上海",
                  relation: "STRANGE",
                  isBlockedByViewer: false,
                },
              ],
              readTrackInfo: {},
              hasPopularEpisodes: true,
              contacts: [
                {
                  type: "wechatOfficialAccounts",
                  name: "作家小野酱",
                  note: "订阅公众号",
                },
                {
                  type: "weibo",
                  name: "作家小野酱",
                  url: "https://weibo.com/n/作家小野酱",
                },
                {
                  type: "jike",
                  name: "小野酱大漂亮",
                  url: "https://m.okjike.com/users/88a397c5-59a4-4616-968a-d4b9df3bb325",
                },
                {
                  type: "email",
                  name: "vera.huang@foxmail.com",
                  note: "商务合作",
                },
                {
                  type: "custom",
                  name: "小红书",
                  url: "https://www.xiaohongshu.com/user/profile/55cbde0ec2bdeb1dc94edd10?xhsshare=CopyLink&appuid=55cbde0ec2bdeb1dc94edd10&apptime=1705103347",
                },
              ],
              isCustomized: false,
              recommendation: "关注社科、商业、艺术和人心",
              detailedDescription:
                "《食野之评》的节目名称取自于《诗经·小雅》...",
            },
            {
              type: "PODCAST",
              pid: "660d0654c891c06e731420d9",
              title: "音乐进化论MusicEvolution",
              author: "勾勾同学_infj版",
              brief: "一份“立足老歌，放眼世界”的老派少女听歌指南",
              description: "“没有了音乐就退化耳朵”...",
              subscriptionCount: 387,
              image: {
                picUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZyNnpKamtpbHVkSEMzbklrSmVZRU50OEJMdUYuanBn.jpg",
                largePicUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZyNnpKamtpbHVkSEMzbklrSmVZRU50OEJMdUYuanBn.jpg@large",
                middlePicUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZyNnpKamtpbHVkSEMzbklrSmVZRU50OEJMdUYuanBn.jpg@middle",
                smallPicUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZyNnpKamtpbHVkSEMzbklrSmVZRU50OEJMdUYuanBn.jpg@small",
                thumbnailUrl:
                  "https://bts-image.xyzcdn.net/aHR0cHM6Ly9pbWFnZS54eXpjZG4ubmV0L0ZyNnpKamtpbHVkSEMzbklrSmVZRU50OEJMdUYuanBn.jpg@thumbnail",
              },
              color: {
                original: "#B9A087",
                light: "#DF9233",
                dark: "#DF9233",
              },
              topicLabels: [],
              syncMode: "SELF_HOSTING",
              episodeCount: 7,
              latestEpisodePubDate: "2024-05-16T23:30:00.000Z",
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
                  uid: "5ebc04a721ac8580416b8f26",
                  avatar: {
                    picture: {
                      picUrl:
                        "https://image.xyzcdn.net/FtJTh6kWU-YSg2xna1v6xCZEEHuh.jpg",
                      largePicUrl:
                        "https://image.xyzcdn.net/FtJTh6kWU-YSg2xna1v6xCZEEHuh.jpg@large",
                      middlePicUrl:
                        "https://image.xyzcdn.net/FtJTh6kWU-YSg2xna1v6xCZEEHuh.jpg@middle",
                      smallPicUrl:
                        "https://image.xyzcdn.net/FtJTh6kWU-YSg2xna1v6xCZEEHuh.jpg@small",
                      thumbnailUrl:
                        "https://image.xyzcdn.net/FtJTh6kWU-YSg2xna1v6xCZEEHuh.jpg@thumbnail",
                      format: "jpeg",
                      width: 2145,
                      height: 2145,
                    },
                  },
                  nickname: "勾勾同学_听歌版",
                  isNicknameSet: true,
                  bio: "“立足老歌，放眼世界”🌎",
                  gender: "FEMALE",
                  isCancelled: false,
                  readTrackInfo: {},
                  ipLoc: "北京",
                  relation: "STRANGE",
                  isBlockedByViewer: false,
                },
              ],
              readTrackInfo: {},
              hasPopularEpisodes: true,
              contacts: [
                {
                  type: "jike",
                  name: "勾勾同学_听歌版",
                  url: "https://m.okjike.com/users/9CBBF6C4-62F2-41B1-99DD-AB5279773F26",
                },
                {
                  type: "wechat",
                  name: "MeRadio2024",
                  note: "加听友群",
                },
              ],
              isCustomized: false,
              recommendation: "老派少女的老歌收听指南",
              detailedDescription:
                "这是一份”立足老歌，放眼世界”的老派少女听歌指南。...",
            },
          ],
          publishDate: "2024-05-20T00:00:00.000Z",
          serialNumber: "96",
        },
      },
    ],
    loadMoreKey: "pick",
  },
  msg: "OK",
}
```