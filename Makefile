tidy:
	go mod tidy

start-db:
	docker start postgres-16.3-temp

stop-db:
	docker stop postgres-16.3-temp

full_run:
	swag init -g cmd/main.go ; yes | swag fmt ; go run cmd/main.go

run:
	go run cmd4/main.go

swagger:
	swag init -g cmd/main.go

install_deps:
	go get github.com/gofiber/fiber/v2
	go get -u gorm.io/gorm

watch:
	find . -name "*.go" | entr -rz make run

# gi is stands for `gitmoji`
gi:
	nodeactivate ; git add . ; gitmoji -c
