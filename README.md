# The Go App Boot Framework

`good` is a http framework that makes developers write go applications much easier.

## Download and Install
```bash
go get -u github.com/gocurr/good
```

## Usage

```go
import "github.com/gocurr/good"
```

```go
var nameFns = good.NameFns{
	{"demo1", func() {
		fmt.Println("demo1...")
	}},

	{"demo2", func() {
		fmt.Println("demo2...")
	}},
}
good.Configure("./app.yml", false)
if err := good.StartCrontab(nameFns); err != nil {
    panic(err)
}

good.ServerMux(http.NewServeMux())
good.Route("/", func(w http.ResponseWriter, r *http.Request) {
    _, _ = w.Write([]byte("ok"))
})
good.Fire()
```