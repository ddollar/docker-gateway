FROM scratch
ADD build/docker-gateway /docker-gateway
CMD ["/docker-gateway"]
