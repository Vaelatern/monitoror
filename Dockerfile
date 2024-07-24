FROM alpine:latest AS builder
ARG VERSION
WORKDIR /building
RUN apk add yarn
COPY . .
WORKDIR /building/ui
RUN node install
RUN yarn build
WORKDIR /building
RUN go build cmd/monitoror

FROM alpine:latest
RUN apk update && \
    apk --no-cache add ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*
COPY --from=builder /bin/monitoror /bin/monitoror
EXPOSE 8080
CMD [ "/bin/monitoror" ]
