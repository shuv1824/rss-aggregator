# RSS Aggregator

This project is a web server and RSS feed aggregator written in Go. It enables users to subscribe to RSS feeds and aggregates posts from those feeds. It leverages Go's concurrency features to efficiently scrape and update feeds at regular intervals.

## Features

- **User Management**: Create and manage users.
- **Feed Management**: Subscribe to RSS feeds and manage them.
- **Post Aggregation**: Collect posts from various RSS feeds and serve them to users.
- **Concurrency**: Scrape feeds concurrently for improved performance.
- **RESTful API**: Exposes a set of RESTful endpoints to interact with users, feeds, and posts.
- **CORS Support**: Configured to allow cross-origin requests to the API.

## Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/shuv1824/rss-aggregator.git
   cd rss-aggregator
   ```

2. **Dependencies**

   Install the dependencies using Go:

   ```bash
   go mod download
   ```

3. **Configure Environment**

   Create a `.env` file based on the template provided and fill in the required information:

   ```plaintext
   PORT=8080
   DB_URL=your_database_url_here
   ```

4. **Run Database Migrations**

   Use tools like `goose` to apply database migrations located in the `db/schema` directory.

5. **Build and Run**

   ```bash
   make
   ```

## API Endpoints

The following endpoints are available:

- **GET /v1/healthz**: Health check for the service.
- **POST /v1/users**: Create a new user.
- **GET /v1/users**: Fetch information about the authenticated user.
- **POST /v1/feeds**: Add a new RSS feed.
- **GET /v1/feeds**: Retrieve all subscribed feeds.
- **POST /v1/feed_follows**: Follow a feed.
- **GET /v1/feed_follows**: Get all followed feeds.
- **DELETE /v1/feed_follows/{feedFollowID}**: Unfollow a feed.
- **GET /v1/posts**: Retrieve posts for the authenticated user.

## Architecture

- The project is structured to keep domain logic and database abstractions cleanly separated.
- API handlers manage HTTP requests and responses.
- Database operations are handled via `sqlc`-generated methods, providing type-safe, efficient database access.
- The directory layout follows a common Go structure with `cmd`, `internal`, and `pkg` directories.

## Contributing

Contributions are welcome! Please submit a pull request or create an issue for any specific feature requests or bug reports.
