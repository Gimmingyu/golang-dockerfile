FROM golang:1.20 as build

WORKDIR /home/go/src

RUN ls -la

COPY . /home/go/src/primrose

FROM build as release

WORKDIR /home/go/src/primrose

EXPOSE 8080

RUN ls -la

RUN go build -o maind main.go

ENTRYPOINT ["maind"]
