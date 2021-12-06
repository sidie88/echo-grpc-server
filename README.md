## Stockbit Movie API

### Tech Stack

 - Go v1.17
 - LabStack Echo
 - GRpc
 
### Server Config
Default config
```
ECHO_PORT=8080
GRPC_ADDRESS=localhost
GRPC_PORT=8888
OMDB_APIKEY=faf7e5bb
```
You can change this config above by set it on environment variable

### API Endpoint

- Rest API
```
/search
/movie-detail
```
- GRpc
```
grpc.Search/searchWithPagination
grpc.Search/getMovieDetail
```