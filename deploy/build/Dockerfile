# 程序编译镜像
FROM golang:1.23 AS build

ENV GO111MODULE=on                               \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /go/release

ADD . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/ship-management ./cmd/ship-management

# 构建镜像
FROM alpine:3.20 AS prod

ENV TZ=Asia/Shanghai                                        \
    DEBIAN_FRONTEND=noninteractive                          \
    GIN_MODE=release                                        \
    ADDR=":8888"                                            \
    SQL_SERVER_DSN="sqlserver://sa:smart_ship123..@172.31.126.33:1433?database=ISM_Vsl_9977551&encrypt=disable"

# 在build阶段复制时区到
COPY --from=build /go/release/dist/ship-management    /ship-management

CMD [ "/ship-management"]