module github.com/Kydaa/bit/service/user

go 1.14

replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4

require (
	github.com/2bitlab/bit_infra v0.0.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.14
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.24.0
)
