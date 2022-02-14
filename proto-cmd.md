```
go get github.com/micro/go-web
go get github.com/micro/go-micro

go get github.com/micro/protobuf/{proto,protoc-gen-go}
go get github.com/micro/protoc-gen-micro
```

```
// ubuntu success
protoc -I=/home/vesper/go/src/freelancer-go/service/account/proto/ --micro_out=/home/vesper/go/src/freelancer-go/service/account/proto/ --go_out=/home/vesper/go/src/freelancer-go/service/account/proto/ /home/vesper/go/src/freelancer-go/service/account/proto/user.proto --go_opt=/home/vesper/go/src/freelancer-go/service/account/proto/user.proto
```