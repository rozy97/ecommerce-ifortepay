<h1 align="center"> 
    Ecommerce iFortepay
</h1>


This is an API-only Ecommerce App for iFortepay Backend Technical test. This app build using Go and MySQL. The content of this documentation includes:

- How to run the application locally
- How to manual test the API
- How to run unit test
- Entity Relationship Diagram
- Sequence Diagram for every feature

___

<h3>
    How to run the application locally 
</h3>

You can run the application using docker compose command below:
```
docker-compose up
```

___

<h3>
    How to manual test the API
</h3>

Open swagger UI at 123.123.123.123/docs or after run the app locally, access localhost:8080/docs
<p align="center">
<img src="https://static1.smartbear.co/swagger/media/images/tools/opensource/swagger_ui.png"/>
</p>

___

<h3>
    How to run unit test
</h3>

```
go test -v ./... -cover
```

___

<h3>
    Entity Relationship Diagram
</h3>

<p align="center">
<img src="https://raw.githubusercontent.com/rozy97/ecommerce-ifortepay/refs/heads/main/assets/erd.png"/>
</p>

___

<h3>
    Sequence Diagram
</h3>

<p align="center">
Register Feature
<img src="https://raw.githubusercontent.com/rozy97/ecommerce-ifortepay/refs/heads/main/assets/register.png"/>
</p>

<br/>
