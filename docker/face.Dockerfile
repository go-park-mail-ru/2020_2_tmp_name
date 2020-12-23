FROM python:3.8-slim-buster

RUN python3 -m venv /opt/venv
RUN pip3 freeze > requirements.txt

# Install dependencies:
COPY requirements.txt .
RUN . /opt/venv/bin/activate && pip install -r requirements.txt

# Run the application:
COPY myapp.py .
CMD . /opt/venv/bin/activate && exec python myapp.py