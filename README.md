# ./app
Folder with business logic only. This directory doesn't care about what database driver you're using or which caching solution your choose or any third-party things.

./app/controllers folder for functional controllers (used in routes)
./app/models folder for describe business models of your project
./app/queries folder for describe queries for models of your project


# ./docs
Folder with API Documentation. This directory contains config files for auto-generated API Docs by OpenAPI.

# ./pkg
Folder with project specific functionality. This directory contains all the project-specific code tailored only for your business use case, like configs, middleware, routes, utils or else.

./pkg/configs folder for configuration functions
./pkg/middleware folder for add middleware (Fiber and yours)
./pkg/routes folder for describe routes of your project
./pkg/repository folder for describe const of your project
./pkg/utils folder with utility functions (server starter, error checker, etc)

# ./platform
Folder with platform-level logic. This directory contains all the platform-level logic that will build up the actual project, like setting up the database or cache server instance and storing migrations.

./platform/cache folder with in-memory cache setup functions
./platform/database folder with database configuration
./platform/migrations folder with migration files (used with golang-migrate/migrate tool)
