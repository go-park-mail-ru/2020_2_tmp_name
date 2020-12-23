FROM python:3.8-slim-buster

COPY microservices/face_features/delivery/grpc/server/python_app /opt

RUN python3 -m venv /opt/venv

# Install dependencies:
RUN . /opt/venv/bin/activate && pip install --upgrade pip -r opt/requirements.txt

# Run the application:
CMD . /opt/venv/bin/activate && exec python /opt/main.py