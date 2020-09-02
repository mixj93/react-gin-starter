FROM lixueli/lxl-base:1.0.0
COPY . code
RUN cd code && make build-cross
COPY build /starter
RUN ls /starter
EXPOSE 5678
ENTRYPOINT ["/starter/main"]
