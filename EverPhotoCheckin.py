import requests
import time
import os
import json

def main():
    ts = int(time.time()*1000)
    baseurl = 'https://openapi.everphoto.cn/sf/3/v4/PostCheckIn'
    data = json.loads(os.environ['data'])
    data['_rticket'] = str(ts)
    params = '&'.join(f'{k}={v}' for k, v in data.items())
    url = baseurl + '?' + params
    headers = json.loads(os.environ['headers'])
    res = requests.post(url, headers=headers, data='{}')
    res = json.loads(requests.post(url, headers=headers, data='{}').text)
    print(res)
    if res['data']['checkin_result']!=True:
        raise Exception('Check-in ERROR')

if __name__ == '__main__':
    main()
