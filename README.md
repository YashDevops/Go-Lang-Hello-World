# Go-Lang-Hello-World
A simple go application for hello-world.

#### Problem Statement

### 1 ### Create a Golang based program which would create a hello-world webserver on port 8080 such that
### -) GET /hello responds with "hello world". Provide sample curl call for input and output.
### -) POST /sum with data consisting of a list of numbers  responds with the sum of the numbers. Provide sample curl call for input and output.
### -) POST /loglevel with data 'loglevel:debug' changes the logging level from info to debug, i.e, outputs more verbose/debug level logging. Provide sample curl call for input and output.


## Hello Api

#### Sample Curl 
```
curl --location --request GET 'localhost:8080/hello
```
#### Sample Output

```
Hello World
```


## Sum Api

#### Sample Curl

```
curl --location --request POST 'localhost:8080/sum' --header 'Content-Type: application/json' --data-raw '{"Number": [3,6,10,18,27]}'
```

#### Sample Output

```
{"number":[3,6,10,18,27],"sum":64}

```



