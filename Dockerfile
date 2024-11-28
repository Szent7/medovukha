FROM node:alpine AS frontend-builder
WORKDIR /build
ADD ./frontend/package.json .
ADD ./frontend/package-lock.json .
RUN npm install
COPY ./frontend .
RUN npm run build

FROM golang:alpine AS backend-builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o main main.go 

FROM alpine
WORKDIR /build/dist
COPY --from=frontend-builder /build/dist /build/dist/
WORKDIR /build
COPY --from=backend-builder /build/main /build/main

EXPOSE 10015

CMD ["./main"]