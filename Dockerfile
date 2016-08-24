# Use alpine linux which is a
# tiny linux image
FROM alpine

# Copy amigo binary to
# guest machine
COPY ./amigo .

EXPOSE 8080
ENTRYPOINT ./amigo
