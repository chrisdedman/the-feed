
[![Go](https://github.com/chrisdedman/the-feed/actions/workflows/go.yml/badge.svg)](https://github.com/chrisdedman/the-feed/actions/workflows/go.yml)

# THE FEED

  At RSS Feed Network, we're on a mission to empower you with the latest, most reliable news from around the globe. We believe in putting you in control of your newsfeed, free from the noise of sponsored content and any other unwelcome noises.

  With our platform, you have the power to curate your own personalized feeds by subscribing to RSS feeds from your favorite space agencies, scientific journals, STEM influencers, and more. Say goodbye to cluttered timelines and hello to a tailored news experience that matters to you.

  But we're more than just a news aggregator. We're a community-driven platform that fosters meaningful interactions and discussions. Share your latest RSS feed discoveries, engage in thoughtful conversations, and connect with like-minded individuals who share your interests.

  We take your privacy and data security seriously. That's why we're committed to ethical data practices and transparent policies that put you in control of your personal information.

  Join us in our mission to explore and engage with RSS-based content in a spam-free environment. Together, let's reclaim ownership of the content that we see and empower ourselves with knowledge.


### Table of Content
1. [How to set up the project](#whats-needed-to-run-the-server)
2. [How to run the server](#how-to-run-the-server-manual-setup)
3. [How to set up Docker](#how-to-set-up-docker-recommended)
4. [Technologies](#technologies)
5. [License](#license)
6. [Demo](#demo)

## What's needed to run the server
- [Golang](https://golang.org/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker-Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/) (if you don't use Docker)
- [PostgreSQL](https://www.postgresql.org/) (if you don't use Docker)
- ``.env`` file

In the root directory, create a new file named ``.env``.<br>
Copy the content of ``.env.example`` into your ``.env`` file.

Replace the placeholders for the database connection with your own values.

## How to run the server (manual setup)
1. Run the following commands on your terminal to clone the repository and run the server:
    ```bash
    git clone https://github.com/chrisdedman/the-feed.git
    cd the-feed       # Change directory to the project folder
    make build        # Build the server (required Make)
    ./build/server && ./build/database # Start the project
    ```
2. Open your browser and navigate to `http://localhost:3000` (or any other port you specified in the `.env` file)

## How to set up Docker (Recommended)

To run the server using Docker, you need to have Docker and Docker-Compose installed on your machine. Follow these steps to set up Docker:

1. Install **Docker** and **Docker-Compose** by following the instructions for your operating system from the official Docker website: 
    - [Docker](https://docs.docker.com/get-docker/)
    - [Docker-Compose](https://docs.docker.com/compose/install/).
2. Once installed, open a terminal or command prompt and navigate to the root directory of the project.

3. Run the following command to start the server using Docker:

   ```bash
   docker-compose up
   ```

4. The server will start running in a Docker container. You can access the server by opening a web browser and navigating to `http://localhost:3000` (or any other port you specified in the `.env` file).

5. To stop the server, press `Ctrl+C` in the terminal where the server is running, and then run the following command to stop and remove the Docker container (if needed):

   ```bash
   docker-compose down
   ```

## Technologies
- [Golang](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Docker-Compose](https://docs.docker.com/compose/)
- [GORM](https://gorm.io/)
- [Gin Web Framework](https://pkg.go.dev/github.com/gin-gonic/gin#section-readme)
- [PostgreSQL](https://www.postgresql.org/)
- [Make](https://www.gnu.org/software/make/)

## License
This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

## Demo
![Demo](/assets/demo.gif)
