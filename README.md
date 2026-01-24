# FileNest

A virtual file system to store, organize and share your files.

## Development Workflow

This repository follows a simple promotion-based workflow:

- Feature and fix branches are merged into `dev`
- The `dev` branch is used for integration, deployment checks, and continuous delivery to staging
- Once staging is verified, `dev` is merged directly into `main`
- The `main` branch represents production and is deployed only after successful staging validation

This approach keeps production stable while allowing rapid iteration and validation in staging.
