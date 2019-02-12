echo "Build Arena..."
if [ -f "./bin/arena" ]; then
    rm ./bin/arena
fi
go build -o ./bin/arena