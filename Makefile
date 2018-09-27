frontend-install:
	cd frontend && npm i

frontend-run:
	cd frontend && npm start

frontend-build:
	cd frontend && npm run build

backend-run:
	cd backend && go run main.go

backend-build:
	cd backend && go build main.go

build-all: frontend-install frontend-build backend-build
	rm -rf build
	mkdir build
	mv backend/main build
	mv frontend/dist build
	mv build/dist/static/index.html build/dist
