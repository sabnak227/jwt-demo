module github.com/sabnak227/jwt-demo/auth

go 1.13

replace github.com/sabnak227/jwt-demo/scope => ../scope

replace github.com/sabnak227/jwt-demo/user => ../user

replace github.com/sabnak227/jwt-demo/util => ../util

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/go-ozzo/ozzo-validation/v4 v4.2.1
	github.com/gogo/protobuf v1.3.1
	github.com/google/uuid v1.0.0
	github.com/gorilla/mux v1.7.4
	github.com/jinzhu/gorm v1.9.15
	github.com/metaverse/truss v0.1.0
	github.com/pkg/errors v0.9.1
	github.com/sabnak227/jwt-demo/scope v0.0.0-20200808201311-ef1f6b76d9a3
	github.com/sabnak227/jwt-demo/user v0.0.0-20200808201311-ef1f6b76d9a3
	github.com/sabnak227/jwt-demo/util v0.0.0-20200808193236-c280dbae12ae
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200808120158-1030fc2bf1d9 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98 // indirect
	google.golang.org/grpc v1.31.0
	google.golang.org/protobuf v1.25.0 // indirect
)
