FROM golang:alpine AS builder 

COPY . /app

RUN go build -o /tmp/ctf cmd/ctf/main.go


# ---------------------------------------------
FROM scratch AS main 

COPY --from=builder /tmp/ctf /ctf 


CMD ["/ctf"]