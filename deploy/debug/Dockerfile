FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

ARG APP_RELATIVE_PATH

COPY . /src
WORKDIR /src/app/${APP_RELATIVE_PATH}

EXPOSE 8000
EXPOSE 9000

CMD ["bin/server"]