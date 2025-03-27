FROM golang:1.28-alpine as dependencies

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM dependencies as build
COPY . ./
RUN CGO_ENABLE=0 go build -o /main -ldflags="-w -s" .

FROM golang:1.28-alpine
COPY --from=build /main /main
CMD ["/main"]