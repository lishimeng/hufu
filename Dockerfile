FROM golang:1.18 as build
ARG NAME
ARG VERSION
ARG COMMIT
ARG BUILD_TIME
ARG MAIN_PATH
ARG COMPILER="$(go version)"
ENV GOPROXY=https://goproxy.cn,direct
ENV LDFLAGS=" \
    -X 'github.com/lishimeng/app-starter/version.AppName=${NAME}' \
    -X 'github.com/lishimeng/app-starter/version.Version=${VERSION}' \
    -X 'github.com/lishimeng/app-starter/version.Commit=${COMMIT}' \
    -X 'github.com/lishimeng/app-starter/version.Build=${BUILD_TIME}' \
    -X 'github.com/lishimeng/app-starter/version.Compiler=${COMPILER}' \
    "
WORKDIR /release
ADD . .
RUN go mod download && go mod verify
RUN echo "go build -v --ldflags "${LDFLAGS}" -o ${NAME} ${MAIN_PATH}"
RUN go build -v --ldflags "${LDFLAGS}" -o ${NAME} ${MAIN_PATH}

FROM ubuntu:22.04 as prod
ARG NAME
EXPOSE 80/tcp
WORKDIR /
COPY --from=build /release/${NAME} /
RUN ln -s /${NAME} /app
CMD [ "/app"]
