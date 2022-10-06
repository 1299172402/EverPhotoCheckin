import os
import time
import requests

token = os.environ['token']
headers = {
    "content-type": "application/json",
    "authorization": f"Bearer {token}",
    "user-agent": "EverPhoto/5.3.0",
    # "application": "tc.everphoto",
    # "host": "openapi.everphoto.cn",
}

# 签到
CHECKIN_URL = "https://openapi.everphoto.cn/sf/3/v4/PostCheckIn"
res = requests.post(CHECKIN_URL, headers=headers).json()
print('[签到状态]')
print(res)

# 签到报错
if res['data']['checkin_result']==False:
    raise Exception('CHECKIN ERROR')

# 延时
time.sleep(5)

# 任务
task = {'star', 'hide', 'add_to_album', 'remark'}

for mission_id in task:
    # 任务（去完成）
    MISSION_URL = "https://openapi.everphoto.cn/sf/3/v4/MissionReport"
    res = requests.post(MISSION_URL, headers=headers, json={"mission_id": mission_id}).json()
    print(f'[任务 {mission_id} 去完成]')
    print(res)
    if res['code']!=0:
        raise Exception('MISSION ERROR')
    # 延时
    time.sleep(5)

    # 任务（领取奖励）
    REWARD_URL = "https://openapi.everphoto.cn/sf/3/v4/MissionRewardClaim"
    res = requests.post(REWARD_URL, headers=headers, json={"mission_id": mission_id}).json()
    print(f'[任务 {mission_id} 领取奖励]')
    print(res)
    if res['code']!=0:
        raise Exception('REWARD ERROR')
    # 延时
    time.sleep(5)
