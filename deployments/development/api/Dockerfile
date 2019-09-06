FROM golang:1.13.0-alpine3.10
WORKDIR /plank
COPY . /plank
RUN apk add --update --no-cache \
      git && \
      go get github.com/cortesi/modd/cmd/modd && \
      go get -u -t github.com/volatiletech/sqlboiler && \
      go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
CMD ["modd", "-f", "configs/modd.conf"]
