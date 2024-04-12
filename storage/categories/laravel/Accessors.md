# Laravel Accessor 

Laravel Accessor are responsible to manipulate an attribute's value when its accessed.

Let's supppose you want to add a new field to your model when you are retrieving it.

In the follow example i want to add a field called 'image_url' to my Draw Model.

```php

protected $appends = ['image_url'];

public function getImageUrlAttribute()
{
    return Storage::url("imgs/" . $this->image)
}

```

the name of the function related to the new field appended must follow the laravel convetion to the accessor, so, if u are appending the image_url field u must name your function as getImageUrlAttribute.

name convention => get[AttributeName]Attribute
