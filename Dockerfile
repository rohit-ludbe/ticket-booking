FROM golang:1.21.5

WORKDIR /usr/src/app

# add hot reload
RUN go install github.com/cosmtrek/air@v1.48.0

# copy all files from host to container 
COPY . .
# package properly install
RUN go mod tidy

