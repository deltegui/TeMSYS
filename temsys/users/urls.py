from django.urls import path

from . import views

urlpatterns = [
    path('', views.login_index, name='login_index'),
    path('login', views.process_login, name='process_login'),
    path('logout', views.do_logout, name='do_logout'),
]
