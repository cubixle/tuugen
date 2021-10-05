# Tuugen

Tuugen takes a tuugen.yml and a grpc proto definition file to generate common boiler plate logic. Some of the common boiler plate it generates is http routes, grpc service, basic data model structs and a main file with some default plumbing.

Tuugen adds an interactor package which the http and grpc transport layers use for business logic.

Feel free to pull the repo and play around with the example_project.

## Using Tuugen.

```
go install github.com/cubixle/tuugen
```

And the run `tuugen` in the root dir of your project or wherever you have your tuugen and proto file.

## Example tuugen file
```yaml
project: tuugen_test_project
import_path: github.com/cubixle/tuugen/example_project
proto_file: service.proto
grpc_file: internal/pb/service/service_grpc.pb.go
service_name: Service
data_models:
  - name: User
    properties:
      - name: id
        type: varchar
        autoinc: true
      - name: team_id
        type: varchar
      - name: name
        type: varchar
      - name: email
        type: varchar
      - name: created_at
        type: timestamp
      - name: updated_at
        type: timestamp
  - name: Team
    properties:
      - name: id
        type: varchar
      - name: name
        type: varchar
```

## TODO

- [ ] http routes
- [ ] improve data model to struct. add more types.
- [ ] Add golangci-lint
- [ ] Add a useful Makefile
- [ ] Add a useful Dockerfile
- [ ] Add a readme template