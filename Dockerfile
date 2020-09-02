FROM lixueli/lxl-base:1.0.0
COPY . code
RUN cd code && make build-cross
RUN cd code && cp -r build /starter
RUN ls /starter
WORKDIR /starter
EXPOSE 5678
ENTRYPOINT ["/starter/main"]
