# Gates and Policies

Gates: 

- Gates are simple, closure-based authorization checks that can be defined anywhere within your application.
- They are typically used for basic authorization checks, such as determining whether a user can perform a 
specific action based on some condition.
- Gates can be defined in the AuthServiceProvider class using the Gate facade's define method.

```php
use Illuminate\Support\Facades\Gate;

Gate::define('update-post', function ($user, $post) {
    return $user->id === $post->user_id;
});
```

Policies: 

- Policies are classes that encapsulate authorization logic for a particular model or resource.
- They provide a convenient way to organize and manage authorization rules related to specific models.
- Each model typically has an associated policy class that defines authorization methods for various actions (e.g., view, create, update, delete).
- Policies are registered in the AuthServiceProvider class using the policy method.
- Example of registering a policy:

```php
use App\Models\Post;
use App\Policies\PostPolicy;

protected $policies = [
    Post::class => PostPolicy::class,
];
```

```php
namespace App\Policies;

use App\Models\User;
use App\Models\Post;

class PostPolicy
{
    public function update(User $user, Post $post)
    {
        return $user->id === $post->user_id;
    }
}
```