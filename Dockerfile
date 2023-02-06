FROM golang:1.19.2-alpine AS builder
ARG GO_VERSION
ARG CGO
ARG OS
ARG ARCH

RUN apk --no-cache add tzdata
RUN mkdir /go/src/lms
WORKDIR /go/src/lms
COPY . ./

RUN go mod tidy
RUN go mod vendor



RUN CGO_ENABLED=$CGO  GOOS=$OS GOARCH=$ARCH go build -mod vendor -a -installsuffix cgo -ldflags '-extldflags "-static"' -o ./bin/lms ./main.go
RUN cp ./bin/lms /bin/


FROM scratch as production
ENV GIN_MODE=debug
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/src/lms/.env /lms/
COPY --from=builder /go/src/lms/bin/lms /lms/lms
ENV TZ Asia/Kolkata
EXPOSE 3000
WORKDIR /lms
ENTRYPOINT ["/lms/lms"]