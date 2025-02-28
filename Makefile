tidy:
	go mod tidy

full_run:
	swag init -g cmd/main.go ; yes | swag fmt ; go run cmd/main.go

run:
	go run cmd2/main.go

swagger:
	swag init -g cmd/main.go

install_deps:
	go get github.com/gofiber/fiber/v2
	go get -u gorm.io/gorm

watch:
	find . -name "*.go" | entr -r make run

# gi is stands for `gitmoji`
gi:
	nodeactivate ; git add . ; gitmoji -c
