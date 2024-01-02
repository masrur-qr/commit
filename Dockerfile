FROM golang
COPY . /new
WORKDIR /perent
COPY go.mod /perent/
RUN go mod download 
COPY . /perent/
EXPOSE 2121
RUN go build -o main ./
CMD [ "/main" ]