### Exercises:
* add another flag to enable verbose output, showing information like date/time
* add another flag to prevent displaying completed items
* update the custom usage function to include additional instructions on how to provide new tasks to the tool
* include test cases for the remaining options, such as `-complete`
* update the tests to use the `TODO_FILENAME` environment variable instead of hard-coding the test file name so that it doesn't cause conflicts with an existing file
* update the getTask() function allowing it to handle muliline input from STDIN. Each line shoul be a new task in the list