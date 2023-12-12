# coding=utf8
import requests, sys 

help = '''
download 远程文件名 本地文件路径
'''

if len(sys.argv) != 3:
    print(help)
    exit(1)
file_name = sys.argv[1]
download_name = sys.argv[2]
# file_name = "config.yaml-202311201846452768"
# download_name = "/tmp/ingress_monitor/config5.yaml"


try:
    login_url = "http://127.0.0.1/opsv/auth/api/login"
    login_param = {
        "username": "zhhcheng",
        "password": ""
    }
    login_res = requests.post(url=login_url, json=login_param)
    login_result = login_res.json()
    if login_result["status"] != "success":
        print("Login failed")
        exit(1)
    opst = login_res.cookies["opst"]

    download_url = "http://127.0.0.1:5000/ingress_monitor/conf/download"
    header = {"Cookie": "opst=%s" % opst, 'user-Agent': 'Python-http-client/1.1 ywb-custom'}
    download_res = requests.get(url=download_url, headers=header, params="file_name=%s" % (file_name, ))
    with open("%s" %(download_name, ), "wb") as f:
        f.write(download_res.content)
except Exception as e:
    print("Download failed")
    exit(1)



