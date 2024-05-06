FROM python:slim

WORKDIR /app

COPY ./app.py ./requirements.txt ./

RUN pip install -r requirements.txt

CMD ["python", "app.py"]