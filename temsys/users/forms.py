from django import forms

class LoginForm(forms.Form):
    username = forms.CharField(label='Name', max_length=20)
    password = forms.CharField(label='Password',
        max_length=32,
        min_length=2,
        widget=forms.PasswordInput
    )
