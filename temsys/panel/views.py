from django.shortcuts import render
from .gateways import Sapi
from .entities import Sensor

# Create your views here.
def system_status(request):
    sapi = Sapi('localhost:8080')
    response = sapi.get_sensor_status(Sensor(name='salon'))
    return render(request, 'panel/status.html', {
        'response': response,
    })
