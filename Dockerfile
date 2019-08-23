FROM scratch
ADD build/tracking /tracking
EXPOSE 8080
ENTRYPOINT ["./tracking"]