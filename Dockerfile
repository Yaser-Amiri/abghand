FROM scratch
LABEL maintainer="yaser.amiri95@gmail.com"
COPY abghand /
ENTRYPOINT ["/abghand"]