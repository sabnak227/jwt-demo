module github.com/sabnak227/jwt-demo/scope

go 1.13

replace github.com/sabnak227/jwt-demo/util => ../util

require (
	github.com/go-kit/kit v0.10.0
	github.com/go-redis/redis/v8 v8.0.0-beta.7
	github.com/gogo/protobuf v1.3.1
	github.com/gorilla/mux v1.7.4
	github.com/jinzhu/gorm v1.9.16
	github.com/metaverse/truss v0.1.0
	github.com/pkg/errors v0.9.1
	github.com/sabnak227/jwt-demo/auth v0.0.0-20200819191209-9d9043b16789
	github.com/sabnak227/jwt-demo/user v0.0.0-20200819191209-9d9043b16789
	github.com/sabnak227/jwt-demo/util v0.0.0-20200819180714-17e952fb120e
	github.com/sirupsen/logrus v1.6.0
	github.com/streadway/amqp v1.0.0
	google.golang.org/grpc v1.31.0
)
