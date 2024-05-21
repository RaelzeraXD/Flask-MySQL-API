from django.db import models

# Create your models here.
class User(models.Model):
    id = models.AutoField
    name = models.CharField(max_length=50)
    age = models.IntegerField()

    class Meta:
        db_table = 'user'