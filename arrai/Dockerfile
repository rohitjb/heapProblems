ARG INSTALLER=golang:1.13.4
ARG TARGET=alpine:3.10

FROM ${INSTALLER} as installer
ARG ARRAI_VERSION
ENV GO111MODULE=on
RUN go get -d github.com/arr-ai/arrai/...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /arrai github.com/arr-ai/arrai/cmd/arrai

FROM ${TARGET}
COPY --from=installer /arrai /usr/bin