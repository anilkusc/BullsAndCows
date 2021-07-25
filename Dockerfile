FROM golang:1.16 as build_backend
RUN apt-get update && apt-get install sqlite3 -y && mkdir /db 
WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN rm -fr ./frontend && rm -fr ./frontend && rm -fr ./ui-test
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o /bin/app .

FROM golang:1.16 as build_frontend
WORKDIR /src/frontend
COPY ./frontend .
RUN ls -al && cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
RUN go mod download
RUN GOOS=js GOARCH=wasm go build -o main.wasm main.go 
RUN GOOS=js GOARCH=wasm go build -o game.wasm game.go
RUN rm main.go && rm game.go

FROM alpine
RUN apk add python3 bash
WORKDIR /application
COPY --from=build_frontend /src/frontend /application
COPY --from=build_backend /bin/app /application/app
COPY ./entrypoint.sh ./entrypoint.sh
RUN chmod +x ./app
RUN chmod +x entrypoint.sh
ENTRYPOINT ["/bin/bash"]
CMD ["entrypoint.sh"]
