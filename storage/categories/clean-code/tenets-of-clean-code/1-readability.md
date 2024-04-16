# Readability

The cabability to make your code readable and understandable and self explainable.
In other words, readability needs to be reliabe.

For Example:

```javascript
isValidEmail('9VlT3@example.com') // true
isValidEmail('9VlT32_4+3424@example.com') // true
isValidEmail('9VlT32_4+3424@example.$$$') // false
isValidEmail('This is not valid email') // false
```

if we read the function isValidEmail we suppose to think it validates an email address.