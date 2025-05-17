FROM golang:1.23.4-bullseye AS build

ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN go build -o svc ./cmd/vuz-mobapp-backend/.

FROM debian

WORKDIR /application

COPY --from=build /build/svc /application/svc

RUN ls -la

CMD ["./svc"]