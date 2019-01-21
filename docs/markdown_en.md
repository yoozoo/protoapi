# Markdown

Markdown support in protoapi

## Writing the proto file

* protoapi will take comments that are before or at the same line as a variable. It also supports multiple line comments, for example:

```protobuf
message Sample {
    string name = 1; //Account name
    /*
    an Id Comment
    */
    int32 aID = 2;
    //operator ID
	//test
    int32 operator = 3; // test operator
}

//nested message
message NestSample {
	string nest = 1;
}
```

* protoapi supports comments in different levels; message, field, enum, service and rpc level:

```protobuf
//message comment
message SampleResp {
    string msg = 1; //field comment
}

//title enum comment
enum Status {
    UNKNOWN = 0; //enum comment
    VIP_1 = 1;
    VIP_2 = 2;
}

//service comment
service OneService {
	//rpc comment
	rpc logingame(Sample) returns (SampleResp) {
	}
}
```

## Generated code:

**Note that there is a `-ROOT-` beside the definition of a parameter, this is to indicate the root definition of a parameter in case there is a nested parameter, such as `Sample` and `NestSample`**

* The generated md file would look like this:


# logingame

### A brief description:
- rpc comment  

### Request URL:
- 
```URL
    http://user.gtarcade.com/micro/OneService.logingame
```

### Request method:
- post

### parameter:

## Sample -ROOT- 
| parameter name  | required  | type  | description
| :-------------- |:--------- | :---- | :----------
|name        | required     | string  | Account name  
|aID        | required     | int  |  an Id Comment  
|operator        | required     | int  | operator ID test   test operator   

## NestSample  (nested message  )
| parameter name  | required  | type  | description
| :-------------- |:--------- | :---- | :----------
|nest        | required     | string  |  

### Successful Return Example:

```json
{
   "msg": "Success"
}
```

### Return parameter description

## SampleResp -ROOT- ( message comment  )
| parameter name  | type            | description
| :------------   |:--------------- | :----------
|msg        | string  | field comment  



### Enum description

## ValidateErrorType 
| field name  | value   | description
| :---------  |:------- | :----------
|INVALID_EMAIL        | 0 | 
|FIELD_REQUIRED        | 1 | 

## Status (title enum comment  )
| field name  | value   | description
| :---------  |:------- | :----------
|UNKNOWN        | 0 | enum comment  
|VIP_1        | 1 | 
|VIP_2        | 2 | 


### Remarks


- service comment  
