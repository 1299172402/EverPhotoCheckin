import requests
import time
import os

def main():
    ts = int(time.time()*1000)
    baseurl = 'https://openapi.everphoto.cn/sf/3/v4/PostCheckIn'
    data = os.environ['data']
    data['_rticket'] = str(ts)
    params = '&'.join(f'{k}={v}' for k, v in data.items())
    url = baseurl + '?' + params
    headers = os.environ['headers']
    res = requests.post(url, headers=headers, data='{}')
    print(res.text)

if __name__ == '__main__':
    main()
