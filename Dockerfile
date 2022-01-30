FROM golang:latest
RUN mkdir /currency-converter
WORKDIR  /currency-converter
COPY . .
EXPOSE 8000
RUN ls
