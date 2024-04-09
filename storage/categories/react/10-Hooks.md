# Hooks

Hooks are functions that allow you to use state and other React features in functional components.
Example of hooks are: useState, useEffect, useContext

[Doc Referece - React Hooks](https://reactjs.org/docs/hooks-reference.html)

## Example

```jsx
function App() {
    const [count, setCount] = useState(0);

    return (
        <div>
            <p>You clicked {count} times</p>
            <button onClick={() => setCount(count + 1)}>
                Click me
            </button>
        </div>
    );

}
```

[<- Previous](9-Lifecycle-Methods.md) [Next ->](11-Hooks.md)
    