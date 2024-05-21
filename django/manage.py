#!/usr/bin/env python3
"""Django's command-line utility for administrative tasks."""
import os
import sys

# HOW I MADE THIS?
# 
# install django,mysqlclient,djangorestframework
# django-admin startproject djangoproject .
# python3 manage.py startapp app
# add app and rest_framework to installed-apps in setting.py
# change database section in settings.py to mysql
# DATABASES = {
#     "default": {
#         "ENGINE": "django.db.backends.mysql",
#         "NAME": "djangodb",
#         "USER": "root",
#         "PASSWORD": "pass",
#         "HOST": "127.0.0.1",
#         "PORT": "3306",
#     }
# }
# setup the db schema in models.py
# class User(models.Model):
#     id = models.AutoField
#     name = models.CharField(max_length=50)
#     age = models.IntegerField()
#
#     class Meta:
#         db_table = 'user'
# python3 manage.py makemigrations
# python3 manage.py migrate
# create serializers.py
# from rest_framework import serializers
# from .models import User

# class UserSerializer(serializers.ModelSerializer):
#     class Meta:
#         model = User
#         fields = ['id', 'name', 'age']

def main():
    """Run administrative tasks."""
    os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'djangoproject.settings')
    try:
        from django.core.management import execute_from_command_line
    except ImportError as exc:
        raise ImportError(
            "Couldn't import Django. Are you sure it's installed and "
            "available on your PYTHONPATH environment variable? Did you "
            "forget to activate a virtual environment?"
        ) from exc
    execute_from_command_line(sys.argv)


if __name__ == '__main__':
    main()
