# Docker MongoDB 사용법
출처 : https://poiemaweb.com/docker-mongodb
1. Docker 설치
2. MongoDB Docker 이미지 다운로드
```shell
docker pull mongo
```
3. MongoDB 실행
```shell
docker run --name mongodb-container -v ~/data:/data/db -d -p 27017:27017 mongo
```

4. Docker MongoDB CLI 접속
   1. 도커 웹 프로그램에서 접속
   2. 명령어를 통해 접속
   ```shell
    docker exec -it mongodb-container bash
   $ mongo
    ```

5. MongoDB Database 사용되는 명령
    ```shell
   # DB 생성
    use DATABASE_NAME;
   # 현재 사용중인 DB
   db
   # 생성된 DB 목록
    show dbs;
   # DB에 Collection생성하기
   db.createCollections("COLLECTION_NAME");
   
    ```