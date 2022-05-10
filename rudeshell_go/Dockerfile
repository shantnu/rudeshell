FROM golang:1.14 as builder

RUN go get github.com/yudai/gotty

# Create a directory for the rude
RUN mkdir /rude

# Copy all files from the current directory to the rude directory
COPY . /rude

WORKDIR /rude
RUN go build -o rude .


from debian
COPY --from=builder /rude/rude /rude/rude
COPY --from=builder /rude/run.sh /rude/run.sh
COPY --from=builder /go/bin/gotty /rude/gotty
RUN chmod +x /rude/gotty
# Set working directory
WORKDIR /rude

# Run the server executable
ENTRYPOINT [ "./run.sh" ]