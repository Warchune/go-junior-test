FROM golang:1.21

ADD ./bin/goose /goose
RUN chmod +x /goose

ADD ./bin/app /app

CMD /goose up && /app
