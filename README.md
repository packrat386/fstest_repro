fstest_repo
-----------

Reproduce behaviour that fstest.TestFS thinks is an error, but I think should
not be considered an error. When TestFS attempts to open a file named `file.txt`
for example, it will also attempt to open `/file.txt` and if the latter opens it
considers it an error.

See: https://github.com/golang/go/blob/master/src/testing/fstest/testfs.go#L585


However if the other file exists simply by coincidence, or the filesystem being
tested is rooted at `/` the tests will fail.

To run, just `go run fstest_repro.go`. This may require elevated permissions
since it writes a file at `/file.txt`. For simplicity I've included a dockerfile
which you can run with `docker build -t fstest_repro:local . && docker run
fstest_repro:local`.

When I ran it I got the following output:

```
 [fg-386] fstest_repro > docker build -t fstest_repro:local . && docker run fstest_repro:local
Sending build context to Docker daemon   7.68kB
Step 1/4 : FROM golang:latest
 ---> 0debfc3e0c9e
Step 2/4 : WORKDIR /fstest_repro
 ---> Using cache
 ---> a761b21c49f2
Step 3/4 : COPY fstest_repro.go .
 ---> Using cache
 ---> e65642be0d55
Step 4/4 : CMD go run fstest_repro.go
 ---> Using cache
 ---> 5061046bfc5f
Successfully built 5061046bfc5f
Successfully tagged fstest_repro:local
ERROR:  TestFS found errors:
file.txt: Open(/file.txt) succeeded, want error
.: Open(/) succeeded, want error
```
