FROM golang:1.25 as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

# building go with external linking for support of cgo
# building static linked executable
RUN go build -ldflags "-linkmode 'external' -extldflags '-static'"  -o main ./cmd/api/main.go

# magic! we got 30mb image
FROM scratch AS release

WORKDIR /app

COPY --from=build /app/main /main

CMD [ "/main" ]
