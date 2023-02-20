FROM golang:1.20.0 AS builder
WORKDIR /go/src/github.com/minmax1996/lolesports-calendar/
COPY app app
COPY internal internal
COPY main.go .
COPY go.mod .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/calendar .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
COPY --from=builder  /go/bin/calendar /bin/
ENTRYPOINT ["calendar"]