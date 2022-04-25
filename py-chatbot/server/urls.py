from django.urls import path
from django.conf.urls import url
from . import views,serch
 
urlpatterns = [
    path('runoob/', views.runoob),
    path('hello/', views.hello),
    url(r'^chat$', serch.search_post),
    url(r'^chat_rand$', serch.search_rand_post),
]