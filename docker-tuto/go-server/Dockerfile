FROM golang:alpine

# Set a working directory 
WORKDIR /go/src/github.com/Remydeme/server

# copy our directory file into /go/src/github.com/Remydeme/server 
COPY . .

# Install curl and git 
RUN apk update && apk add curl && apk add git 

# Install go dep to be able to download all the dependencies 
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh 

# isntall server dependencies 
RUN dep init 

# expose port 8000
EXPOSE 8000

# This command will run the server when we will execute docker run command 
CMD ["go", "run", "main.go"]
