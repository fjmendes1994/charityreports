FROM golang:latest
RUN go get github.com/fjmendes1994/charityreports
WORKDIR src/github.com/fjmendes1994/charityreports
RUN go build
CMD ["charityreportsgt"]
