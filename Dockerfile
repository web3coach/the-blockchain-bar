FROM golang:1.16

WORKDIR /build
ADD ./ /build

RUN apt install bash

RUN make install

ENTRYPOINT ["tbb"]
CMD ["-h"]
