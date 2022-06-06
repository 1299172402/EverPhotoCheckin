import requests
import json
import os

def checkin(token):
    url = 'https://openapi.everphoto.cn/sf/3/v4/PostCheckIn'
    headers = {
        "user-agent": "EverPhoto/4.5.0",
        "content-type": "application/json",
        "authorization": f"Bearer {token}",
    }
    res = json.loads(requests.post(url, headers=headers).text)
    print(res)
    if res['data']['checkin_result']!=True:
        raise Exception('Check-in ERROR')

if __name__ == '__main__':
    checkin(os.environ['token'])
