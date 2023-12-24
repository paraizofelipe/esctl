FROM golang:alpine as base
RUN apk add --no-cache make cmake

FROM base as build
COPY . /opt/esctl
WORKDIR /opt/esctl
RUN make deps && make build 

FROM scratch
COPY --from=build /opt/esctl /opt/paraizo/
ENTRYPOINT ["/opt/paraizo/esctl"]
