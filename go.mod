// +heroku install ./... bitbucket.org/liamstask/goose/cmd/goose

module github.com/bryutus/brute

// +heroku goVersion go1.15
go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.2.0
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pressly/goose v2.7.0+incompatible
	github.com/stretchr/testify v1.5.1 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)
