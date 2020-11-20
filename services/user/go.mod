module user

go 1.15

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/grpc v1.33.2
	gopkg.in/yaml.v2 v2.2.8 // indirect
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.6
	protobuf v0.0.0-00010101000000-000000000000
)

replace protobuf => ../protobuf
