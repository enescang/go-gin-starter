module github.com/enescang/go-gin-starter

go 1.16

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/golang-jwt/jwt v3.2.1+incompatible
	github.com/joho/godotenv v1.3.0
	github.com/stretchr/testify v1.6.1
	go.mongodb.org/mongo-driver v1.5.4
)

replace github.com/enescang/go-gin-starter/db => ./db
replace github.com/enescang/go-gin-starter/utils => ./utils
replace github.com/enescang/go-gin-starter/models => ./models
replace github.com/enescang/go-gin-starter/controllers => ./controllers
