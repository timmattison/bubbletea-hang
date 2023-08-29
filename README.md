# bubbletea-hang

This repository shows how Bubbletea can hang if a simple program panics

## How to reproduce - the working code

This shows the program running as normal:

```sh
go run main1.go model.go
```

## How to reproduce - the hanging code

WARNING! You will need to have another terminal open to kill this next program since it will hang.

```sh
go run main2.go model.go
```

AGAIN, this will hang and will need to be killed from another terminal!

The issue appears to be in
the [standard renderer](https://github.com/charmbracelet/bubbletea/blob/8f3464a75600e991bbce22229d6e5b99975416f0/standard_renderer.go#L121)
where it writes to the `r.done` channel but nobody ever receives the message.
