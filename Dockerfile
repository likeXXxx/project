FROM internal-registry.ghostcloud.cn/ekos/alpine:3.6

COPY ProjectManage /
COPY views /views
COPY conf /conf
COPY bootstrap-4.3.1-dist /bootstrap-4.3.1-dist
COPY jquery /jquery
COPY static /static

RUN chmod +x ProjectManage

ENV BEEGO_RUNMODE=prod
ENTRYPOINT ["./ProjectManage"]
