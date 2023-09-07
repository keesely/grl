# get
```bash
curl 'https://apiv2.hengfenglife.com.cn/mch/coupons/stocks?appends=global_max_coupon_per_user_by_type7_by_month&page=1' \
  -H 'authority: apiv2.hengfenglife.com.cn' \
  -H 'accept: application/json, text/plain, */*' \
  -H 'accept-language: zh,en-US;q=0.9,en;q=0.8,zh-Hans;q=0.7' \
  -H 'authorization: APPCODE 4a39ec89ecfd46de85d264469230b816' \
  -H 'cache-control: no-cache' \
  -H 'dnt: 1' \
  -H 'origin: https://admin.hengfenglife.com.cn' \
  -H 'pragma: no-cache' \
  -H 'referer: https://admin.hengfenglife.com.cn/' \
  -H 'sec-ch-ua: "Chromium";v="116", "Not)A;Brand";v="24", "Microsoft Edge";v="116", "Vivaldi";v="6.2"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-site' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36' \
  -H 'x-ca-stage: RELEASE' \
  -H 'x-ticket: eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IlViYVFpdHQ5MUF5WTQwMGMifQ.eyJqdGkiOiJleUowZVhCbElqb3pMQ0pwWkNJNk15d2lkR2x0WlNJNk1UWTRNalkwT1RVMk1TNDBNek15TXpNc0lsOXlZVzVrSWpveU16RTJPVGw5IiwiaXNzIjoiSGVuZ0ZlbmdMaWZlIiwiaWF0IjoxNjgyNjQ5NTYxLCJhdWQiOiIzIiwiZXhwIjoxNzAxODgzMTYxLCJvdl9pZCI6ImtlZS1RaSJ9.ceQt7gQ9I2FlLs5h04l3_ecqMK6rtbhb8YYXTilhuOzZ4y_G2UUyIvcUC_w9VNaLxvbu3WQgQrH9ibVmTZmMrrDul0lfRacAkwnTSn3vPyTZwtacOlzJfbimpw0c_X_sM2CAaEpOpZTTLi237iuUlULLu0w5Agg4HIyr4pXUUN8W0rNGSpDwioy6bd2UtBp5te4IClTlUH6h3R5MJhl8h2WVFau7nktqhZYCM0naOJd9dntNgdA8olP3ZAV4nNlKJ50AxxU1Kc5HeTuQv0hyMKY3f8t6sqQ69Pn9_etF7MF_YDzdJZur6ZBFOyECpuNJQktWV8gTKsTkWy0avLCoSQ' \
  --compressed
```

