# 功能列表

- basic_convert:
    - 常见基础数据类型互转，数组互转
- base64
    - 常见base64加密，解密
- json
    - 序列化，反序列化
- md5
    - md5加密
- random
    - 区间随机
- rsa
    - 公钥加密
    - 私钥解密
- time
    - 时间格式化通用大类
- uuid
    - 36位uuid
    - 区间长度uuid

如果大家觉得好用,右上角帮忙点个star吧。:)
> 收集日常好用的golang工具包。

# 联系我们

- 技术支持/咨询请联系作者QQ: 554486586
- 作者邮箱: 554486586@qq.com

##### 单元测试

```
go test -coverpkg=./... -coverprofile=coverage.data -timeout=5s ./...
go tool cover -html=coverage.data -o coverage.html
````