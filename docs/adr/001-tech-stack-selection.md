## Status
Accepted

## Context
The Tailscale project aims to provide a secure and easy-to-use solution for WireGuard and 2FA. To achieve this, we need to decide on a programming language for the project. After considering factors such as performance, concurrency, and simplicity, we are choosing between Go and other languages.

## Decision
We will use Go as the primary programming language for the Tailscale project. Go's simplicity, concurrency features, and performance make it an ideal choice for building a secure and efficient networking solution.

## Consequences
Using Go for the Tailscale project will have the following consequences:
* **Improved performance**: Go's lightweight goroutine scheduling and concurrency features will enable us to build a high-performance networking solution.
* **Simplified development**: Go's simplicity and minimalistic design will reduce the complexity of our codebase and make it easier for developers to contribute to the project.
* **Better security**: Go's built-in concurrency features and performance will enable us to build a more secure solution, reducing the risk of vulnerabilities and exploits.