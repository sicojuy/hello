#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
docker build -t docker.gf.com.cn/dev/go/hello .
docker push docker.gf.com.cn/dev/go/hello
