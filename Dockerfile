FROM google/golang-runtime

RUN curl https://yt-dl.org/latest/youtube-dl -o /bin/youtube-dl && \
    chmod a+rx /bin/youtube-dl

RUN mkdir /steamerDataDir
VOLUME ["/steamerDataDir"]
