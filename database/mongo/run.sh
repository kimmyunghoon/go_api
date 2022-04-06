echo "MongoDB Container 실행"
docker run --name mongodb-container -v ~/data:/data/db -d -p 27017:27017 mongo