# Snek

A very simple environment variable-based config management system, focused on container/kubernetes use cases. Very small fangs.

## Why?

Viper is powerful, but its complicated. Like many developers, I find myself working in a Docker/Kubernetes world more often than not. So, who needs config files? I just want envinroment variables and/or a ConfigMap. This simplifies config management from within my app - all I really need is to look for my config values in environment variables, and to use a default value if the environment variable isn't set. Fin.

I'm tired of writing this (admitttedly simple) logic into every app, so here we are :)

## Usage

Snek is really simple. You can do two things. First, if you want to, you can set default values for a particular configuration parameter:

```go
snek.Default("MY_ENV_VAR", "mydefaultvalue")
```

I tend to put a big list of these in a file called `config.go`, because 1) its extremely convenient to be able to see the entire application config in a single place, and 2) I like to have default values for all config parameters - its handy for running the app in a default development mode:

config.go:
```go 
package main

func init() {
    snek.Default("CONFIG_VAL_1", "value1")
    snek.Default("CONFIG_VAL_2", "value2")
    snek.Default("CONFIG_VAL_3", "value3")
    ...
}
```

Second, you can read your configuration values from environment variables:

```go
value := snek.Get("MY_ENV_VAR")
```

This will first look for and return the value of the environment variable `MY_ENV_VAR`. If this environment variable has not been set, snek will look for any default value you defined for `MY_ENV_VAR` and return that. If neither has been set, it will return an empty string.

One more option is to use `snek.GetOrError()` instead of `snek.Get()`:

```go
value, err := snek.GetOrError("MY_ENV_VAR")
```

This function behaves the exact same way as `snek.Get()`, except instead of returning an empty string when neither the specified environment variable or a default value has been defined, it throws an error. 

