# EverPhoto Checkin

Use Github Action to check-in EverPhoto everyday.

## Secrets

| KEY     | Value                                                        |
| ------- | ------------------------------------------------------------ |
| TOKEN    | The `Bearer` in the `authorization`. |

## How to get token

### The best way

1. Login https://www.everphoto.cn
2. Press F12
3. Get it as explained in the picture below
   Application => Cookies => access_token

![image](https://user-images.githubusercontent.com/29673994/172185521-ec8af7d7-61a0-4a90-84a5-8f19fdfcb0be.png)

### Other way

1. Fork this repository
2. Go to Action => GetEverPhotoToken => Run workflow
3. Enter your username and password

![image](https://user-images.githubusercontent.com/29673994/172188331-7459c9b1-6b0e-4e33-b234-253431df6f9a.png)

4. View the token in the log
5. **REMOVE the workflow**

![image](https://user-images.githubusercontent.com/29673994/172188063-a38588c2-b415-40a1-a2ad-8f996a92e6d0.png)

