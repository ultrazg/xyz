### 更新收听数据概览

更新收听数据概览与收听历史。轮询调用

#### 请求地址

> /mileage_update

#### 请求方式

> POST

#### 支持格式

> JSON

#### 请求头

| 参数                | 必填 | 类型   | 说明                |
| :------------------ | :--- | :----- | ------------------- |
| x-jike-access-token | true | string | x-jike-access-token |

#### 请求参数

| 参数                           | 必填  | 类型    | 说明                                                         |
| :----------------------------- | :---- | :------ | ------------------------------------------------------------ |
| tracking                       | true  | array   | 请求体                                                       |
| tracking.eid                   | true  | string  | 单集 id                                                      |
| tracking.pid                   | true  | string  | 节目 id                                                      |
| tracking.startPlayingTimestamp | true  | number  | 开始播放时的时间戳                                           |
| tracking.endPlayingTimestamp   | true  | number  | 当前播放时的时间戳。再次调用时将此值赋值给 **startPlayingTimestamp** |
| tracking.isSpeaker             | false | boolean | 是否外放                                                     |
| tracking.isOffline             | false | boolean | 是否已下载单集                                               |
| tracking.isTrial               | false | boolean | 是否属于试听单集                                             |
| tracking.withSpeed             | true  | number  | 播放速度（1）                                                |

#### 返回字段

无


#### 示例

> 地址：https://www.example.com/mileage_update

参数

```javascript
{
  "tracking": [
        {
            "isOffline": false,
            "isSpeaker": true,
            "withSpeed": 1,
            "isTrial": true,
            "eid": "67fef8f0cdd692da1591dc98",
            "startPlayingTimestamp": 1744899862591.384,
            "endPlayingTimestamp": 1744899894736.4399,
            "pid": "600834f0bd898cad7ea07801"
        }
    ],
}
```

响应

```javascript
{
  code: 200,
  data: {},
  msg: "OK",
}
```