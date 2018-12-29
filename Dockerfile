FROM golang:1.11-stretch AS builder
WORKDIR /go/src/github.com/pollosp/gin-test/
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
COPY . .
RUN dep ensure -v
RUN useradd -u 10001 scratchuser
RUN go test && CGO_ENABLED=0 go build

FROM scratch
WORKDIR /app/
COPY --from=builder /go/src/github.com/pollosp/gin-test/gin-test .
COPY --from=builder /etc/passwd /etc/passwd
USER scratchuser
ENTRYPOINT ["./gin-test"]
CMD [ ]
