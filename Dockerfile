FROM golang:1.14.4 as gobuilder

WORKDIR /work
COPY . .
RUN CGO_ENABLED=0 go build

FROM node:14.4.0 as nodebuilder

WORKDIR /work
COPY view/package.json view/package-lock.json ./
RUN npm ci
COPY view .
RUN npm run build

FROM alpine:3.12.0

WORKDIR /app
COPY --from=gobuilder /work/git-bucket-to-lab ./
COPY --from=nodebuilder /work/dist ./view

EXPOSE 1323
RUN ls
RUN ls /app -la

CMD ["/app/git-bucket-to-lab"]
