# pull official base image
FROM jjanzic/docker-python3-opencv

# set work directory
WORKDIR /app

# install dependencies
RUN pip install --upgrade pip
COPY ./requirements.txt /app/requirements.txt
RUN pip install -r requirements.txt

# copy project
COPY . /app/

EXPOSE 8083

CMD python microservices/face_features/delivery/grpc/server/python_app/main.py