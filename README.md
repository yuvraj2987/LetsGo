# Repository to Learn Go Language
The directory structure is based on
https://github.com/golang-standards/project-layout.


# /cmd
All binaries

# /internal
All packages before they are ready to be exposed to everyone. Mostly playground
area.


## Build Individual Packages
```
	make build pkg=$(pkg_path)

```

e.g.

```
	make build pkg=internal/ccr/ch1

```

## Test Individual Packages
```
	make test pkg=$(pkg_path)

```


