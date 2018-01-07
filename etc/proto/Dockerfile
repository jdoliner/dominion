FROM python
LABEL maintainer="jdoliner@gmail.com"

RUN apt-get update -yq
RUN apt-get install -yq unzip

RUN wget https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip -O protoc.zip
RUN unzip protoc.zip -d /
RUN cp -r /include /bin

ADD run /
ENTRYPOINT ["/run"]
WORKDIR /work
