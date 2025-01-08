# golang requests library benchmark

## Introduction
To compare the performance of golang and python request libraries, we mainly compare the following request libraries:
- [gospider007/requests](github.com/gospider007/requests)
- [go-resty](github.com/go-resty/resty)
- [imroc/req](github.com/imroc/req)
- [wangluozhe/requests](github.com/wangluozhe/requests)

- [psf/requests](https://github.com/psf/requests)
- [httpx](https://github.com/encode/httpx)
- [curl_cffi](https://github.com/yifeikong/curl_cffi)

### Random 1k data 10w requests
| Library name | time consuming |
| --- | --- |
| gospider007/requests | 6.9499259s |
| go-resty  | 7.1604138s |
| imroc/req | 7.2530969s | 
| wangluozhe/requests |  8.4787594s |
| curl_cffi | 26.215140342712402s |
| httpx | 61.01818656921387s |
| psf/requests | 76.01117539405823s |

### Random 10k data 10w requests
| Library name | time consuming | 
| --- | --- |
| gospider007/requests | 9.4316189s |
| imroc/req | 10.1804427s | 
| go-resty  |  10.2600992s |
| wangluozhe/requests | 11.7542946s|
| curl_cffi | 27.91108274459839s |
| httpx | 64.91894054412842s |
| psf/requests | 77.11325931549072s |

### Random 100k data 10w requests
| Library name | time consuming |
| --- | --- |
| gospider007/requests | 19.228713s |
| imroc/req |31.6547403s | 
| go-resty  |  32.5609081s |
| wangluozhe/requests | 33.4119696s|
| curl_cffi | 43.912320137023926s |
| httpx | 84.05615258216858s |
| psf/requests | 92.94585227966309s |

## Results
[gospider007/requests](https://github.com/gospider007/requests) > [imroc/req](github.com/imroc/req) > [go-resty](github.com/go-resty/resty) > [wangluozhe/requests](github.com/wangluozhe/requests) > [curl_cffi](https://github.com/yifeikong/curl_cffi) > [httpx](https://github.com/encode/httpx) > [psf/requests](https://github.com/psf/requests)