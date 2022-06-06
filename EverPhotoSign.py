import requests
import hashlib
import json
import os

def sign(mobile, password):
    url = 'https://web.everphoto.cn/api/auth'
    headers = {
        "user-agent": "EverPhoto/4.5.0",
        "application": "tc.everphoto",
    }
    password = hashlib.md5(('tc.everphoto.' + password).encode()).hexdigest()
    data = {
        "mobile": mobile,
        "password": password,
    }
    res = json.loads(requests.post(url, headers=headers, data=data).text)
    return res['data']['token']

if __name__ == '__main__':
    print(sign(os.environ['mobile'], os.environ['password']))
