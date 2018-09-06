# 错误处理

为了规范API代码的错误处理， 我们定义了protoapi的错误处理规范

## 返回结果

* 结果处理一共分以下4种:

    1. **正常结果**: 调用API成功时返回包含所求数据的结果
    2. **异常结果**: 调用API时出现业务错误， 如: 用户名/密码错误，账号被屏蔽等, 返回包含错误内容的结果
    3. **常见异常**: 调用API时出现框架，中间件内置的错误, 如: API不存在, 输入参数不合要求等, 与具体业务无关
    4. **错误**: 程序运行时遇到的错误， 如: `TimeOutError`, `OutOfMemoryError`, `SegmentFaultError` 等, 与业务无关

* 在protoapi种, 我们把这4种结果分别用不同的 `HTTP status code`来区分:
    1. 正常结果 `Response`: 200
    2. 异常结果 `BizError`: 400
    3. 常见异常 `CommonError`: 420
    4. 错误 `Error`: 500

## proto定义

* 正常结果 与 异常结果 定义于使用者所写的proto文件中， 比如`test/example.proto`中:

    ```
    // 正常结果
    message HelloResponse {
        string hi = 1;  
    }
    // 异常结果， 可以自行定义
    message Error {
        ErrorCode code = 1;
        string message = 2;
    }
    ```

* 常见错误 则定义在`proto/protoapi_common.proto`中， 可被引用生成相关的错误处理代码
    * 例如`CommonError`我们定义了以下四种:

    ```
    message CommonError {
        GenericError genericError = 1; 
        AuthError authError = 2; //认证错误
        ValidateError validateError = 3; //验证错误
        BindError bindError = 4; // 参数错误
    }
    ```

## 前端处理

* 如上所述，我们定义了不同结果用不同的 `HTTP status code`来区分， 在生成的`helper.ts`里：
    ```
    /**
    * Defined Http Code for response handling
    */
    export enum httpCode {
        DEFAULT = 0,
        NORMAL = 200,
        BIZ_ERROR = 400,
        COMMON_ERROR = 420,
        INTERNAL_ERROR = 500,
    }
    ```
* 生成的ts代码中， 错误处理会使用一个switch， 根据error的status来决定返回不同的错误内容:


    ```
    switch(err.response.status){
        case httpCode.BIZ_ERROR:
            return Promise.reject(err.response.data.error  as Error)
        case httpCode.COMMON_ERROR:
            let returnErr = mapCommonErrorType(err.response.data.error);
            return Promise.reject(returnErr)
        case httpCode.INTERNAL_ERROR:
            return Promise.reject(err.response.data.error)
        default:
            return Promise.reject(new Error("Unknown Error"))
    }
    ```
* `mapCommonErrorType` 会根据返回的常见错误类型再做区分：

    ```
    for (let key in commonErr) {
        switch (key) {
            case 'genericError':
                return commonErr[key] as GenericError
            case 'authError':
                return commonErr[key] as AuthError

            case 'validateError':
                return commonErr[key] as ValidateError

            case 'bindError':
                return commonErr[key] as BindError
            default:
                return "Unknow Error"
        }
    }
    ```