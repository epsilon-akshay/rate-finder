FROM golang:latest
RUN mkdir /currency-converter
WORKDIR  /currency-converter
COPY . .
RUN make build
EXPOSE 8000
CMD ["./out/currency-converter"]