import requests
import httpx
from curl_cffi import requests as curl
import time
def requestsTest(href,count):
    session=requests.session()
    sucess=0
    total=100000
    t=time.time()
    for _ in range(total):
        r=session.get(href)
        if len(r.content)==count:
            sucess+=1
    if sucess!=total:
        print("test failed")
    else:
        print("requests time:",time.time()-t)
def httpxTest(href,count):
    session=httpx.Client()
    sucess=0
    total=100000
    t=time.time()
    for _ in range(total):
        r=session.get(href)
        if len(r.content)==count:
            sucess+=1
    if sucess!=total:
        print("test failed")
    else:
        print("httpx time:",time.time()-t)
def curlTest(href,count):
    session=curl.Session()
    sucess=0
    total=100000
    t=time.time()
    for _ in range(total):
        r=session.get(href)
        if len(r.content)==count:
            sucess+=1
    if sucess!=total:
        print("test failed")
    else:
        print("curl time:",time.time()-t)


requestsTest("http://127.0.0.1:3334/1k",1024)
httpxTest("http://127.0.0.1:3334/1k",1024)
curlTest("http://127.0.0.1:3334/1k",1024)


requestsTest("http://127.0.0.1:3334/10k",10240)
httpxTest("http://127.0.0.1:3334/10k",10240)
curlTest("http://127.0.0.1:3334/10k",10240)

requestsTest("http://127.0.0.1:3334/100k",102400)
httpxTest("http://127.0.0.1:3334/100k",102400)
curlTest("http://127.0.0.1:3334/100k",102400)


