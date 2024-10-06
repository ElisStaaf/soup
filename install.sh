#!/bin/bash
DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

if [ ! -f  $DIR/soup ]; then
      echo "Building executable..."
      go build $DIR/soup.go
      echo "Build successful. You can now run the executable."
else
    echo "Nothing to do."
fi
exit(0)
