FROM gobuffalo/buffalo:v0.16.21 as builder

ENV GO111MODULE on
ENV GOPROXY http://proxy.golang.org

RUN mkdir -p /src/rockgate
WORKDIR /src/rockgate

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

ADD . .
RUN go build -o /bin/rockgate

FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

COPY --from=builder /bin/rockgate .

# Uncomment to run the binary in "production" mode:
# ENV GO_ENV=production

# Bind the app to 0.0.0.0 so it can be seen from outside the container
ENV ADDR=0.0.0.0

EXPOSE 8181

CMD ["rockgate"]