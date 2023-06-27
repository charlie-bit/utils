# function list

- basic_convert
    - common basic data type conversion, array conversion
- base64
    - common base64 encryption, decryption
- json
    - serialization, deserialization
- md5
    - md5 encryption
- random
    - interval random
- rsa
    - public key encryption
    - private key decryption
- time
    - time formatting general class
- uuid
    - 36 bit uuid
    - interval length uuid
- goroutines
    - catch panic errors, enable security
- db
    - redis
      > GUI RedisInsight is recommended https://redis.com/redis-enterprise/redis-insight/
        - initialization method
    - mongo
      > Recommended to use GUI Studio3T https://studio3t.com/
        - initialization method
    - mysql
      > GUI TablePlus is recommended https://tableplus.com/
        - initialization method
    - postgresql
      > GUI TablePlus is recommended https://tableplus.com/
        - initialization method
- log
    - go.uber.org/zap
    - github.com/sirupsen/logrus
- monitor
    - DingTalk
- config
    - yaml
- cron github.com/robfig/cron
    - new cron job
    - show the example configuration file
- http_client: native processing
    - log plugin
    - timeout logic
- grpc_server
    - log plugin
    - recovery plugin
    - timeout plugin
- grpc_client:
    - log plugin
    - timeout plugin
- websockets
    - log plugin
    - timeout plugin

If you think it is easy to use, please click a star in the upper right corner. :)
> Collect golang toolkits that are useful for daily use.

##### unit test

```
go test -coverpkg=./... -coverprofile=coverage.data -timeout=5s ./...
go tool cover -html=coverage.data -o coverage.html
````