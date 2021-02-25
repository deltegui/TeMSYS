import os
from flask import Flask
from oligo import Iber

connection = Iber()
connection.login(os.environ.get('IBER_USER'), os.environ.get('IBER_PASS'))

app = Flask(__name__)

@app.route('/')
def serve_data():
    watt = connection.watthourmeter()
    return {
        "watts": watt
    }
