FROM docker.gf.com.cn/ubuntu:14.04.5

ADD ./hello /hello

CMD ["/hello"]
