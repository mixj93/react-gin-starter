FROM scratch
COPY build/ starter/
EXPOSE 5678
ENTRYPOINT ["/starter/main"]