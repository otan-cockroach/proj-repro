To reproduce:

```
git submodule init --recursive
bazel clean
bazel build //:all
```
