cd ./ui
rm -rf src/lib/proto
mkdir src/lib/proto
npm run proto

cd ../

rm -rf internal/proto
mkdir -p internal/proto
protoc --go_out=./internal/proto \
 --go_opt=paths=source_relative \
 --go-grpc_out=./internal/proto \
 --go-grpc_opt=paths=source_relative \
 --proto_path=./proto \
 ./proto/*.proto