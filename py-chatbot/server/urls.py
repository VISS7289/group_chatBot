from django.urls import path
from django.urls import re_path as url
from . import views,serch
 
urlpatterns = [
    path('runoob/', views.runoob),
    path('hello/', views.hello),
    # 聊天接口
    url(r'^chat$', serch.search_post),
    # 随机聊天接口
    url(r'^chat_rand$', serch.search_rand_post),
]