frontend-install:
	cd frontend && yarn install

frontend-run:
	cd frontend && yarn start

frontend-build:
	cd frontend && yarn run build

backend-run:
	cd backend && go run cmd/main.go --dev --debug

backend-local-build:
	cd backend && go build cmd/main.go

backend-cross-build:
	cd backend && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build cmd/main.go

build-local: frontend-install frontend-build backend-local-build
	rm -rf build
	mkdir build
	mv backend/main build
	mv frontend/build build/dist

build-cross: frontend-install frontend-build backend-cross-build
	rm -rf build
	mkdir build
	mv backend/main build
	mv frontend/build build/dist