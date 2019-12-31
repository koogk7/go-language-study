# 이미지 설정
FROM golan:1.13

# 작성자 정보
MAINTAINER pnu.lee

# 이미지 내부의 커맨드가 실행될 디렉토리 설정
WORKDIR ./example

# ./example 디렉토리에 있는 파일들을 이미지 내부 /app 디렉터리에 추가
ADD ./example /app

CMD ["go", "run", "webserver"]

# Port 설정
EXPOSE 8080
