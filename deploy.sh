#!/bin/bash
ComposeConfig=/opt/web-system/blog-compose.yml
WorkDir=/opt/web-system/go-build/
RepoDir=/opt/web-system/go-build/gin-blog
Starttime=`date +"%Y-%m-%d_%H-%M-%S"`
ImageExistedNum=`docker images | grep alexcld/gin-blog | wc -l`
ImageTageNum=`expr $ImageExistedNum + 1`
CleanOldRepo() {
    cd $WorkDir
    rm -rf ./gin-blog
    PullCode
}

PullCode() {
    cd $WorkDir
    git clone https://gitee.com/alexcld/gin-blog.git
}

BuildImage() {
    cd $RepoDir
    docker build -t alexcld/gin-blog:0.0.$ImageTageNum .
}

UpdateImgae() {
    sed -ri "s@image: alexcld/gin-blog.*@image: alexcld/gin-blog:0.0.$ImageTageNum@g"  $ComposeConfig
    cd /opt/web-system/
    docker-compose -f blog-compose.yml up -d
}

main() {
    CleanOldRepo
    BuildImage
    UpdateImgae
}

main
