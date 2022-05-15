package components

func WebIDPFileSrc(idpSvc string) []byte {

	src := ` 
##### Stage 1

FROM node:lts-buster as builder

WORKDIR /app

# Copy project files to the docker image
COPY . .

# if you prefer npm, replace the above command with
RUN npm install @angular/cli@latest -g

# FOR NPM
RUN npm install

# SET ENVIRONMENT VARIABLES
ENV ENVIRONMENT=production
ENV AUTH_SVC=` + idpSvc + `
# Build Angular Application in Production
RUN ng build --prod

#### STAGE 2
#### Deploying the application

FROM nginx:alpine

VOLUME  /var/cache/nginx

# Copy the build files from the project
# replace "angular-docker-environment-variables" with your angular project name
COPY --from=builder /app/dist/release /usr/share/nginx/html

# Copy Nginx Files
COPY --from=builder /app/.docker/.config/nginx.conf /etc/nginx/conf.d/default.conf

# EXPOSE Port 8888
EXPOSE 8888
`
	return []byte(src)

}
