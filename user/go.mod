module github.com/sabnak227/jwt-demo/user

go 1.13

replace github.com/sabnak227/jwt-demo/util => ../util

require (
	github.com/go-kit/kit v0.10.0
	github.com/go-ozzo/ozzo-validation/v4 v4.2.1
	github.com/gogo/protobuf v1.3.1
	github.com/gorilla/mux v1.7.4
	github.com/jinzhu/gorm v1.9.15
	github.com/metaverse/truss v0.1.0
	github.com/pkg/errors v0.9.1
	github.com/sabnak227/jwt-demo/auth v0.0.0-20200812201839-e9805e935abe
	github.com/sabnak227/jwt-demo/util v0.0.0-20200812201839-e9805e935abe
	github.com/sirupsen/logrus v1.6.0
	github.com/streadway/amqp v1.0.0
	golang.org/x/sys v0.0.0-20200812155832-6a926be9bd1d // indirect
	google.golang.org/genproto v0.0.0-20200812160120-2e714abc8b50 // indirect
	google.golang.org/grpc v1.31.0
)
