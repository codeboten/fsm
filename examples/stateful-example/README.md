# stateful example
This example persists tracing information using Consul.

## pre-requisites
Consul
### OSX
```bash
brew install consul
consul agent -dev
```

## build
```bash
make
```

## run
```bash
./watcher
```