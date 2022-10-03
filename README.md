# EverPhoto Checkin

使用 Github Action 完成时光相册的每日签到和每日任务。

## Secrets

| KEY     | Value                                                        |
| ------- | ------------------------------------------------------------ |
| TOKEN    | 在 `authorization` 项中的 `Bearer` 值 |

## 如何获取 token

### 方法一（推荐）

1. 登录 https://www.everphoto.cn
2. 按下 F12
3. 如下图所示获取
   Application => Cookies => access_token

![image](https://user-images.githubusercontent.com/29673994/172185521-ec8af7d7-61a0-4a90-84a5-8f19fdfcb0be.png)

### 方法二

1. Fork 本仓库
2. 以此点击 Action => GetEverPhotoToken => Run workflow
3. 输入用户名和密码

![image](https://user-images.githubusercontent.com/29673994/172188331-7459c9b1-6b0e-4e33-b234-253431df6f9a.png)

4. 在日志 log 中查看 token
5. **删除此 workflow （防止他人盗用 token）**

![image](https://user-images.githubusercontent.com/29673994/172188063-a38588c2-b415-40a1-a2ad-8f996a92e6d0.png)

