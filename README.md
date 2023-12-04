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
| go-resty  | 7.3855073s |
| imroc/req | 7.3961296s | 
| gospider007/requests | 7.7989687s |
| wangluozhe/requests | 8.9908882s |
| curl_cffi | 26.215140342712402s |
| httpx | 61.01818656921387s |
| psf/requests | 76.01117539405823s |

### Random 10k data 10w requests
| Library name | time consuming | 
| --- | --- |
| gospider007/requests | 10.3116086s |
| wangluozhe/requests | 12.402695s|
| imroc/req | 10.7307247s | 
| go-resty  |  10.7186106s |
| curl_cffi | 27.91108274459839s |
| httpx | 64.91894054412842s |
| psf/requests | 77.11325931549072s |

### Random 100k data 10w requests
| Library name | time consuming |
| --- | --- |
| gospider007/requests | 19.7566719s |
| wangluozhe/requests | 33.9291227s|
| imroc/req |30.3512156s | 
| go-resty  |  31.2344602s |
| curl_cffi | 43.912320137023926s |
| httpx | 84.05615258216858s |
| psf/requests | 92.94585227966309s |

