FROM golang:1.19.2-alpine3.15 as builder

WORKDIR /build
COPY go.mod ./
RUN go mod download && go mod tidy
COPY . .
RUN go build -o main .


# Input parameters for the Dockerfile expected in os.Getenv




FROM golang:1.19.2-alpine3.15
# Add maintainer info
LABEL maintainer="Danail Surudzhiyski"
ARG dbhost
ARG dbuser
ARG dbpass
ENV DATABASE_HOST=$dbhost
ENV DATABASE_USER=$dbuser
ENV DATABASE_PASS=$dbpass

COPY --from=builder /build/main .
COPY --from=builder /build/static ./static

EXPOSE 8080

CMD ["./main"]