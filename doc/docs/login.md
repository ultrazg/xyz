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
| verifyCode        | true  | string | 验证码        |
| areaCode          | false | string | 区号，默认+86 |

#### 返回字段

| 返回字段             | 类型   | 说明                                     |
| :------------------- | :----- | :--------------------------------------- |
| type            | string | 用户角色                                 |
| uid             | string | 用户 uid                                 |
| avatar          | object | 用户头像信息                             |
| nickname        | string | 用户昵称                                 |
| wechatUserInfo  | object | 绑定的微信信息                           |
| jikeUserInfo    | object | 绑定的即刻信息                           |
| x-jike-access-token  | string | 后续请求都需要，注意保存      |
| x-jike-refresh-token | string | 刷新 token 时需要，注意保存 |
| ...                  | ...    | ...                                      |

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

```javascript
{
	"code": 200,
	"msg": "OK",
	"data": {
		"data": {
			"type": "USER",
			"uid": "YOUR-UID",
			"avatar": {
				"picture": {
					"picUrl": "https://www.example.jpg",
					"largePicUrl": "https://www.example.jpg@large",
					"middlePicUrl": "https://www.example.jpg@middle",
					"smallPicUrl": "https://www.example.jpg@small",
					"thumbnailUrl": "https://www.example.jpg@thumbnail",
					"format": "jpeg",
					"width": 1003,
					"height": 1003
				}
			},
			"nickname": "YOUR-NICKNAME",
			"isNicknameSet": true,
			"bio": "YOUR-BIO",
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
			"isBlockedByViewer": false
		},
		"x-jike-access-token": "YOUR-ACCESS-TOKEN",
		"x-jike-refresh-token": "YOUR-REFRESH-TOKEN"
	}
}
```
