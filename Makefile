download:
	go mod download
compose:
	docker-compose -f docker-compose.yml up -d
install:
	go install
setup: compose install