FROM denoland/deno:1.27.0

ARG PORT
ARG GIT_REVISION
ENV DENO_DEPLOYMENT_ID=${GIT_REVISION}

WORKDIR /app

COPY . .
RUN deno cache main.ts --import-map=import_map.json

EXPOSE ${PORT}

CMD ["run", "-A", "main.ts"]
