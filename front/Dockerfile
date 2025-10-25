FROM oven/bun:latest as build

WORKDIR /app

COPY package.json bun.lock tsconfig.json ./
RUN bun install

COPY bun-env.d.ts bunfig.toml  ./
COPY src src

RUN bun run build

FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html


