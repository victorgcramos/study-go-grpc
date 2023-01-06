# GO GRPC STUDY

> This repo is used for studying purposes.

I've been doing a lot of research these days regarding the Go
GRPC microservices, and this repo describes a basic and very
minimalistic project that has an API `gateway`, that is used to
make calls with the `user` service client.

### Run
1. Generate the Proto Buffer files `(*.pb.go)` 
   ```sh 
   $ make proto
   ```
2. Run user GRPC service
   ```shell
   $ make user-service
   ```
3. Run gateway server
   ```shell
   $ make gateway
   ```

### Gateway HTTP Routes
> WIP: I'll add new routes further.

| **endpoint** | **method** | **params**                               |
|--------------|------------|------------------------------------------|
| `/version`   | GET        | -                                        |
| `/user/new`  | POST       | `{ username: string, password: string }` |
