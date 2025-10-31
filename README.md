# bubbles-hlist

`bubbles-hlist` is a horizontal variant of Charm's [`bubbles/list`](https://github.com/charmbracelet/bubbles) component.

It lays out items side by side instead of top to bottom, while keeping (most of) the original list behavior.

## Features

Just like the original list component, you get:

- Pagination
- Fuzzy filtering
- Auto-generated help
- Activity spinner
- Status messages

All of these can be configured or replaced by the user.

## API

The API is almost exactly the same as `bubbles/list`, so it should feel familiar if you’ve used the original component.

The main difference: because items are arranged horizontally, you need to specify the width of each item "cell".

You do that with:

```go
list.SetCellWidth(width)
```
