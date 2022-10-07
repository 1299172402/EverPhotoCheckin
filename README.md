# EverPhoto Checkin

完成时光相册的每日签到和每日任务。

注意：每日任务需要使用中国IP。

## 环境变量

| Name    | Value                                                        |
| ------- | ------------------------------------------------------------ |
| token | 在 `authorization` 项中的 `Bearer` 值 |

## 如何获取 token

### 方法一（推荐）

1. 登录 https://www.everphoto.cn
2. 按下 F12
3. 如下图所示获取
   Application => Cookies => access_token

![image](https://user-images.githubusercontent.com/29673994/172185521-ec8af7d7-61a0-4a90-84a5-8f19fdfcb0be.png)

### 方法二

   运行 EverPhotoSign.py 中的 sign() 函数，获取 token 。

