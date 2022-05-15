package components

func BuildScript(tag string) []byte {

	src := `
echo "building IDP UI"
cd idp &&
docker build --tag idp-ui:` + tag + ` . --no-cache &&
cd .. &&
echo "building auth service"
cd token-auth-svc &&
go build -ldflags "-s -w"  &&
docker build --tag auth-svc:` + tag + `  . --no-cache &&
cd .. &&
docker system prune --force &&
docker pull redis`
	return []byte(src)
}
