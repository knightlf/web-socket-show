#FROM octahub.8lab.cn:5000/base/cicd-go:v1.18.2
FROM octahub.8lab.cn:5000/base/ubuntu:18.04
MAINTAINER devcli-server
### sources.list from http://140.143.184.36/ octa main
#COPY sources.list /etc/apt/sources.list
RUN mkdir -p /data/devcli/conf
COPY ./conf/tsconfig.json /data/devcli/conf/
COPY ./web-socket-show /data/devcli/
COPY ./Docker/entrypoint.sh /usr/bin/entrypoint.sh
#COPY neighbors /iri/conf/neighbors

#RUN apt-get update && apt-get install -y --force-yes pypy vim libmysqlclient-dev python-dev python-pip curl libblas-dev liblapack-dev libatlas-base-dev gfortran 
#RUN apt-get update && apt-get install -y  vim tzdata
#RUN mkdir /data


# Setting this to 1 will have socat exposing 14266 and pointing it on
# localhost. See /entrypoint.sh
# !!! DO NOT DOCKER EXPOSE (-p) 14266 as the remote api settings
#     will not be applied on that port !!!
# You also have to maintain $DOCKER_IRI_MONITORING_API_PORT_DESTINATION
# based on the actual API port exposed via IRI
#ENV DOCKER_IRI_MONITORING_API_PORT_ENABLE 0
#ENV DOCKER_IRI_MONITORING_API_PORT_DESTINATION 14700
# PORT
EXPOSE 9055
EXPOSE 9033

#VOLUME ["/iri/conf","/iri/data"]
WORKDIR /data/devcli
ENTRYPOINT [ "entrypoint.sh" ]

