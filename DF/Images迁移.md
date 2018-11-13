### 将镜像保存为压缩包文件

    [root@CentOS-7 ~]# docker save  nginx | gzip > nginx-latest.tar.gz



### 加载镜像

    [root@CentOS-7 ~]#docker load -i nginx-latest.tar.gz 


通过这种方式进行镜像的快速迁移。带名字的迁移过程。
