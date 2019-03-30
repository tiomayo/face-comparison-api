[![N|Solid](https://gitlab.docotel.net/uploads/-/system/appearance/header_logo/1/Docotel_Logo_28_px__2_.png)](https://docotel.com) [Tiomayo](https://github.com/tiomayo/face-comparison-api/) ,
[Kaonashi](https://github.com/adhityasan/face-comparison-api/)


# face-comparison-api

Restful API to process images
 - Get FaceId from Azure
 - Get OCR scan result from Azure `(on work)`
 - Face mathing result from azure
 - Get OCR scan result from AWS `(on work)`
 - Face mathing result from AWS `(on work)`
 - Get OCR scan result from Google Vision `(on work)`
 - Face mathing result from Google Vision `(on work)`

## Running the app locally

```sh
$ go build
$ ./face-comparison-api
2019/02/03 11:38:11 Starting Server
``` 
- [route test](http://localhost:8000/go/aisatsu?name=Guest)

```sh
$ curl http://localhost:8000/go/aisatsu?name=Kaonashi
Hello, Kaonashi
```

## Building and running the docker image

```sh
$ docker build -f "Dockerfile" -t face-comparison-api:1.0.0 .
$ docker run -d -p 8000:8000 face-comparison-api:1.0.0
2019/02/03 11:38:11 Starting Server at :8000...
```

If you want to build the apps with log, use the Dockerfile.volume file to build
 - Notice how it mount a directory of the Host OS to the volume specified by the docker container -

```sh
$ docker build --rm -f "Dockerfile.volume" -t face-comparison-api:1.0.0 .
$ mkdir ~/app-logs
$ docker run -d -p 8000:8000 -v ~/app-logs:/face-comparison-api/logs face-comparison-api:1.0.0
2019/02/03 11:38:11 Starting Server at :8000...
```

If you want to optimize / reduce the size of the docker image, you could use the Dockerfile.multistage file, It use a very lightweight [Alpine linux](https://alpinelinux.org) image and will only contain the binary executable built by the first stage. DOcker image size should be about 20MB-30MB 

```sh
$ docker build --rm -f "Dockerfile.multistage" -t face-comparison-api:1.0.0 .
```

Read the tutorial: [Building Docker Containers for Go Applications](https://www.callicoder.com/docker-golang-image-container-example/) 


## Contributing
Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

#### License
[MIT](https://choosealicense.com/licenses/mit/)