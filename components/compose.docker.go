package components

func DockerComposeSrc() []byte {

	src := `version: '3.1'
services:
  redis-svc:
    image: redis
    restart: always
    ports:
      - 6379:6379
    volumes:
      - "/etc/timezone:/etc/timezone:ro"
      - "/etc/localtime:/etc/localtime:ro"
  idp:
    image: idp
    container_name: idp_ui
    restart: always
    ports:
      - 8888:8888
    volumes:
      - "/etc/timezone:/etc/timezone:ro"
      - "/etc/localtime:/etc/localtime:ro"
  auth-svc:
    image: auth-svc
    restart: always
    ports:
      - 4000:4000
    volumes:
      - /opt/data/sqldb:/go/src/app/data/db
      - "/etc/timezone:/etc/timezone:ro"
      - "/etc/localtime:/etc/localtime:ro"
    depends_on:
      - "redis-svc"
`

	return []byte(src)

}
