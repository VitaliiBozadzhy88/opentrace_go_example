# **README**

## 📀 Preparing to launch a project

1. Clone project (need to push Code button and choose in what way you want to get project)
2. Open with IntelliJ IDEA (if you do not have this program - download it [here](https://www.jetbrains.com/idea/download/#section=mac))
3. You will also need Docker to this project. Download it from [here](https://www.docker.com)
4. You have to SET UP ROOT. How to do it you find [here](https://www.jetbrains.com/help/idea/configuring-goroot-and-gopath.html)
5. For quick start you may open MySql and write Query -> `create database usersdb` then another one Query -> 
`create table Users ( id int auto_increment primary key, email varchar(30), activation_code varchar(30))`
and the last one ``insert into usersdb.Users (`email`, `activation_code`, `name`) values ('test@test.com', 'Test_Code_8976', 'TEST_Name')``


## 📌 How the project works

1. open terminal/cmd and put `docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest`
2. Find  [main.go](main.go) and press Run button or ^R in Intellij IDEA to start server
3. Open new tab in browser and type http://localhost:8080/getPerson
4. Fill in the field Email `test@test.com`(we add it before)
5. Open new tab in browser and type http://localhost:16686
6. Push button FIND TRACES. In Service menu find Get_email_service and push FIND TRACES
7. If you need more information about JAEGER visit [site](https://www.jaegertracing.io) and read documentation# opentrace_go_example
