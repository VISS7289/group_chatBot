# -*- coding: utf-8 -*-
import sys
sys.path.append(r"../myChat")
from django.shortcuts import render
from django.views.decorators import csrf
from django.http import HttpResponse
import json
from myChat import myChat

# 接收POST请求数据
def search_post(request):
    json_str = request.body
    json_dict = json.loads(json_str)
    key = json_dict.get("chat_content", None)
    return HttpResponse(myChat.chatWithMyChat(key))

def search_rand_post(request):
    return HttpResponse(myChat.chatWithRandom()[0])