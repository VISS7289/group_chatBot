from django.http import HttpResponse
from django.shortcuts import render

def hello(request):
    return HttpResponse("hello")

def runoob(request):
    context          = {}
    context['hello'] = '菜鸟教程'
    return render(request, 'runoob.html', context)