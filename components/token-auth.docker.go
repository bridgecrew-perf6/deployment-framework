package components

import "os"

func TokenAuthSrc(release string, host string, domain string, web string, secureMode string, redisHost string, dbhost string, user string, password string) []byte {

	src :=
		`FROM golang:latest AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
# second stage
FROM alpine:latest AS production
RUN apk add --no-cache libc6-compat
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
COPY --from=builder /app .
CMD ["./main"]
`

	return []byte(src)

}
