# Automatically detecting and resolving deprecations using Semgrep demo

This repository is a demo that accompanies the blog post [Automatically detecting and resolving deprecations using Semgrep](https://www.lucasmelin.com/blog/2023-11-26/).

## Application

Here is our application. It wraps the code-buddy CLI.

Our code does the equivalent of:

```sh
YOURAPP="somename"
code-buddy run YOURAPP --please
```

You can run this application using:

```go
go run main.go "somename"
```
