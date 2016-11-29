FROM docker.gf.com.cn/busybox

ADD ./hello /hello

CMD ["/hello"]
