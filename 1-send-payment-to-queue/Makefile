clean:
	rm -rf build
	
build:
	cd src && GOOS=linux GOARCH=amd64 go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o ../build/bin/bootstrap

init:
	terraform -chdir=infra init

plan:
	terraform -chdir=infra plan

apply:
	terraform -chdir=infra apply --auto-approve

destroy:
	terraform -chdir=infra destroy --auto-approve