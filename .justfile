set export

bin := "./bin/calvu"

build:
  go build -o {{bin}}

install: build
  chmod +x bin/calvu
  ln -sf $(pwd)/bin/calvu /usr/local/bin/calvu
  echo "✅ 'calvu' is now available globally!"


