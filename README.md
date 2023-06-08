# 功能列表

- basic_convert:
  - 常见基础数据类型互转，数组互转


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