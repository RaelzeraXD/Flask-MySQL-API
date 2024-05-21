from django.contrib import admin
from django.urls import path, include
from app import views

urlpatterns = [
    path('users/', views.get_users),
    path('users/<int:id>', views.get_users_by_id),
    path('create', views.create_user),
    path('update/<int:id>', views.update_user),
    path('delete/<int:id>', views.delete_user),
]
