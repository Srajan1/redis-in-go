# WHAT?
Redis server implementation in go.

## Sources:
Has been inspired from this article:  https://www.build-redis-from-scratch.dev/en/introduction
![alt text](image.png)
<small>This repo is 100% organic, home grown and hand written, no LLMs were involved in writing it.</small>

## Setup
This repo implements redis' server, and not the client. You can use original Redis client to interact with the server.
To install the client on macos run 

```bash
brew install redis
```
Also stop the redis server if it's running

```bash
brew services stop redis
```

## Basics
Redis uses RESP (Redis Serializaion Protocol) to communucate between client and server. The SET command sends the key as a string and the value could be int, bool, string, etc. When you do 
```bash
SET admin Dalinar
```
this gets converted to 
```
*3\r\n$3\r\nSET\r\n$5\r\nadmin\r\n$7\r\nDalinar
```
by RESP. Or we can simply visualize like below.
```golang
*3         # Indicates that it is a array of size 3
$3         # Indicates that next token is of length 3
SET        # Token itself
$5         # Indicates that next token is of length 5
admin      # Token itself
$7         # Indicates that next token is of length 7
Dalinar    # Token itself
```
This sets the admin as [Dalinar](https://stormlightarchive.fandom.com/wiki/Dalinar_Kholin)

Similarly when we do `GET ADMIN`. Server returns `$7\r\nDalinar\r\n`

