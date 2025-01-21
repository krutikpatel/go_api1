go mod init api1

# Install required packages
go get github.com/gin-gonic/gin
go get github.com/sirupsen/logrus
go get github.com/spf13/viper    # for configuration
go get -u gorm.io/gorm          # if using GORM
go get -u gorm.io/driver/postgres  # for PostgreSQL driver
go get github.com/spf13/viper

# After adding dependencies, tidy up modules
go mod tidy