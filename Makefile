frontend-install:
	cd frontend && npm i

frontend-run:
	cd frontend && npm start

frontend-build:
	cd frontend && npm run build

backend-run:
	cd backend && go run main.go

backend-local-build:
	cd backend && go build main.go

backend-cross-build:
	cd backend && GOOS=linux  GOARCH=amd64 CGO_ENABLED=0 go build main.go

build-cross: frontend-install frontend-build backend-cross-build
	rm -rf build
	mkdir build
	mv backend/main build
	mv frontend/dist build
	mv build/dist/static/index.html build/dist
