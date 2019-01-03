# ServerSideForRxPractice
server side program for RxPractice

### docker command

#### docker container まとめて停止

```
docker stop $(docker ps -q)
```

#### docker container まとめて削除

```
docker container prune
```

もしくは、

```
docker rm $(docker ps -q -a)
```

#### docker image まとめて削除

```
docker image prune -a
```

```
docker rmi $(docker images -q)
```

##### 参考

https://qiita.com/shisama/items/48e2eaf1dc356568b0d7

#### docker container 内に入る

```
docker container exec -it <IDなど> <command>
```

よく使うやつ

```
docker container exec -it <ID> mysql -u root -p
```

#### docker-compose 

image をビルドして実行

```
docker-compose up -d --build
```

```
docker-compose up -d
```

### mysql command

#### option

##### -u

ユーザー名指定

##### -p

パスワード指定

##### -D

データベース名指定