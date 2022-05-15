package components

func ComposeBuildScript() []byte {

	src := `
echo "building token auth service"
cd auth-svc-deploy &&
docker build --tag auth-svc  . --no-cache &&
cd .. &&
echo "Getting Redis..."
docker pull redis &&
echo "bringing down current application and services"
docker-compose down &&
echo "bringing up application and services"
docker-compose up -d --no-cache`
	return []byte(src)
}
