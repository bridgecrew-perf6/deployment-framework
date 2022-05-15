package components

import "os"

func TokenAuthSrc(release string, host string, domain string, web string, secureMode string, redisHost string, dbhost string, user string, password string) []byte {

	src :=
		`
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base gcc
ADD . /src
RUN cd /src && go build -ldflags "-s -w" -o main .
# final stage
FROM alpine
WORKDIR /app
ENV ADMIN_PASS=admin
ENV REDIS_HOST=` + redisHost + `
ENV REDIS_PORT=6379
ENV REDIS_PASSWORD=
ENV MYSQL_HOST=` + dbhost + `
ENV MYSQL_USER=` + user + `
ENV MYSQL_PASSWORD=` + password + `
ENV ACCESS_SECRET=` + os.Getenv("ACCESS_SECRET") + `
ENV REFRESH_SECRET=` + os.Getenv("REFRESH_SECRET") + `
ENV AESKEY=` + os.Getenv("AESKEY") + `
ENV ALLOWED=` + host + `
ENV IDP_UI=` + web + `
ENV COOKIE_DOMAIN=` + domain + `
ENV COOKIE_SECURE=` + secureMode + `
ENV PORT=4000
ENV GIN_MODE=` + release + `
EXPOSE 4000
COPY --from=build-env /src/main /app/
CMD ["./main"]
`

	return []byte(src)

}
