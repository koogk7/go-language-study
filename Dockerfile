# 이미지 설정
FROM golang:1.13

# 작성자 정보
MAINTAINER pnu.lee

# 이미지 내부의 커맨드가 실행될 디렉토리 설정
WORKDIR ./app

# ./example 디렉토리에 있는 파일들을 이미지 내부 /app 디렉터리에 추가
COPY /example .
COPY ./docker.sh .  


CMD /bin/bash -c "source ./docker.sh"

# Port 설정
EXPOSE 8080
