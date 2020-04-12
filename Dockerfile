# Server Build Stage
FROM golang:latest AS builder

# Copy app into docker container
COPY . /app
WORKDIR /app/cmd/go-fishbowl
RUN go mod download

# Build server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .

# Frontend Build Stage
FROM node:alpine AS node_builder
COPY --from=builder /app/frontend ./
RUN npm install
RUN npm run build

# Final Build, what is actually deployed
FROM alpine:latest
COPY --from=builder /main ./
COPY --from=node_builder /build ./web
RUN chmod +x ./main

EXPOSE 8080
CMD ["./main"]

# RUN chmod +x ./bin/go-fishbowl

# CMD ["./bin/go-fishbowl"]
