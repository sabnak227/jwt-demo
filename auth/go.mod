module github.com/sabnak227/jwt-demo/auth

go 1.13

replace github.com/sabnak227/jwt-demo/util => ../util

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/go-ozzo/ozzo-validation/v4 v4.2.1
	github.com/go-redis/redis/v8 v8.0.0-beta.7
	github.com/gogo/protobuf v1.3.1
	github.com/google/uuid v1.0.0
	github.com/gorilla/mux v1.7.4
	github.com/jinzhu/gorm v1.9.15
	github.com/lestrrat-go/jwx v1.0.3 // indirect
	github.com/metaverse/truss v0.1.0
	github.com/pkg/errors v0.9.1
	github.com/sabnak227/jwt-demo/scope v0.0.0-20200811004440-54a0fb6ac2fc
	github.com/sabnak227/jwt-demo/user v0.0.0-20200811004440-54a0fb6ac2fc
	github.com/sabnak227/jwt-demo/util v0.0.0-20200811004440-54a0fb6ac2fc
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200810151505-1b9f1253b3ed // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200808173500-a06252235341 // indirect
	google.golang.org/grpc v1.31.0
	google.golang.org/protobuf v1.25.0 // indirect
)
