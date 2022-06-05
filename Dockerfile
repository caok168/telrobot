#FROM harbor.aibee.cn/online-pipeline/golang:latest as builder
FROM golang:1.17 as builder

WORKDIR /root/telrobot
RUN go env -w GOPROXY=https://goproxy.cn,direct
ADD . .
RUN make

#FROM registry.aibee.cn/parkinglot/alpine_plus:latest
FROM harbor.aibee.cn/online-pipeline/alpine_plus:3.7

WORKDIR /root
COPY --from=builder /root/telrobot/bin/TelRobot /root/TelRobot
RUN ln -s /root/TelRobot /bin/TelRobot

CMD ["TelRobot"]