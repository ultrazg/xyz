### 更新用户偏好设置

更新用户偏好设置

#### 请求地址

> /user_preference_update

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数 | 必填 | 类型    | 说明                                         |
| ---- | ---- | ------- | -------------------------------------------- |
| type | true | string  | 设置项。由 **/user_preference_get** 接口获取 |
| flag | true | boolean | 是否启用                                     |



#### 返回字段

| 返回字段                       | 类型    | 说明                       |
| ------------------------------ | ------- | -------------------------- |
| isRecentPlayedHidden           | boolean | 不让他人看到的收听记录     |
| isListenMileageHiddenInComment | boolean | 不在评论区展示收听时长标签 |
| isStickerLibraryHidden         | boolean | 不让他人看到我的贴纸库     |
| isStickerBoardHidden           | boolean | 不让他人看到我的贴纸装扮   |
| rejectHotPush                  | boolean | 接收热门内容推送           |
| rejectRecommendation           | boolean | 个性化推荐                 |




#### 示例

> 地址：https://www.example.com/user_preference_update

参数

```javascript
{
  "type": "isRecentPlayedHidden",
  "flag": true
}
```

响应

``` javascript
{
  code: 200,
  data: {
    data: {
      enableNextPlayLiveActivity: false,
      enableTimestampCommentLiveActivity: false,
      isListenMileageHiddenInComment: false,
      isRecentPlayedHidden: true,
      isStickerBoardHidden: false,
      isStickerLibraryHidden: false,
      rejectHotPush: false,
      rejectRecommendation: true,
    },
  },
  msg: "OK",
}
```