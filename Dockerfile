# Build stage
FROM golang:1.16-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server


# Run stage
FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/server /app/server

CMD ["/app/server"]