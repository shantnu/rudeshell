# The base go-image
FROM golang:1.14


RUN go get github.com/yudai/gotty

# Create a directory for the rude
RUN mkdir /rude

# Copy all files from the current directory to the rude directory
COPY . /rude

# Set working directory
WORKDIR /rude

RUN go build -o rude .

# Run the server executable
ENTRYPOINT [ "./run.sh" ]