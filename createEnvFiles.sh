# ./createEnvFiles.sh <sql_password> <domain> <jwtSecretKey>

echo "sql_password: $1"
echo "domain: $2"
echo "jwtSecretKey: $3"

checkEnvFile() {
  cd $1

  if [ -f .env ];
    then
        rm -r .env
        touch .env
    else
        touch .env
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

checkEnvFile ./backend/services/auth
echo "CONNECTION_STRING=user:${1}@tcp(database:3306)/destination_spot" >> .env
echo "DOMAIN=${2}" >> .env
echo "JWT_SECRET_KEY=${3}" >> .env
cd ../../..

checkEnvFile ./backend/services/core
cd ../../..