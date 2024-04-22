### 短信登录

用于登录

#### 请求地址

> /login

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求参数

| 参数              | 必填  | 类型   | 说明          |
| :---------------- | :---- | :----- | ------------- |
| mobilePhoneNumber | true  | string | 手机号        |
| verifyCode | true  | string | 验证码        |
| areaCode          | false | string | 区号，默认+86 |

#### 返回字段

|返回字段|类型|说明                              |
|:-----   |:------|:-----------------------------   |
|user.type|string|用户角色|
|user.uid   |string    |用户 uid   |
|user.avatar   |object    |用户头像信息   |
|user.nickname   |string    |用户昵称   |
|user.wechatUserInfo   |object    |绑定的微信信息   |
|user.jikeUserInfo   |object    |绑定的即刻信息   |
|x-jike-access-token   |string    |token信息。后续请求都需要，注意保存   |
|x-jike-refresh-token   |string    |refresh-token。刷新token时需要，注意保存   |
|...   |...    |...   |


#### 示例

> 地址：https://www.example.com/login

参数

```javascript
{
  "areaCode": "+86",
  "verifyCode": "1234",
  "mobilePhoneNumber": "13111111111"
}
```

响应

``` javascript
{
	"code": 200,
	"msg": "OK",
	"data": {
		"data": {
			"data": {
				"isSignUp": false,
				"showNewbieGuide": true,
				"newbieGuideFeatureGroup": "SKIP",
				"subscriptionGuideFeatureGroup": "DISABLED",
				"user": {
					"type": "USER",
					"uid": "YOUR-UID",
					"avatar": {
						"picture": {
							"picUrl": "https://www.example.com",
							"largePicUrl": "https://www.example.com",
							"middlePicUrl": "https://www.example.com",
							"smallPicUrl": "https://www.example.com",
							"thumbnailUrl": "https://www.example.com",
							"format": "jpeg",
							"width": 132,
							"height": 132
						}
					},
					"nickname": "YOUR-NICKNAME",
					"isNicknameSet": true,
					"gender": "MALE",
					"isCancelled": false,
					"readTrackInfo": {},
					"ipLoc": "广东",
					"birthYear": 1998,
					"industry": "互联网/IT",
					"wechatUserInfo": {
						"nickName": "YOUR-WECHAT-NICKNAME"
					},
					"phoneNumber": {
						"mobilePhoneNumber": "131****1111",
						"areaCode": "+86"
					},
					"phoneNumberNeeded": false,
					"jikeUserInfo": {
						"nickname": "YOUR-JIKE-NICKNAME"
					},
					"isBlockedByViewer": false,
					"isInvited": true,
					"userCanDebug": false
				}
			}
		},
		"x-jike-access-token": "YOUR-ACCESS-TOKEN",
		"x-jike-refresh-token": "YOUR-REFRESH-TOKEN"
	}
}
```