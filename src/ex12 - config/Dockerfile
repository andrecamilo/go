FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get "github.com/crgimenes/goconfig"
ENV DOMAIN=DOMAIN_XXX
ENV MONGODB_HOST=MONGODB_XXX
RUN go build -o main . 
CMD ["/app/main"]







# FROM golang:1.12

# WORKDIR ${GOPATH}/config
# COPY . ${GOPATH}/config

# RUN echo ${GOPATH}/config

# RUN chown -R ${APP_USER}:${APP_USER} ${GOPATH}
# USER ${APP_USER}

# ENV DOMAIN=DOMAIN_XXX
# ENV MONGODB_HOST=MONGODB_XXX
# ENTRYPOINT [${GOPATH}"/config"]