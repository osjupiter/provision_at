
#!/bin/bash
# required package: inotify-tools

events=(-e CREATE -e MODIFY -e MOVED_TO)
while inotifywait ${events[@]} "./"; do
    go test ./ -v
    go run ../main.go 
done
