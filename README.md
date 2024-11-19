# My-TODO-Manager

## Features :
- Add Task.
- Mark completed tasks.
- See list 
    - See list of pending tasks.
    - See list of complete tasks.
- Delete task.
    - Delete a single tasks.
    - Delete all tasks.
    - Delete all complete task.
- Data store into *JSON* file.

## Requirements
- Go **1.21** or later.
- Basic knowledge of the Go *os* and *flag* packages (for contributors).


## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/munnaMia/My-TODO-Manager.git
    ```

2.  Navigate to the project directory:
    ```bash
    cd My-TODO-Manager
    ```

3.  Build the application:
    ```bash
    go build -o todo
    ```


## Usage
Run the application using the following commands:

1. **ADD TASK**
    - add : Use add command to add task.
    - -t : Use -t flag after add command to add a title for the task.
    - -d : Use -d flag to add description. Description is optional.
    ```bash
    ./todo add -t "Title here" 
    ```
    OR
    ```bash
    ./todo add -t "Title here" -d "Description here"
    ```
   
   



