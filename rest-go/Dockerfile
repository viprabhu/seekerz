FROM golang
RUN mkdir /rest-go
ADD main.go /app/
WORKDIR /app
RUN go build -o rest-go .
EXPOSE 8080
CMD ["/app/rest-go"]