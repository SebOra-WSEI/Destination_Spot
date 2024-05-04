# ./createEnvFiles.sh <sql_password>

echo "sql_password: $1"

checkEnvFile() {
  cd $1

  if [ -f .env ];
    then
        rm -r .env
        touch .env
    else
        touch .envs
fi
}

checkEnvFile ./database
echo "MYSQL_ROOT_PASSWORD=${1}" >> .env
echo "MYSQL_DATABASE=destination_spot" >> .env
echo "MYSQL_USER=user" >> .env
echo "MYSQL_PASSWORD=${1}" >> .env
cd ..

checkEnvFile ./frontend
cd ..

checkEnvFile ./backend/auth
echo "CONNECTION_STRING=user:${1}@tcp(database:3306)/destination_spot" >> .env
cd ../..

checkEnvFile ./backend/core
cd ../..