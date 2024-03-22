
FROM centos:latest
ENV MYPATH /root/trend-ring
WORKDIR  $MYPATH
# RUN yum -y update
RUN yum -y install vim
RUN yum -y install git
RUN yum install -y gcc-c++
RUN yum -y install wget
RUN wget -P /root/ https://dl.google.com/go/go1.22.0.linux-amd64.tar.gz
RUN tar -zxvf /root/go1.22.0.linux-amd64.tar.gz -C /usr/local
RUN echo export PATH=$PATH:/usr/local/go/bin >> /etc/profile
RUN source /etc/profile && go version
RUN echo "source /etc/profile" >> /root/.bashrc
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GO111MODULE=on
