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