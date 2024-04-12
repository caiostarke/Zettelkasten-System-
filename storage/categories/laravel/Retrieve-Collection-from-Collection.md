# How To Retrieve a new Collection from a value in a Collection.

Let's suppose you have a collection with the likes of a post, and you want to retrieve the posts u have likes.

```php
$likes = Auth::user()->likes;

if ($likes) {
    $posts = Post::whereIn('id', likes->pluck('id'))->get();
}

return $posts;
```