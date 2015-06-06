FROM google/golang-runtime

RUN curl https://yt-dl.org/latest/youtube-dl -o /bin/youtube-dl && \
    chmod a+rx /bin/youtube-dl

RUN mkdir /steamerDataDir
RUN ls /gopath/bin/team-iDareX &>/dev/null && ln -s /gopath/bin/team-iDareX /gopath/bin/app
WORKDIR /gopath/src/app

VOLUME ["/steamerDataDir"]
