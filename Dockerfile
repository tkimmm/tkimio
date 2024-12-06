FROM golang:1.22.1 as base

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  dockeruser

WORKDIR /app

# COPY go.mod go.sum ./
COPY . ./
RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /tkimio .

FROM gcr.io/distroless/static-debian12

COPY --from=base /etc/group /etc/group
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /app/static/ /static
COPY --from=base /tkimio .

USER dockeruser:dockeruser

CMD ["./tkimio"]
