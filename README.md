# function list

- basic_convert:
    - [X] common basic data type conversion, array conversion
- base64
    - [X] common base64 encryption, decryption
- json
    - [X] serialization, deserialization
- md5
    - [X] md5 encryption
- random
    - [X] interval random
- rsa
    - [X] public key encryption
    - [X] private key decryption
- time
    - [X] time formatting general class
- uuid
    - [X] 36 bit uuid
    - [X] interval length uuid
- goroutines
    - [X] Catch panic errors, enable security
- db
    - redis
      > GUI RedisInsight is recommended https://redis.com/redis-enterprise/redis-insight/
        - [X] initialization method
    - mongo
      > Recommended to use GUI Studio3T https://studio3t.com/
        - [X] initialization method
    - mysql
      > GUI TablePlus is recommended https://tableplus.com/
        - [X] initialization method
    - postgresql
      > GUI TablePlus is recommended https://tableplus.com/
        - [X] initialization method
- log
    - [X] go.uber.org/zap
    - [X] github.com/sirupsen/logrus
- monitor
    - [X] DingTalk
- config
    - [X] yaml
- cron github.com/robfig/cron
- db: gorm.io/gorm
    - log plugin
    - timeout plugin
- http_client: native processing
    - [X] log plugin
    - [X] timeout plugin
- grpc_server:
    - log plugin
    - metric plugin
    - recovery plugin
    - timeout plugin
- grpc_client:
    - log plugin
    - metric plugin
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