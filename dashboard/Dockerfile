FROM mhart/alpine-node:0.10
MAINTAINER x0rzkov <x0rzkov@protonmail.com>

RUN apk add --no-cache --no-progress make ca-certificates git nano bash jq python

WORKDIR /opt/service
COPY package.json .

RUN apk add --no-cache --no-progress g++ gcc musl-dev
RUN npm install

COPY . .

RUN make all

EXPOSE 8000

CMD ["make", "server"]
