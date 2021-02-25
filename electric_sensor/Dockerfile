FROM python:3.10.0a5-buster
WORKDIR /app
COPY . .
RUN pip install -r requirements.txt
RUN pip install cffi
RUN pip install gunicorn[gevent]
EXPOSE 5000
ENV IBER_USER=""
ENV IBER_PASS=""
CMD gunicorn --worker-class gevent --workers 2 --bind 0.0.0.0:5000 app:app --max-requests 10000 --timeout 100 --keep-alive 5 --log-level info