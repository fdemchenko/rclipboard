# Remote Clipboard 💻🔁💻

This application allows you to share clipboard text with computers across the network with TCP sockets on both Linux and Windows.

## How does it work?

Program follows `pull` model, where user triggers action to pull clipboard from another machine by making TCP request. Read more about triggers and their delivery to the app in the next section.

## Triggers and delivery

- On Linux systems app waiting for SIGUSR1 signals to trigger pull action.
- On Windows machines app listening local TCP socket for `CLIPBOARD_PULL_START` message. Maybe this will be changed in the future and replaced with another IPC method.

Later you can create custom key mappings in your OS, to control clipbord pulling. See [this](deployment/) directory for references.

Recommended way to start app is to create user specific systemd service on Linux, and scheduled task on Windows.

## Configuration
```
Usage of rclipboard:
  -listen-addr string
        TCP listen address
  -remote-addr string
        TCP remote address of another clipboard
```




