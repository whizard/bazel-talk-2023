# Some Useful Stuff For Using BazelAs
Here are some basic and useful Bazel commands.

## Bazel build

The build command builds your packages.

### Build all packages/targets

```
bazel build //...
```

### Build a specific package

```
bazel build //cmd/myapp1
```

### Build a target for a given platform

```
bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/myapp1
```

## Bazel Run

The bazel runs the specific target.

### Run a given target's binary

```
bazel run //cmd/myapp1
```

## Bazel Gazelle - Update Go dependencies and BUILD.bazel files

The command do update all of the BUILD.bazel files if imports change in a go file

```
bazel run //:gazelle update-repos
```

## Bazel Query

The Query command is a way to display information about your workspace.

### Bazel query the label of file
```
bazel query cmd/myapp1/main.go
```

### Bazel query a target

```
bazel query //cmd/...
bazel query //cmd/myapp1
bazel query //cmd/myapp1/internal/app.go
bazel query //cmd/myapp2
```

### Bazel query all targets or of a package

```
bazel query //...
bazel query cmd/...
```

## Bazel query to list out all platforms
```

bazel query 'kind(platform, @io_bazel_rules_go//go/toolchain:all)'
```

### Bazel query to list out all platform conditions

```
bazel query @bazel_tools//src/conditions:all
```

### Bazel query of packages

```
bazel query //... --output package
```

### Bazel query output label_kind of all targets

```
bazel query //... --output label_kind
```

### Bazel query type rules

```
bazel query 'kind(rule, //...)' --output label_kind
```
### Bazel query for binaries

```
bazel query 'kind(.*_binary, //...)' --output label_kind
```
### Bazel query for test cases

```
bazel query 'kind(.*_test, //...)' --output label_kind
```

## Bazel Tests

Run tests using Bazel

### Run all tests

```
bazel tests //...
```

## Compute coverage

```
bazel coverage //...
```

## Execute a binary
To run a given built binary us the Bazel run command

## Action Query
Executes a query on the post-analysis action graph.

```
bazel aquery 'deps(//...)'
```

## Buildifier

```
buildifier --mode=fix BUILD.bazel
```

# References

- https://github.com/systemlogic/learn-bazel
- https://skia.googlesource.com/buildbot/+/refs/heads/main/BAZEL_CHEATSHEET.md
