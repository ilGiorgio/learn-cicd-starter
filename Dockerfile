FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

ADD notely /usr/bin/notely

CMD ["notely"]
