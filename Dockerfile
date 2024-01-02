FROM golang
COPY . /new
WORKDIR /perent
COPY . /perent/
RUN go build -o main ./
CMD [ "/main" ]