### 发送短信验证码

用于接收验证码，确保手机号已经注册/绑定小宇宙

> 注意：请勿频繁调用该接口！

#### 请求地址

> /sendCode

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求参数

|参数|必填|类型|说明|
|:-----  |:-------|:-----|:----                               |
|mobilePhoneNumber    |true    |string|手机号                          |
|areaCode    |false    |string   |区号，默认+86|

#### 返回字段

无

#### 示例

> 地址：https://www.example.com/sendCode

参数
```json
{
	"mobilePhoneNumber":"13100000000",
	"areaCode":"+86"
}
```
响应
``` javascript
{
    "code": 200,
    "msg": "OK",
    "data": {},
}
```
