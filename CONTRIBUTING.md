# Contributing to Tailscale
=====================================

Thank you for considering contributing to Tailscale, the easiest, most secure way to use WireGuard and 2FA. We appreciate your help in making Tailscale better.

## Fork Guide
-------------

To contribute to Tailscale, follow these steps:

1. Fork the Tailscale repository to your own GitHub account.
2. Clone the forked repository to your local machine using `git clone https://github.com/your-username/tailscale.git`.
3. Create a new branch for your feature or bug fix using `git checkout -b feature/your-feature-name`.
4. Make your changes and commit them using `git commit -m "feat: your commit message"` (see Conventional Commits section below).
5. Push your changes to your forked repository using `git push origin feature/your-feature-name`.
6. Open a pull request against the Tailscale repository.

## Branch Naming
----------------

We follow a consistent naming convention for our branches:

* `main`: The main branch where the latest production-ready code is stored.
* `feature/*`: Feature branches for new features or enhancements.
* `fix/*`: Fix branches for bug fixes.
* `docs/*`: Documentation branches for changes to the documentation.
* `refactor/*`: Refactor branches for code refactoring.
* `test/*`: Test branches for test additions or modifications.
* `chore/*`: Chore branches for miscellaneous changes.

## Conventional Commits
-----------------------

We follow the Conventional Commits specification for our commit messages. This helps us to automatically generate changelogs and determine the next version number.

* `feat`: New feature or enhancement.
* `fix`: Bug fix.
* `docs`: Documentation change.
* `refactor`: Code refactoring.
* `test`: Test addition or modification.
* `chore`: Miscellaneous change.
* `perf`: Performance improvement.
* `revert`: Revert a previous commit.
* `style`: Code style change.
* `ci`: Continuous integration change.
* `build`: Build system change.

Example commit message:
```
feat: add support for new authentication protocol
```

## PR Checklist
----------------

Before opening a pull request, make sure to:

* [ ] Follow the fork guide above.
* [ ] Create a feature or fix branch with a descriptive name.
* [ ] Write a clear and concise commit message following the Conventional Commits specification.
* [ ] Include a brief description of the changes in the pull request body.
* [ ] Ensure all tests pass and there are no linting errors.
* [ ] Review the code changes and make sure they align with the Tailscale code style.
* [ ] Test the changes manually to ensure they work as expected.

We appreciate your contribution to Tailscale and look forward to reviewing your pull request!