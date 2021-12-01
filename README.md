#### Installation
add proper configurations in src/config/config.json file
->If you want to change the message legth in queue change src/config/config.json ->Kafka->MessageLength
->If you want to controll the dequeue then change src/config/config.json ->Kafka->RoutineCount


### Producer

run ->go run producer.go
Function Enqueue - which add the data into queue


### Consumer


run ->go run consumer.go

Function deQueueMessage - which add dequeu the message from queue

### error handiling
Jira mechanism is used for Error handling

add proper configurations in src/config/config.json file

  "Jira": {
                "Host": "sample.atlassian.net", //jira host
                "Username": "sample@test.com", //jira bug reporter username
                "ApiKey":  "Ra3k22TYUuLL4wI8lpbLHBQ6F9E", //API key 
                "ProjectKey" : "MC", //Project key, which project we need to assign the errors as the bug
                "AssigneeId" : "g34567890" // Bug assignee Id 
        }
 
        
        
        
        





	
