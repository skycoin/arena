echo "Build Arena..."
if [ -f "./bin/arena" ]; then
    rm ./bin/arena
fi
go build -o ./bin/arena

echo "-------------------------------------"
echo "Arena demos."
echo "-------------------------------------"


echo "1. Run arena: \n"
echo "./bin/arena \n"
echo ">> Result: \n"

./bin/arena


echo "\n-------------------------------------"
echo "2. Get version: \n"
echo "./bin/arena --version \n"
echo ">> Result: \n"

./bin/arena --version


echo "\n-------------------------------------"
echo "3. Run help: \n"
echo "./bin/arena --help add \n"
echo ">> Result: \n"

./bin/arena --help add


echo "\n-------------------------------------"
echo "4. Run add (OK, Scenario 1): \n"
echo "./bin/arena add 1 2\n"
echo ">> Result: \n"

./bin/arena add 1 2


echo "\n-------------------------------------"
echo "5. Run add (OK, Scenario 2): \n"
echo "./bin/arena add 1,2 2.3\n"
echo ">> Result: \n"

./bin/arena add 1,2 2.3


echo "\n-------------------------------------"
echo "6. Run add (ERROR, Scenario 1): \n"
echo "./bin/arena add \n"
echo ">> Result: \n"

./bin/arena add


echo "\n-------------------------------------"
echo "7. Run add (ERROR, Scenario 2): \n"
echo "./bin/arena add a b \n"
echo ">> Result: \n"

./bin/arena add a b

echo "\n-------------------------------------"