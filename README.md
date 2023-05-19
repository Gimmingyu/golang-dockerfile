# golang-dockerfile

## Instruction

### FROM

`FROM` 인스트럭션은 도커 이미지의 바탕이 될 베이스 이미지를 지정한다.

`Dockerfile`로 이미지를 빌드 할 때 먼저 `FROM` 인스트럭션으로 지정된 이미지를 받아온다.

### RUN

`RUN` 인스트럭션은 도커 이미지가 빌드될 때 컨테이너 안에서 실행할 명령을 정의하는 인스트럭션이다.

### COPY

`COPY` 인스트럭션은 동작중인 호스트 머신의 파일, 디렉터리를 도커 컨테이너 안으로 복사하는 인스트럭션이다.

### CMD

`CMD` 인스트럭션은 도커 컨테이너를 실행할 때 컨테이너 안에서 실행할 프로세스를 지정한다.

### ENTRYPOINT

`ENTRYPOINT` 인스트럭션은 컨테이너의 명령 실행 방식을 조절할 수 있다.

`CMD`와 마찬가지로 컨테이너를 실행 할 때 실행 할 프로세스를 지정하는 인스트럭션으로 `ENTRYPOINT` 인스트럭션을 지정하면 `CMD`의 인자가 `ENTRYPOINT` 에서 실행하는 파일에 인자로 주어진다.

만약 `ENTRYPOINT` 를 사용하여 컨테이너 수행 명령을 정의한 경우, 해당 컨테이너가 수행될 때 반드시 `ENTRYPOINT` 에서 지정한 명령을 수행되도록 지정 한다.

하지만, `CMD`를 사용하여 수행 명령을 경우에는 컨테이너를 실행할때 인자값을 주게 되면 `Dockerfile` 에 지정된 CMD 값을 대신 하여 지정한 인자값으로 변경하여 실행되게 된다.

## Image build

```bash
docker image build -t go_server:latest

# docker image build -t ${이미지명}:${태그명} ${Dockerfile 경로}
```

-t 옵션을 사용해 이미지명과 태그명을 지정할 수 있으며 태그명을 생략시 `latest`로 자동적으로 태그가 붙는다.

```bash
docker image ls

# docker container image 확인 
```

> Dockerfile이 아닌 다른 이름의 파일로 빌드하고 싶은 경우 -f 옵션을 사용하여 빌드할 수 있다.
>
> ```bash
> docker image build -f Dockerfile-example -t go_server:latest
> ```

## Docker commands

### 컨테이너 이름 짓기

```bash
docker container run -t --name my_server go_server:latest
```

위와 같이 `my_server`라는 이름으로 컨테이너를 실행할 때 이름을 지어줄 수 있다.

### 컨테이너 중지

```bash
docker container stop $(docker container ls --filter "ancestor=go_server" -q)
```

### 컨테이너 재시작

```bash
docker container restart 
```

### 컨테이너 제거

```bash
docker container rm go_server:latest
```

중지된 컨테이너를 완전히 파기하려면 위와 같이 `rm` 명령어를 이용해 파기 할 수 있다.

다음과 같이 여러개의 중지된 컨테이너들을 확인 할 수도 있다.

```bash
docker container ls --filter "status=exited"

docker container run -t --name my_server --rm go_server:latest
```

`my_server` 라는 이름을 가진 컨테이너가 이미 존재 할 때 같은 이름으로 생성 시 오류를 발생시키는데, 이때 --rm 옵션을 적용해 컨테이너를 실행한다면 컨테이너 중지 시 제거까지 해주어 오류 발생 여지를 줄여준다.

### 가지고 있는 이미지 목록 조회

```bash
docker images

docker image ls
```

### 실행중인 컨테이너 목록 조회

```bash
docker container ls
```

### 포트 포워딩

```bash

docker container run -d -p 9000:8080 example/echo:latest
```

위와 같이 `-p` 옵션을 이용해 로컬 포트와 컨테이너 포트를 포워딩 할 수 있다.

### 컨테이너 목록 필터링

```bash
docker container ls --filter "필터명=값"
```

필터명에는 name(컨테이너 명), ancestor 등이 존재한다.

### 컨테이너에 표준출력 연결

```bash
docker container logs -f $(docker container ls --filter "ancestor=go_server" -q)
```

위와 같이 ancestor에 해당하는 컨테이너의 로그를 표준출력에 보여줄 수 있다.

### 실행중인 컨테이너에서 명령 실행

```bash
docker container exec echo pwd
```

위와 같이 현재 실행중인 컨테이너에서 현재 경로를 가져오는 echo pwd 라는 명령을 내릴 수 있다.

표준 입력 연결을 유지하는 `-i` 옵션과 유사 터미널을 할당하는 `-t` 옵션을 조합해 컨테이너를 셸에서 다룰수 있게 해준다.

### 파일 복사하기

```bash
docker container cp dummy.txt go_server:/usr
```

위와 같이 dummy.txt라는 파일을 도커 컨테이너인 go_server 하위 usr 폴더에 복사 할 수 있다.


## 여러 컨테이너 실행하기

서버를 운용할 때 하나로만 운영하는게 아니라 마이크로 서비스가 될 수도 있고 로드밸랜서가 필요할 수도 있다.

하여튼 다양한 의존관계로 하나의 시스템이 구성된다.

### docker compose

`docker compose`는 yaml 포맷으로 기술된 설정 파일로, 여러 컨테이너의 실행을 한번에 관리할 수 있도록 해준다.

```bash
> docker compose version
Docker Compose version v2.13.0
```
