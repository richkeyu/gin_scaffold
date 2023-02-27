FROM golang:1.17

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

ENV APP_HOME /app
WORKDIR $APP_HOME
COPY . .

CMD sh startup.sh
# ENTRYPOINT ["startup.sh"]
