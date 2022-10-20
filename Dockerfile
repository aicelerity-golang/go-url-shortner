FROM golang:1.19
WORKDIR /app
ADD . /app/

# Below RUN commands are added in Dev, it's not needed in Prod
RUN go mod tidy
RUN go mod verify

EXPOSE 5000

CMD ["go", "run", "main.go"]