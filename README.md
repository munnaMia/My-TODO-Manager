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
    - *add* : Use add command to add task.
    - *-t* : Use -t flag after add command to add a title for the task.
    - *-d* : Use -d flag to add description. Description is optional.
    ```bash
    ./todo add -t "Title here" 
    ```
    OR
    ```bash
    ./todo add -t "Title here" -d "Description here"
    ```
<br>

2. **COMPLETE TASK**
    - *complete* : Use this command to mark a task when it's done.
    - *-id* : Use -id (id number) flag to mark a task that complete.
    ```bash
    ./todo complete -id <id>
    ```
<br>

3. **SEE TASKS**
    - *list* : Use list command to see all tasks.
    - *-p* : Use -p flag to see all the pending task.
    - *-c* : Use -p flag to see all the completed task.
    
    **See all the pending task**
    ```bash
    ./todo list -p
    ```
    **See all the completed task**
    ```bash
    ./todo list -c
    ```
<br>

4. **DELETE TASK**
    - *delete* : Use this command to delete tasks
    - *-id* : Use this -id (id) flag to delete a single pending task.
    - *-all* : Use this -all flag to delete all pending tasks
    - *-c* : Use this -c flag to delete all completed tasks.

    **Delete a single task**
    ```bash
    ./todo delete -id <id>
    ```
    **Delete all pending tasks**
    ```bash
    ./todo delete -all
    ```
    **Delete all complete tasks**
    ```bash
    ./todo delete -c
    ```
<br>

4. **HELP**
    - *help* : Use this command to see all the commands

    ```bash
    ./todo help
    ```
<br>
<br>



## Packages are used in this project :
- os package.
- flag package
- fmt package
- path/filepath package
- time
- encoding/json
- text/tabwriter

<br>
<br>

**This is a simple practise project to understand how CLI apps are work.**
