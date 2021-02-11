From golang

RUN go get -u github.com/go-chi/chi
COPY ./main.go /
CMD ["go", "run", "/main.go"]