FROM golang:1.21.0

WORKDIR /app

COPY . .

RUN go build -o telegram-bot

# EXPOSE 8080 # currently our bot aren't using ports to be exposed

CMD ["./telegram-bot"]
