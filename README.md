# Documentation

## 0. Requirements

Before running the project, ensure you have the following installed:

- **Docker**: You need to have Docker installed on your system. Follow the official installation guide [here](https://docs.docker.com/get-docker/).
- **Task**: A task runner to simplify execution. Install it by following the guide [here](https://taskfile.dev/).

## 1. How to Start

To start the project, execute the following commands:

- **Start Docker Services**
  ```sh
  task docker
  ```
  This command will start the required services in a Docker container. Once running, you can access the Swagger documentation at:
  [http://localhost:8080/swagger/](http://localhost:8080/swagger/).

- **Build the CLI Tool**
  ```sh
  task build-cli
  ```
  This will compile the solution, generating an executable called `quiz` in the current directory.

## 2. How to Use

Once the `quiz` CLI tool is built, it provides three main functionalities:

1. **Retrieve All Questions**
   ```sh
   quiz questions
   ```
   This command fetches all available survey questions.

2. **Retrieve a Specific Question**
   ```sh
   quiz question X
   ```
   Replace `X` with any question ID obtained from the `quiz questions` command.

3. **Evaluate Answers**
   ```sh
   quiz evaluate "A B C D"
   ```
   Here, `A B C D` represents the selected answers in sequence for questions 1, 2, 3, etc.

## 3. Architecture

The project follows a **Clean Code** architecture, ensuring:

- **Separation of concerns**: Code is organized in a way that each module has a single responsibility.
- **Maintainability**: Easy to extend and modify without affecting unrelated parts of the system.
- **Testability**: Facilitates unit testing with clearly defined modules and interfaces.

By following these principles, the code remains structured, readable, and scalable for future enhancements.

