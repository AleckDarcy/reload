#go mod download
#go mod vendor

# reload grpc
rm -rf vendor/google.golang.org/grpc/
cp -r ../../../../go/google.golang.org/grpc vendor/google.golang.org/grpc

# reload go-grpc-middleware
rm -rf vendor/github.com/grpc-ecosystem/go-grpc-middleware/
cp -r ../../../../go/github.com/grpc-ecosystem/go-grpc-middleware/ vendor/github.com/grpc-ecosystem/go-grpc-middleware/

# reload core
rm -rf vendor/github.com/AleckDarcy/reload/
cp -r ../../../../../core vendor/github.com/AleckDarcy/reload/