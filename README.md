# Snek

[![cacois](https://circleci.com/gh/cacois/snek.svg?style=svg)](https://app.circleci.com/pipelines/github/cacois/snek)
[![Coverage Status](https://coveralls.io/repos/github/cacois/snek/badge.svg?branch=master)](https://coveralls.io/github/cacois/snek?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/cacois/snek?style=flat-square)](https://goreportcard.com/report/github.com/cacois/snek)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Snek** is a very simple Go environment variable-based config management system, designed for docker/kubernetes use cases. Very small fangs.

## Why?

Go already has some great config management tools, like [Viper](https://github.com/spf13/viper). Powerful...but complex. Managing precedence between config sources, many file formats, aliases...sometimes your needs are simpler.

Like many developers, I find myself working in a Docker/Kubernetes world more often than not. So, who needs config files? All I need is environment variables and/or a [ConfigMap](https://kubernetes.io/docs/concepts/configuration/configmap/). This simplifies config management from within my app - all I really need is something that looks for my config values in environment variables, and allows me to set default values in case the desired environment variables are empty. Fin.

I'm tired of writing this (admittedly simple) logic into each of my apps, so here we are. :)

## Usage

Start by importing the module:

```bash
$ go get "github.com/cacois/snek"
```

Snek is really simple. You can do two things. First, if you want to, you can set default values for a particular configuration parameter:

```go
snek.Default("MY_ENV_VAR", "mydefaultvalue")
```

Then you can read your configuration values from anywhere in your app:

```go
value := snek.Get("MY_ENV_VAR")
```

This will first look for and return the value of the environment variable `MY_ENV_VAR`. If this environment variable has not been set, snek will look for any default value you defined for `MY_ENV_VAR` and return that. If neither has been set, it will return an empty string.

One more option is to use `snek.GetOrError()` instead of `snek.Get()`:

```go
value, err := snek.GetOrError("MY_ENV_VAR")
if err != nil {
    // do something about it
}
```

This function behaves the exact same way as `snek.Get()`, except instead of returning an empty string when neither the specified environment variable or a default value has been defined, it throws an error. 

## Patterns

### Inline Config Value Access

The easiest way to use snek is just to set all your defaults in your `main.go` file, in the `init()` or `main()` function:

main.go:
```go
package main

func init() {
    snek.Default("CONFIG_VAL_1", "somevalue")
    snek.Default("CONFIG_VAL_2", "anothervalue")
}
...
```

Its convenient to set the defaults for all config values in your app in a single place, to easily keep track of them.

You can then pull your config values from anywhere in your app, except from `init()` functions or package-level variables, since those pieces of code will be [executed before](https://yourbasic.org/golang/package-init-function-main-execution-order/) the `init()` function in `main.go` - meaning you will try to access your variables before your defaults have been set. (If you don't set or care about default values, this warning is irrelevant)  I refer to this as "inline" variable access, since it means I'm accessing the values from within my business logic, inside functions:

someotherfile.go:
```go
package somepackage

func SomeFunction() {
    // if env var CONFIG_VAL_1 is empty, default value "somevalue" will be returned
    value1 := snek.Get("CONFIG_VAL_1") 
}
```

### Package-level Config Value Access

Sometimes, I like to be able to also set my config values in package-level variables:

somefile.go:
```go
package somepackage

var (
    value1 := snek.Get("CONFIG_VAL_1") // package-level variable
)

func SomeFunction() {
    fmt.Sprintf("value1: %s", value1)
}
```

This presents an interesting challenge, because package-level variable definitions are executed very early in Go - before your entrypoint `main()` function or even the `init()` function in your `main.go` file. This means all my default values are not yet set, `value1` above is set as `""`. So, to be able to assign the variables in this way, I tend to make a package called `config` with a single file named `config.go`:

config.go:
```go
package config

func init() {
    snek.Default("CONFIG_VAL_1", "value1")
    snek.Default("CONFIG_VAL_2", "value2")
    snek.Default("CONFIG_VAL_3", "value3")
}

func Init() {
    // dummy function to allow me to force early execution of
    // the above init() function
}
```

I then put a big list of default values in the `init()` function. To make sure this function is executed to set the defaults before any other package-level code, I add an empty `Init()` function and call it early in the `main()` function in `main.go` (you have to call something in the `config` package, or Go will not allow you to import it):

main.go:
```go
package main

import config

func main() {
    config.Init()
}
```

A little legwork, but problem solved.