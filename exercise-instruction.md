1. Take the interface definition file (swagger.yaml) and cut-paste its contents into https://editor.swagger.io 
2. If the interface didn’t make much sense to you in its text form, explore it in this tool. It’s more human friendly. 
3. Generate some starter code with Generate Server -> go-server. That creates a go-server-server-generated.zip file. 
4. Extract the file into your Go environment. 
5. Write your loan calculator microservice in Go, using the generated code as a starting point. Start with one that returns dummy data. Get it running locally and test with Postman. 
6. Complete the microservice to calculate the required results. Test with Postman using a range of values. 
7. Create a GitHub Action that triggers a build for each push to the main branch. 
8. If you haven’t already, add unit tests to your microservice to achieve a test coverage of at least 80%. 
9. Upgrade your GitHub Action to run tests before building. 
10. Add a linter to your project. 
11. Run linter and resolve all issues. 
12. Upgrade your GitHub Action to run linter before running tests. 
13. Upgrade GitHub Action to deploy microservice to Google App Engine. Test with Postman.

