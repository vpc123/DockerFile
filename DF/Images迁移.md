### 将镜像保存为压缩包文件

    [root@CentOS-7 ~]# docker save  nginx | gzip > nginx-latest.tar.gz



### 加载镜像

    [root@CentOS-7 ~]#docker load -i nginx-latest.tar.gz 


通过这种方式进行镜像的快速迁移。带名字的迁移过程。



### 快捷命令
将镜像从一个主机迁移到另一个主机：
docker save


    [root@CentOS-7 ~]# docker images
    REPOSITORY  TAG IMAGE IDCREATED SIZE
    docker.io/hello-world   latest  48b5124b27683 months ago1.84 kB
    [root@CentOS-7 ~]# 
    [root@CentOS-7 ~]# docker save hello-world | bzip2 | ssh root@10.140.1.120 "cat | docker load"
    root@10.140.1.120's password: 
    Loaded image: docker.io/hello-world:latest
    [root@CentOS-7 ~]# 