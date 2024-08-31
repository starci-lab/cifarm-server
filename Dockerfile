FROM node:alpine AS node-builder

WORKDIR /backend
COPY typescript-runtime .

RUN npm install
RUN npx tsc

FROM heroiclabs/nakama-pluginbuilder:3.23.0 AS builder

ENV GO111MODULE on
ENV CGO_ENABLED 1

WORKDIR /backend
COPY go-runtime .
COPY local.yml .

RUN go build --trimpath --buildmode=plugin -o backend.so

FROM heroiclabs/nakama:3.23.0

COPY --from=node-builder /backend/build/*.js /nakama/data/modules/build/
COPY --from=builder /backend/backend.so /nakama/data/modules
COPY --from=builder /backend/local.yml /nakama/data/
COPY --from=builder /backend/*.json /nakama/data/modules