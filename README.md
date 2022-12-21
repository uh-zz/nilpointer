# nilpointer

`nilpointer` checks for nil returning if the multiple return value contains pointer types.

`nilpointer` was created under the influence of [`nilerr`](https://github.com/gostaticanalysis/nilerr).

Powerd by [skeleton](https://github.com/gostaticanalysis/skeleton).

## How to use

```
go install github.com/uh-zz/nilpointer/cmd/nilpointer@latest
go vet -vettool=`which nilpointer` ./...
```

## Analyze

Checks for return of nil if multiple return values and pointer type is included.

```go
func do() (*int, error) {
	if err != nil {
		return nil, nil // NG
	}

	i1 := 1
	return &i1, nil // OK
}
```

Not checked if return value is one.

```go
type Something struct {
    E error
}

func do() *Something {
    if err := something(); err != nil {
        return &Something{E: err}
    }
    return nil // OK
}

if err := do(); err != nil {
    return err
}

```

`nilpointer` ignores code which has a miss with ignore comment.

```go
func do() (*int, error) {
	if err != nil {
        //lint:ignore nilpointer reason
		return nil, nil // ignore
	}

	i1 := 1
	return &i1, nil // OK
}
```
