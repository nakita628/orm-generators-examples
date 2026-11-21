"""Abstract, multi-table and proxy inheritance, and a custom user model."""

from django.contrib.auth.models import AbstractUser
from django.db import models


class User(AbstractUser):
    """The project's own user, in place of auth.User (AUTH_USER_MODEL)."""

    nickname = models.CharField(max_length=30, blank=True)


class Timestamped(models.Model):
    created = models.DateTimeField(auto_now_add=True)
    modified = models.DateTimeField(auto_now=True)
    owner = models.ForeignKey(User, on_delete=models.CASCADE, related_name="%(app_label)s_%(class)s_owned")

    class Meta:
        abstract = True
        ordering = ["-created"]


class Article(Timestamped):
    title = models.CharField(max_length=100)


class Video(Timestamped):
    url = models.URLField()

    class Meta(Timestamped.Meta):
        ordering = ["url"]


class Place(models.Model):
    name = models.CharField(max_length=50)
    neighbours = models.ManyToManyField("self", blank=True)


class Restaurant(Place):
    serves_pizza = models.BooleanField(default=False)


class Pizzeria(Restaurant):
    oven = models.CharField(max_length=20)


class Cafe(Place):
    # The link to the parent named, rather than place_ptr
    place = models.OneToOneField(Place, on_delete=models.CASCADE, parent_link=True, primary_key=True, related_name="cafe")
    has_wifi = models.BooleanField(default=True)


class OpenRestaurant(Restaurant):
    class Meta:
        proxy = True


class Step(models.Model):
    recipe = models.ForeignKey("Recipe", on_delete=models.CASCADE)
    text = models.TextField()

    class Meta:
        order_with_respect_to = "recipe"


class Recipe(models.Model):
    name = models.CharField(max_length=50)
