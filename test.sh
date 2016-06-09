
#!/bin/bash
# required package: inotify-tools

events=(-e CREATE -e MODIFY -e MOVED_TO)
while inotifywait ${events[@]} -r "./"; do
    go test ./proAt -v
    go run ./main.go 
done
