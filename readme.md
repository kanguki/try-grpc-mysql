## Dev note:

- To exclude `XXX_` fields in message (cause errors when automigrating), go get github.com/gogo/protobuf/protoc-gen-gogofasterinstall to use --gogofaster_out:
  protoc --proto_path=proto proto/*.proto --gogofaster_out=plugins=grpc:internal/core/movie
-

## Call grpc enpoints: (assume port = 2381)

- grpcurl  -plaintext localhost:2381 list
- grpcurl -d '{"movie":{"title":"run bts", "director":"mo", "thumbnail":"link", "status":1, "country":"korea"}}' -plaintext localhost:2381 movie.MovieService/AddMovie
- grpcurl -d '{"title":"bts","director":"mo"}' -plaintext localhost:2381 movie.MovieService/SearchMovie
- grpcurl -d '{"movie":{"id":4,"title":"Run BTS"}}' -plaintext localhost:2381 movie.MovieService/UpdateMovie
- grpcurl -plaintext localhost:2381 movie.MovieService/GetAllMovies
- grpcurl -d '{"id":[4]}' -plaintext localhost:2381 movie.MovieService/DeleteMovie
