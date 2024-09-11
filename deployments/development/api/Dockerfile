FROM golang:1.23.0-alpine3.20
WORKDIR /plank
COPY . /plank
RUN apk add --update --no-cache \
      git && \
      go install github.com/cortesi/modd/cmd/modd@latest && \
      go install github.com/volatiletech/sqlboiler@latest && \
      go install github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql@latest
CMD ["modd", "-f", "configs/modd.conf"]
