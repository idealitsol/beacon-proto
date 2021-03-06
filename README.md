# beacon-proto
gRPC protobuf generator for beacon

### Versioning

To bump the version automatically, install bumpversion and move the pre-commit script into your githooks folder:

```bash
$ pip install bumpversion
$ cp pre-commit .git/hooks
```

The script will bump the path version every time you commit something.
If you need to change the minor or major versions you can do the following **(Don't forget to disable the git hook before committing)**:

```bash
# minor
$ bumpversion --allow-dirty minor

# major
$ bumpversion --allow-dirty major
```
### Tagging

To tag a version automatically, copy the post-commit into your local .git/hooks folder:

```bash
cp post-commit .git/hooks
```


### Linting Protoc files

Install the [protolint tool](https://github.com/yoheimuta/protolint) for linting or fixing `.protoc` file

```bash
brew tap yoheimuta/protolint
brew install protolint
```