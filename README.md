# lxd-tui

A TUI for LXD. The goal is to implement a k9s for LXD. Right now it supports:

- List projects and instances (toggle with the tab key)
- Select and delete multiple instances at once (use space/enter to select and ctrl-d to delete)
  - The UI refreshes automatically so you see progress in real time

It uses the DAO pattern so it can be extended to support different container managers in the future, like Incus.

Example:
```
// InstanceDAO is the interface that wraps the basic instance methods.
// DAO allows to abstract the data layer from the business logic.
// For example, the UI layer can call the GetInstances method without knowing
// how the data is retrieved. This allows to easily switch between different
// data sources (e.g. LXD, Incus, etc.) without changing the UI code.
```

## Run

`go run cmd/main.go`
