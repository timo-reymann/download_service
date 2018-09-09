FROM golang:latest
EXPOSE 8086
RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY ./src/form.html ./form.html
RUN go build -o server ./src
CMD ["/app/server"]
