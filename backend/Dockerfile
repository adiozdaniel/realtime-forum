FROM golang:bullseye AS build1
# ENV CGO_ENABLED=1 GCC_GO="gccgo" CC="gcc"
WORKDIR /backend
COPY . .
RUN go mod download
RUN go build -o ./forum-dockerize .
# FROM scratch
# WORKDIR /
# COPY --from=build1 /backend/forum-dockerize ./forum-dockerize
EXPOSE 4000
#USER nonroot:nonroot
# CMD ./main
CMD [ "./forum-dockerize" ]
# ENTRYPOINT [ "/forum-dockerize" ]