module github.com/wingfeng/backend

go 1.15

replace github.com/wingfeng/idx => ../server

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/bwmarrin/snowflake v0.3.0
	github.com/casbin/casbin/v2 v2.17.0
	github.com/casbin/xorm-adapter/v2 v2.1.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.1
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/spec v0.19.14 // indirect
	github.com/go-openapi/swag v0.19.12 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/labstack/gommon v0.3.0
	github.com/lestrrat/go-jwx v0.0.0-20180221005942-b7d4802280ae
	github.com/lestrrat/go-pdebug v0.0.0-20180220043741-569c97477ae8 // indirect
	github.com/lunny/log v0.0.0-20160921050905-7887c61bf0de
	github.com/magiconair/properties v1.8.4
	github.com/mattn/go-sqlite3 v1.14.5
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/spf13/viper v1.7.1
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.6.9
	github.com/ugorji/go v1.2.0 // indirect
	github.com/wingfeng/idx v0.0.1
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20201117144127-c1f2f97bffc9
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/sys v0.0.0-20201117222635-ba5294a509c7 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20201118030313-598b068a9102 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/guregu/null.v4 v4.0.0
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.9
)
