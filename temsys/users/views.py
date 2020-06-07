from django.shortcuts import render, redirect
from django.views.decorators.http import require_http_methods
from .forms import LoginForm
from django.contrib.auth import authenticate, login, logout

@require_http_methods(['GET'])
def login_index(request):
    if request.user.is_authenticated:
        return redirect('/panel')
    return render(request, 'users/index.html', {'form': LoginForm()})


@require_http_methods(['POST'])
def process_login(request):
    form = LoginForm(request.POST)
    if not form.is_valid():
        return redirect('login_index')
    data = form.cleaned_data
    user = authenticate(username=data['username'], password=data['password'])
    if user is None:
        return redirect('login_index')
    login(request, user)
    return redirect('/panel')


@require_http_methods(['GET'])
def do_logout(request):
    logout(request)
    return redirect('/user')
