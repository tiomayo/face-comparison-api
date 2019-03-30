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
$ ./go-docker
2019/02/03 11:38:11 Starting Server
``` 
- [route test](http://localhost:8000/go/aisatsu?name=Guest)

```sh
$ curl http://localhost:8000/go/aisatsu?name=Kaonashi
Hello, Kaonashi
```

## Building and running the docker image

```sh
$ docker build --rm -f "Dockerfile" -t face-comparison-api:1.0.0 .
$ docker run -d -p 8000:8000 face-comparison-api:1.0.0
2019/02/03 11:38:11 Starting Server at :8000...
```

Read the tutorial: [Building Docker Containers for Go Applications](https://www.callicoder.com/docker-golang-image-container-example/) 


## Contributing
Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

#### License
[MIT](https://choosealicense.com/licenses/mit/)