from django.shortcuts import render
from rest_framework.decorators import api_view
from rest_framework.response import Response
from rest_framework import status

from .models import User
from .serializers import UserSerializer

@api_view(['GET'])
def get_users(request):
    if request.method == "GET":
        queryset = User.objects.all()
        serializer = UserSerializer(queryset, many=True)
        return Response(serializer.data)
    return Response(status.HTTP_400_BAD_REQUEST)

@api_view(['GET'])
def get_users_by_id(request,id):
    if request.method == "GET":
        try:
            query = User.objects.get(pk=id)
            serializer=UserSerializer(query)
            return Response(serializer.data)
        except:
            return Response(status.HTTP_400_BAD_REQUEST)

@api_view(["GET",'POST'])        
def create_user(request):
    if request.method == "POST":
        new_user = request.data
        serializer = UserSerializer(data=new_user)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data,status=status.HTTP_201_CREATED)
        return Response(status.HTTP_404_NOT_FOUND)
    return Response(status.HTTP_200_OK)

@api_view(["GET","PUT"]) 
def update_user(request, id):
    if request.method == "PUT":
        user = User.objects.get(pk=id)
        
        serializer = UserSerializer(user, data=request.data)
        
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=status.HTTP_200_OK)
        else:
            return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)
    return Response(status=status.HTTP_200_OK)

@api_view(["GET","DELETE"]) 
def delete_user(request,id):
    if request.method == "DELETE":
        try:
            user_to_delete = User.objects.get(pk=id)
            user_to_delete.delete()
            return Response(status=status.HTTP_202_ACCEPTED)
        except:
            return Response(status=status.HTTP_400_BAD_REQUEST)
    return Response(status=status.HTTP_200_OK)