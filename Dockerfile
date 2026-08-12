FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install

ADD notely /usr/bin/notely

CMD ["notely"]
