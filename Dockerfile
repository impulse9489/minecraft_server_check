FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
ADD check_server /
CMD ["/check_server"]
