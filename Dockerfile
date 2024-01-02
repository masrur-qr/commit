FROM golang
COPY . /new
WORKDIR /perent
COPY go.mod /perent/
RUN go mod download 
COPY . /perent/
EXPOSE 2121
RUN go build -o main ./
CMD [ "./main" ]
# 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