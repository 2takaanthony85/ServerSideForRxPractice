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

containerをまとめて停止、削除

```
docker-compose down
```

### mysql command

#### option

##### ユーザー名指定

```
-u
```

##### パスワード指定

```
-p
```

##### データベース名指定

```
-D
```

##### 参考

https://qiita.com/yulily@github/items/54cb6ccaacf39977455c