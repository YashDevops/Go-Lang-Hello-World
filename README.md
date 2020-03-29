# Go-Lang-Hello-World

A simple GO application 

### The Source Code for the following Application : [Code Link ](https://github.com/YashDevops/Go-Lang-Hello-World)

#### Problem Statement

### Create a Golang based program which would create a hello-world webserver on port 8080 and have the following api's

### 1) GET /hello responds with "hello world". Provide sample curl call for input and output.
### 2) POST /sum with data consisting of a list of numbers responds with the sum of the numbers. Provide sample curl call for input and output.
### 3) POST /loglevel with data 'loglevel:debug' changes the logging level from info to debug, i.e, outputs more verbose/debug level logging. Provide sample curl call for input and output.


## How to use this image

Note: Pulling an image requires using  latest tag.


```sh
docker run -itd --name grab-assignment -p 8080:8080 yashshahdevops/assignment:latest 
```

#### Attention 
###### If you face this error while running the docker container

```sh
docker: Error response from daemon: driver failed programming external connectivity on endpoint gracious_lehmann (c4c441fe06e65ec7d0907c3b4c7fa309f6faba0aa012bcaacee1d7ab8ce50eb2): Error starting userland proxy: listen tcp 0.0.0.0:8080: bind: address already in use.
```

##### change the host port to another other port as some other service may be listening to that port 8080

## Hello API

#### Sample Curl 
```sh
curl --location --request GET 'localhost:8080/hello'
```

#### Sample Output


```sh
Hello World
```


## Sum API

#### Sample Curl

```sh
curl --location --request POST 'localhost:8080/sum' --header 'Content-Type: application/json' --data-raw '{"Number": [3,6,10,18,27]}'
```

#### Sample Output

```sh
{"number":[3,6,10,18,27],"sum":64}

```


##### NOTE :- The following build has been built automatically using Github.

#### Build Screenshot
![Image](https://raw.githubusercontent.com/YashDevops/Go-Lang-Hello-World/master/build.png)



