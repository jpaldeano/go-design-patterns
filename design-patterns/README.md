# Design Patterns in relation to SOLID

This structure highlights both the strengths and potential pitfalls of each pattern in relation to SOLID principles.
Please refer to the code in the respective folders.

## Summary table

| **Pattern**         | **Use Case**                                                                                          | **Where It Complies**                                                                                                                                                         | **Where It May Not Comply**                                                                                                   |
|----------------------|------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------|
| **Factory**         | Centralizes object creation and abstracts away concrete class dependencies.                          | - **SRP:** Keeps creation logic separate.<br>- **DIP:** Abstracts dependencies through interfaces.                                                                            | - **OCP:** Adding new types may require changes to the factory itself, violating extensibility.                             |
| **Strategy**        | Allows dynamic selection or change of behavior (algorithm) at runtime.                               | - **OCP:** Easily extendable with new strategies.<br>- **LSP:** Behavior is consistent via abstraction.<br>- **DIP:** Depends on abstractions, not implementations.           | - **SRP:** If strategy implementations become complex, they may contain logic violating single responsibility.              |
| **Decorator**       | Adds functionality (e.g., logging, caching) dynamically without modifying the original component.     | - **OCP:** Adds behavior without altering the original component.<br>- **SRP:** Separates base behavior and added functionality.                                             | - **DIP:** If decorators tightly couple to the component, it violates dependency inversion.                                 |
| **Singleton**       | Ensures a class has a single, globally accessible instance (e.g., configuration, logging).           | - **SRP:** Provides a single point of control for instance creation.                                                                                                         | - **DIP:** Encourages global state, increasing tight coupling.<br>- **SRP:** If singleton takes on additional responsibilities. |
| **Builder**         | Constructs complex objects step by step, useful for configurations or object hierarchies.            | - **SRP:** Isolates object construction logic.<br>- **OCP:** Allows variations in object configurations without modifying the client.                                         | - **OCP:** Changes in object structure may require updating the builder class.                                               |
| **Observer**        | Establishes event-driven communication between objects (e.g., pub-sub systems, notifications).       | - **OCP:** Listeners can be added without modifying the subject.<br>- **SRP:** Separates event producer and consumers.                                                       | - **DIP:** Complex observer relationships may increase indirect coupling between subjects and observers.                     |
| **Adapter**         | Bridges incompatible interfaces, allowing reuse of existing classes.                                 | - **OCP:** New adapters can be introduced without modifying clients or adaptees.<br>- **SRP:** Keeps interface conversion logic separate.                                    | - **DIP:** Adapter may introduce tight coupling if not properly abstracted.                                                  |
| **Proxy**           | Controls access to an object, adding behavior like security or lazy initialization.                  | - **SRP:** Separates additional access-related concerns from the primary component.<br>- **OCP:** New proxies can be added without changing client code.                     | - **DIP:** Proxy may tightly couple access logic to the real component.                                                      |
| **Composite**       | Treats individual objects and object groups uniformly (e.g., file system hierarchy).                 | - **LSP:** Components and composites are interchangeable.<br>- **OCP:** New composite types can be added without changing the client code.                                   | - **SRP:** Complex hierarchies may violate single responsibility in composite implementation.                                |
| **Command**         | Encapsulates requests as objects, enabling undo/redo or deferred execution.                          | - **SRP:** Encapsulates action logic.<br>- **OCP:** Adding new commands doesn’t modify existing logic.<br>- **DIP:** Loose coupling between sender and receiver.             | - **SRP:** Complex commands may end up violating single responsibility by containing too much business logic.                |
| **State**           | Allows an object to change its behavior when its internal state changes (e.g., state machines).       | - **OCP:** New states don’t require changes to the client.<br>- **SRP:** Isolates state transitions into individual classes.                                                  | - **DIP:** Poor design can lead to state-specific classes becoming tightly coupled to the context object.                    |
| **Template Method** | Defines a skeleton algorithm, allowing subclasses to redefine parts without altering structure.       | - **OCP:** Algorithm extension is achieved via subclassing.<br>- **LSP:** Subclasses conform to base class behavior.                                                         | - **SRP:** The base class may violate single responsibility if the algorithm grows too complex.                              |


## Factory pattern

### Problems with the anti-pattern:
- Tightly Coupled Code: The `main` function directly depends on the concrete implementations (`GasOven`, `ElectricOven`, `WoodOven`), violating the dependency inversion principle.
- Limited Extensibility: Adding a new type of oven requires modifying the `main` function. This approach does not follow the Open/Closed Principle (code should be open for extension but closed for modification).
- Code Duplication: The `Bake()` method call is repeated across the different types, leading to verbosity.
- Scalability Issues: As the logic grows, this approach becomes harder to manage.

### Benefits of Using Factory:
- Encapsulation: The creation logic is encapsulated in the `OvenFactory`, removing it from the `main` function.
- Extensibility: Adding a new oven type only requires changes to the factory, not the client code.
- Loose Coupling: The `main` function interacts with the `Oven` interface rather than specific implementations.
- Readability: The `main` function becomes cleaner and easier to understand.

### Pattern pitfall
- OCP: Adding new types may require changes to the factory itself, violating extensibility.

## Strategy pattern

### Problems with the anti-pattern:
- Violation of SRP: The DecorateTree function handles both the logic for deciding the decoration style and the actual decoration process.
- Violation of Open/Closed Principle: Adding a new decoration style requires modifying the DecorateTree function, which increases the chance of breaking existing functionality.
- Scalability Issues: As the number of decoration styles grows, the switch statement becomes harder to maintain.

### Benefits of using the strategy:
- Adherence to SRP: Each decoration style has its own class, responsible only for its specific behavior. The TreeContext handles only the execution of the selected strategy.
- Open/Closed Principle: Adding a new style (e.g., "vintage") doesn’t require modifying existing code. Just implement the TreeDecorator interface.
- Flexible Runtime Behavior: The decoration style can be changed at runtime by setting a different decorator.
- Scalability: Adding new decoration styles is easy and doesn’t increase the complexity of the TreeContext.


### Pattern pitfall
- SRP: If strategy implementations become complex, they may contain logic violating single responsibility.
