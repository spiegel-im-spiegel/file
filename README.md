# [file] -- Extend filepath.Glob function

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/file.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/file)
[![GitHub license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/file/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/file.svg)](https://github.com/spiegel-im-spiegel/file/releases/latest)

## Usage

```go
matches, err := file.Glob("**/*.[ch]")
if err != nil {
    fmt.Fprintf(os.Stderr, "%+v\n", err)
    return
}
for _, path := range matches {
    fmt.Println(path)
}
// Output:
// testdata/include/source.h
// testdata/source.c
```

### Glab with context.Context

```go
matches, err := file.GlobWithContext(context.Background(), "**/*.[ch]")
if err != nil {
    fmt.Fprintf(os.Stderr, "%+v\n", err)
    return
}
for _, path := range matches {
    fmt.Println(path)
}
// Output:
// testdata/include/source.h
// testdata/source.c
```

### Glab with flags

```go
matches, err := file.Glob("**/*.[ch]", file.WithFlags(file.StdFlags|file.AbsolutePath))
if err != nil {
    fmt.Fprintf(os.Stderr, "%+v\n", err)
    return
}
for _, path := range matches {
    fmt.Println(path)
}
// Output:
// /home/username/go/src/github.com/spiegel-im-spiegel/file/testdata/include/source.h
// /home/username/go/src/github.com/spiegel-im-spiegel/file/testdata/source.c
```

| Flag             | Note                                                |
| ---------------- | --------------------------------------------------- |
| `ContainsFile`   | contains file                                       |
| `ContainsDir`    | contains directory                                  |
| `SeparatorSlash` | use slash character for separator character in path |
| `AbsolutePath`   | output absolute representation of path              |
| `StdFlags`       | `ContainsFile \| ContainsDir` (default)             |

[file]: https://github.com/spiegel-im-spiegel/file "spiegel-im-spiegel/file: Extend filepath.Glob function"
