#Requires AutoHotkey v2.0

; Win+Shift+R to pull clipboard from another machine
#+r::Run 'pinger.exe --address=localhost:3000 CLIPBOARD_PULL_START'