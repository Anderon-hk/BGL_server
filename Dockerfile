From golang
LABEL maintainer="Tam Chi Fung chiftam2-c@my.cityu.edu.hk"
RUN go get -u github.com/go-chi/chi
COPY ./main.go /
CMD ["go", "run", "/main.go"]