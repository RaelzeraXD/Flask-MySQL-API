FROM python

WORKDIR /app

COPY . .

RUN pip install -r requirements.txt

EXPOSE 8000

# the migrate command needs to mysql be up, thats why i move from RUN to CMD 
CMD ["sh", "-c", "python3 manage.py makemigrations && python3 manage.py migrate && python3 manage.py runserver 0.0.0.0:8000"]