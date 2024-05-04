# ./createEnvFiles.sh <sql_password>

echo "sql_password: $1"

checkIfExist() {
  if [ -f .env ];
    then
        rm -r .env
        touch .env
    else
        touch .env
fi
}

cd ./database
checkIfExist

echo "MYSQL_ROOT_PASSWORD=${1}" >> .env
echo "MYSQL_DATABASE=destination_spot" >> .env
echo "MYSQL_USER=user" >> .env
echo "MYSQL_PASSWORD=${1}" >> .env

cd ..