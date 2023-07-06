FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
COPY static/ ./static/
COPY *.go ./
COPY *.html ./

# Build
RUN go build -o /kube-ingress-dashboard

# Stage 2: Scratch image to minimise size
FROM scratch
WORKDIR /bin
COPY --from=builder /kube-ingress-dashboard .
COPY *.html ./

EXPOSE 8080
CMD ["./kube-ingress-dashboard"]