# put
```bash
curl 'https://apiv2.hengfenglife.com.cn/mch/merchants/116' \
  -X 'PUT' \
  -H 'authority: apiv2.hengfenglife.com.cn' \
  -H 'accept: application/json, text/plain, */*' \
  -H 'accept-language: zh,en-US;q=0.9,en;q=0.8,zh-Hans;q=0.7' \
  -H 'authorization: APPCODE 582fccd04e7b495cb8baec3a6b027fb4' \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'dnt: 1' \
  -H 'origin: http://admin.stage.hengfenglife.com.cn' \
  -H 'pragma: no-cache' \
  -H 'referer: http://admin.stage.hengfenglife.com.cn/' \
  -H 'sec-ch-ua: "Chromium";v="116", "Not)A;Brand";v="24", "Microsoft Edge";v="116", "Vivaldi";v="6.2"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: cross-site' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36' \
  -H 'x-ca-stage: TEST' \
  -H 'x-ticket: eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IlViYVFpdHQ5MUF5WTQwMGMifQ.eyJqdGkiOiJleUowZVhCbElqb3pMQ0pwWkNJNk15d2lkR2x0WlNJNk1UWTVOREExTmpNd05TNDBOalU0T1RNc0lsOXlZVzVrSWpvNU5qazNNRFY5IiwiaXNzIjoiSGVuZ0ZlbmdMaWZlIiwiaWF0IjoxNjk0MDU2MzA1LCJhdWQiOiIzIiwiZXhwIjoxNjk0MDU4MTA1LCJvdl9pZCI6ImtlZS1RaSJ9.uqCnzRwkkU0lN9nCP-l22R_KMAm5bJfUcAZmPp9U5-qm3lhptpjfGcATAiOhln_WptOpCiq1NtcKcb0G01smJ8abSvEJ1cbwYOpk_AF6VR0HQlg0euCnFQNRWtsGL_AIU7VLs0-ubBMqLy0sX0BDFdRBqvdjjgIomogi0d5rS__WDOGC7SINTo8FLTAH1deDJVVpDvZgW3Xqnm48l_0k8IhVxO9r-tEH8_r_fGjH5yQBRcKYRurvHr-NOyo-eGlHWNz5sxrHppzixmxiBcUTh6Eby4wkllxYwwZaM6GVoAA_tEv4ATMVQSmb1qUaHJ0oB1QLpytNSeG8NwLRhORdwg' \
  --data-raw '{"id":116,"name":"Stage桃城区瑶雅食品店","type":2,"status":1,"logo":"https://cdn.hengfenglife.com.cn/general/0/2022/02/16/3852c9ecec407b9bdf72fccb59ec3131bb8842bc.jpeg","description":"\nOPYID: 161-720-1477\n\n匹配任何包含至少一个 n 的字符串。\n\n例如，/a+/ 匹配 \"candy\" 中的 \"a\"，\"caaaaaaandy\" 中所有的 \"a\"。\n\nn*\t\n匹配任何包含零个或多个 n 的字符串。\n\n例如，/bo*/ 匹配 \"A ghost booooed\" 中的 \"boooo\"，\"A bird warbled\" 中的 \"b\"，但是不匹配 \"A goat grunted\"。\n","address":"CDAA","expired_at":"2022-01-01 00:00:00","opening_times":["08:00","20:00"],"pictures":["https://tva1.sinaimg.cn/mw1024/9d52c073gy1gz49cths8gj20xc1cb10h.jpg","https://t7.baidu.com/it/u=963301259,1982396977&fm=193&f=GIF","https://t7.baidu.com/it/u=737555197,308540855&fm=193&f=GIF","https://t7.baidu.com/it/u=1297102096,3476971300&fm=193&f=GIF","https://tva1.sinaimg.cn/mw1024/9d52c073gy1gz49cv4j1uj20xc1cbtg4.jpg","https://tva1.sinaimg.cn/mw1024/9d52c073gy1gz49cy43q0j20yi0lcwk6.jpg"],"attach":{"mchid":"1616865058","sort_id":9527,"showable":1,"categorys":[],"sub_mchid":"1642060288","printer_sn":null,"printer_key":null,"printer_ret":["922668306成功"],"delivery_times":[["08:00","20:00"],["08:00","20:00"]],"min_delivery_fee":1213},"shop_numbers":{"10000011601":{"shop_number":"10000011601","shop_name":"Stage桃城区瑶雅食品店-1","close":true},"10000011602":{"shop_number":"10000011602","shop_name":"Stage桃城区瑶雅食品店-2","close":true},"10000011603":{"shop_number":"10000011603","shop_name":"Stage桃城区瑶雅食品店-3","close":true}},"can_sender":2,"lat":37.754213,"lng":115.666901,"created_at":"2021-12-29 14:31:57","updated_at":"2023-07-06 11:09:09","join_code":"789d386358aee910627a20421001be2f","users_count":3,"canPay":true,"can_sender_bool":true,"min_delivery_price":12.13,"users":[{"user_id":8641,"state":1,"role":2,"account":null,"created_at":"2022-08-24 10:41:07","updated_at":"2022-10-14 15:51:53","account_type":"wechat","name":null,"role_desc":"成员","user":{"id":8641,"username":"n_136660364751","nickname":"Kee2021","realname":"郭","status":1,"mobile":"13666036475","email":null,"sex":1,"level":1,"user_type":1,"supper":-1,"estate_role":1}},{"user_id":17119,"state":1,"role":1,"account":null,"created_at":"2022-03-18 16:34:46","updated_at":"2022-03-18 16:34:46","account_type":"wechat","name":null,"role_desc":"管理员","user":null},{"user_id":18696,"state":1,"role":1,"account":null,"created_at":"2022-10-19 17:33:35","updated_at":"2022-10-19 17:33:35","account_type":"wechat","name":null,"role_desc":"管理员","user":null}],"mchid":"1616865058","sub_mchid":"1642060288","pay":true,"categorys":[],"printer":-1,"show":1,"shop_items":[{"close":true,"shop_name":"Stage桃城区瑶雅食品店-1","shop_number":"10000011601","open":false},{"close":true,"shop_name":"Stage桃城区瑶雅食品店-2","shop_number":"10000011602","open":false},{"close":true,"shop_name":"Stage桃城区瑶雅食品店-3","shop_number":"10000011603","open":false}]}' \
  --compressed
```
