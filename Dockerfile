FROM arm64v8/golang:1.20-alpine AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY static/ ./static/
COPY *.go ./

# Build
RUN go build -o /kube-ingress-dashboard

EXPOSE 8080
CMD ["./kube-ingress-dashboard"]

# Stage 2: Scratch image
FROM scratch
WORKDIR /bin
COPY --from=builder /kube-ingress-dashboard .

EXPOSE 8080
CMD ["./kube-ingress-dashboard"]