FROM ubuntu:20.04
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install -y python3.9
RUN apt-get install -y libreoffice --no-install-recommends
RUN apt-get -y install default-jre
WORKDIR /app

COPY requirements.txt ./

# RUN python3 -m pip install -r requirements.txt

COPY mainLinux.py ./

EXPOSE 8080

CMD ["python3", "mainLinux.py"]