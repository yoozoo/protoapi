# Markdown

protoapi支持生成markdown文档。

## 编辑proto文件

* protoapi会把注释作为对对应元素的描述

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

* protoapi支持多个层级的注释

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

## 生成的文档:

**在某些参数后面有 `-ROOT-` 字样, 这表示这是其它参数的root.**

* 生成的文档会是如下样子：


    # logingame

    ### 简要描述：
    - test method comment

    ### 请求URL：
    - `LoginService.logingame`

    ### 请求方式：
    - post

    ### 参数：

    ## LoginReq -ROOT-
    | parameter name  | required  | type  | description
    | :-------------- |:--------- | :---- | :----------
    |account        | required     | string  | Account name
    |game_id        | required     | int  |  game id refer to game table
    |op_id        | required     | int  | operation Id   test table
    |server_id        | required     | int  |


    ### 返回示例：

    ```json
    {
    "code": "0",
    "data": "Success",
    "loginInfo": {
        "accountInfo": {
            "account_name": "Success",
            "join_time": "Success"
        },
        "is_vip": false,
        "last_login": "Success",
        "user_id": "Success",
        "vipStatus": "UNKNOWN"
    },
    "msg": "Success"
    }
    ```

    ### 返回参数说明：

    ## LoginResp -ROOT- ( login request return  )
    | parameter name  | type            | description
    | :------------   |:--------------- | :----------
    |code        | int  |
    |msg        | string  |
    |loginInfo        | LoginInfo  |
    |data        | string Array |

    ## LoginInfo
    | parameter name  | type            | description
    | :------------   |:--------------- | :----------
    |user_id        | string  |
    |last_login        | string  |
    |is_vip        | bool  |
    |accountInfo        | AccountInfo Array |
    |vipStatus        | VipStatus  | this is an enum

    ## AccountInfo
    | parameter name  | type            | description
    | :------------   |:--------------- | :----------
    |account_name        | string  |
    |join_time        | string  |



    ### Enum说明：

    ## ValidateErrorType
    | field name  | value   | description
    | :---------  |:------- | :----------
    |INVALID_EMAIL        | 0 |
    |FIELD_REQUIRED        | 1 |

    ## VipStatus
    | field name  | value   | description
    | :---------  |:------- | :----------
    |UNKNOWN        | 0 | default value for status
    |VIP_1        | 1 |
    |VIP_2        | 2 |


    ### 备注


    - test service comment
