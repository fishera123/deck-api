FROM golang:1.14.3-alpine3.11 AS builder

ENV HOME /home

RUN apk add --no-cache gcc libc-dev

WORKDIR $HOME

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go mod tidy && go build -a .


FROM alpine:3.11.6 AS final

RUN apk add --no-cache bash

ENV HOME /home

WORKDIR $HOME

COPY --from=builder $HOME/deck-api .

CMD [ "./deck-api" ]
