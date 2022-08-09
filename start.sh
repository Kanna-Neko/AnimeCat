cd backend
go build
cd ..
docker-compose up -d --build
rm ./backend/AnimeCat