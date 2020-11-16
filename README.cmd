# Instruction
1. Put this program to your go-workspace/src folder
2. Get golang library by "go get lib-name" command
3. Run program by "run go main.go" command from gqlgen-sqlc-example folder
4. Open browser with http://localhost:8080 or with your address
5. Typing your query-command and query your data from postgresql database
    example :
    - query-member:
        query {
            members {
                id
                name
                age
                skills {
                id
                }
            }
        }
    - query-skills:
        query {
            skills {
                id
                name
                descs
            }
        }
* Note: You also could create, update, delete data from the database by this program 
