FROM golang:1.13
ENV GO111MODULE=on
ENV TEMPDBPASS root
WORKDIR /sensorapi/
COPY . .
RUN ls
RUN GOOS=linux GOARCH=amd64 go build ./main.go
RUN mv ./main ./sapi
CMD ./sapi -url 0.0.0.0:8080 -dbname "root:${TEMPDBPASS}@tcp(db:3306)/tempanalizr" -rabbit "amqp://guest:guest@rabbit:5672"
