from django.http import HttpResponse
from django.shortcuts import render

def hello(request):
    return HttpResponse("hello")

def runoob(request):
    context          = {}
    context['hello'] = 'θιΈζη¨'
    return render(request, 'runoob.html', context)