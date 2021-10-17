#go mod download
#go mod vendor

# reload go-grpc-middleware
cp -r ../../../../go/github.com/grpc-ecosystem/go-grpc-middleware/ vendor/github.com/grpc-ecosystem/go-grpc-middleware/

# reload core
cp -r ../../../../../core vendor/github.com/AleckDarcy/reload/