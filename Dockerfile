FROM golang:latest as builder

ENV SERVICE_NAME=api-base
ENV APP /src/${SERVICE_NAME}/
ENV WORKDIR ${GOPATH}${APP}

WORKDIR $WORKDIR
ADD . $WORKDIR

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $SERVICE_NAME .

FROM alpine

ENV SERVICE_NAME=api-base
ENV APP /src/${SERVICE_NAME}/
ENV GOPATH /go
ENV WORKDIR ${GOPATH}${APP}

COPY --from=builder ${WORKDIR}${SERVICE_NAME} $WORKDIR

WORKDIR $WORKDIR
CMD ./${SERVICE_NAME}
