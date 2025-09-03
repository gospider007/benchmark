# golang requests library benchmark

## Introduction
To compare the performance of golang and python request libraries, we mainly compare the following request libraries:
- [gospider007/fingerproxy](github.com/gospider007/fingerproxy)
- [azuretls-client](https://github.com/Noooste/azuretls-client)
- [wangluozhe/requests](github.com/wangluozhe/requests)
- [imroc/req](github.com/imroc/req)
- [go-resty](github.com/go-resty/resty)

### Random 1k data 1k requests
| Library name | time consuming |
| --- | --- |
| Fingerproxy | 2.26485075s |
| AzureTest | 1.295575208s |
| WangluozheRequest | 1.455702625s |
| ImrocReq | 1.333406458s |
| GoResty | 1.285543375s |
| Http | 1.214584083s |

---

### Random 1k data 10k requests
| Library name | time consuming |
| --- | --- |
| Fingerproxy | 4.163065s |
| AzureTest | 2.439226875s |
| WangluozheRequest | 2.293516959s |
| ImrocReq | 2.279090209s |
| GoResty | 2.217924125s |
| Http | 2.185611417s |

---

### Random 1k data 100k requests
| Library name | time consuming |
| --- | --- |
| Fingerproxy | 9.732467167s |
| AzureTest | 8.879041542s |
| WangluozheRequest | 7.8297485s |
| ImrocReq | 9.092916042s |
| GoResty | 9.015154208s |
| Http | 9.143848291s |
