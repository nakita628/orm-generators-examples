"""Every way a Django model refers to another."""

from django.conf import settings
from django.contrib.contenttypes.fields import GenericForeignKey, GenericRelation
from django.contrib.contenttypes.models import ContentType
from django.db import models


def default_country():
    return "JP"


class Country(models.Model):
    code = models.CharField(max_length=2, unique=True)
    name = models.CharField(max_length=50)


class City(models.Model):
    # A foreign key to a unique column other than the primary key
    country = models.ForeignKey(Country, to_field="code", on_delete=models.PROTECT, default=default_country)
    name = models.CharField(max_length=50)


class Person(models.Model):
    name = models.CharField(max_length=50)
    user = models.OneToOneField(settings.AUTH_USER_MODEL, on_delete=models.CASCADE, null=True, related_name="person")
    # Every on_delete
    city = models.ForeignKey(City, on_delete=models.SET_NULL, null=True, blank=True, related_name="people")
    born_in = models.ForeignKey(City, on_delete=models.RESTRICT, null=True, related_name="+")
    works_in = models.ForeignKey(City, on_delete=models.SET_DEFAULT, default=None, null=True, related_name="workers")
    lives_near = models.ForeignKey(City, on_delete=models.DO_NOTHING, null=True, db_constraint=False, related_name="neighbours")
    mentor = models.ForeignKey("self", on_delete=models.SET(None), null=True, blank=True, related_name="mentees")
    # Many-to-many: symmetrical to itself, one-way to itself, and through a model of its own
    friends = models.ManyToManyField("self", blank=True)
    follows = models.ManyToManyField("self", symmetrical=False, related_name="followers", blank=True)
    groups = models.ManyToManyField("Group", through="Membership", through_fields=("person", "group"), related_name="members")
    staff_only = models.ForeignKey(
        settings.AUTH_USER_MODEL, on_delete=models.CASCADE, null=True, limit_choices_to={"is_staff": True}, related_name="managed_people"
    )
    tags = GenericRelation("TaggedItem", related_query_name="person")


class Group(models.Model):
    name = models.CharField(max_length=50)
    tags = GenericRelation("TaggedItem")


class Membership(models.Model):
    person = models.ForeignKey(Person, on_delete=models.CASCADE)
    group = models.ForeignKey(Group, on_delete=models.CASCADE)
    inviter = models.ForeignKey(Person, on_delete=models.SET_NULL, null=True, related_name="invitations")
    joined = models.DateField()
    role = models.CharField(max_length=20, default="member")

    class Meta:
        unique_together = [("person", "group")]


class TaggedItem(models.Model):
    tag = models.SlugField()
    content_type = models.ForeignKey(ContentType, on_delete=models.CASCADE)
    object_id = models.PositiveBigIntegerField()
    content_object = GenericForeignKey("content_type", "object_id")

    class Meta:
        indexes = [models.Index(fields=["content_type", "object_id"])]


class Order(models.Model):
    pk = models.CompositePrimaryKey("store", "number")
    store = models.ForeignKey(City, on_delete=models.CASCADE)
    number = models.PositiveIntegerField()


class Shipment(models.Model):
    store = models.ForeignKey(City, on_delete=models.CASCADE, related_name="+")
    number = models.PositiveIntegerField()
    order = models.ForeignObject(Order, on_delete=models.CASCADE, from_fields=("store", "number"), to_fields=("store", "number"))
