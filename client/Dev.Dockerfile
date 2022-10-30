FROM denoland/deno:1.25.0

ARG GIT_REVISION
ENV PORT=${PORT}
ENV DENO_DEPLOYMENT_ID=${GIT_REVISION}

WORKDIR /app

COPY . .
RUN deno cache main.ts --import-map=import_map.json

EXPOSE ${PORT}

CMD ["run", "-A", "--watch=static/,routes/,islands/,components/", "main.ts"]
