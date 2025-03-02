FROM node:22-bookworm

WORKDIR /app

COPY . .

RUN rm /bin/sh && ln -s /bin/bash /bin/sh

RUN npm install -g pnpm \
    && SHELL=bash pnpm setup \
    && source /root/.bashrc \
    && pnpm install turbo --global

COPY ./package.json ./pnpm-lock.yaml /app/

RUN pnpm i
# RUN npx playwright install
# RUN npx playwright install-deps

CMD sh -c "pnpm run dev"
