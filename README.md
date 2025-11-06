# bubbles-hlist

`bubbles-hlist` is a horizontal variant of Charm's [`bubbles/list`](https://github.com/charmbracelet/bubbles) component.

It lays out items side by side instead of top to bottom, while keeping (most of)
the original list behavior.

![A GIF showing the usage of hlist](https://github.com/marcelblijleven/bubbles-hlist/blob/main/examples/hlist.gif)

## Features

Just like the original list component, you get:

- Pagination
- Fuzzy filtering
- Auto-generated help
- Activity spinner
- Status messages

All of these can be configured or replaced by the user.

## Usage

The API is almost exactly the same as `bubbles/list`, so it should feel familiar
if you’ve used the original component.

The main difference: because items are arranged horizontally, you can specify
the width of each item "cell" through the delegate.

You do that with:

```go
delegate.SetWidth(width)
```

## Attribution

This project’s `list` package is based on the [`list`](https://github.com/charmbracelet/bubbles/tree/master/list) component from [Charmbracelet/bubbles](https://github.com/charmbracelet/bubbles),  
modified to adjust the functionality from a vertical list to a horizontal list.

Original code © 2020-2025 Charmbracelet, released under the MIT License.
