FROM docker.gf.com.cn/ubuntu:14.04.4

ADD ./hello /hello

CMD ["/hello"]
