/**
* 该文件生成于protoapi
* 文件包含前端调用API的代码，供微信小程序使用
* 文件内代码使用TypeScript
*/
import {
    EnvListRequest,
    EnvListResponse,
    KVHistoryRequest,
    KVHistoryResponse,
    KeyListRequest,
    KeyListResponse,
    KeyValueListRequest,
    KeyValueListResponse,
    KeyValueRequest,
    KeyValueResponse,
    ProductListRequest,
    ProductListResponse,
    RegisterServiceRequest,
    RegisterServiceResponse,
    SearchKeyValueListRequest,
    ServiceListRequest,
    ServiceListResponse,
    ServiceSearchRequest,
    TagListRequest,
    TagListResponse,
    UpdateServiceRequest,
    UpdateServiceResponse,
    UploadProtoFileRequest,
    UploadProtoFileResponse,
    
} from './AppServiceObjs';
import { generateUrl, errorHandling } from './helper';

const promisify = (wx) => {
  return (method) => {
    return (option) => {
      return new Promise ((resolve,reject) => {
        wx[method]({
          ...option,
          success:(res) => { resolve(res) },
          fail: (err) => {reject(err)}
        })
      })
    }
  }
}

const wxPromisify = promisify(wx)
const wxRequest = wxPromisify('request')

var baseUrl = "http://192.168.115.60:8080";

export function SetBaseUrl(url: string) {
    baseUrl = url;
}

var authToken = "token";

export function SetAuthToken(token: string) {
  authToken = token;
}

export function getEnv(params: EnvListRequest): Promise<EnvListResponse | never> {
  let url: string = generateUrl(baseUrl, "AppService", "getEnv");

  return wxRequest({ url: url, data: params, method:'POST', header:{'Authorization': 'token ' + authToken}}).then(res => {
    if (typeof res.data === 'object') {
      try {
        return Promise.resolve(res.data as EnvListResponse)
      } catch (e) {
        return Promise.reject(res.data);
      }
    }
    return Promise.reject(res.data);
  }).catch(err => {
    // handle error response
    return errorHandling(err)
  });
}


export function registerService(params: RegisterServiceRequest): Promise<RegisterServiceResponse | never> {
  let url: string = generateUrl(baseUrl, "AppService", "registerService");

  return wxRequest({ url: url, data: params, method:'POST', header:{'Authorization': 'token ' + authToken}}).then(res => {
    if (typeof res.data === 'object') {
      try {
        return Promise.resolve(res.data as RegisterServiceResponse)
      } catch (e) {
        return Promise.reject(res.data);
      }
    }
    return Promise.reject(res.data);
  }).catch(err => {
    // handle error response
    return errorHandling(err)
  });
}


export function updateService(params: UpdateServiceRequest): Promise<UpdateServiceResponse | never> {
  let url: string = generateUrl(baseUrl, "AppService", "updateService");

  return wxRequest({ url: url, data: params, method:'POST', header:{'Authorization': 'token ' + authToken}}).then(res => {
    if (typeof res.data === 'object') {
      try {
        return Promise.resolve(res.data as UpdateServiceResponse)
      } catch (e) {
        return Promise.reject(res.data);
      }
    }
    return Promise.reject(res.data);
  }).catch(err => {
    // handle error response
    return errorHandling(err)
  });
}


export function uploadProtoFile(params: UploadProtoFileRequest): Promise<UploadProtoFileResponse | never> {
  let url: string = generateUrl(baseUrl, "AppService", "uploadProtoFile");

  return wxRequest({ url: url, data: params, method:'POST', header:{'Authorization': 'token ' + authToken}}).then(res => {
    if (typeof res.data === 'object') {
      try {
        return Promise.resolve(res.data as UploadProtoFileResponse)
      } catch (e) {
        return Promise.reject(res.data);
      }
    }
    return Promise.reject(res.data);
  }).catch(err => {
    // handle error response
    return errorHandling(err)
  });
}


export function getTags(params: TagListRequest): Promise<TagListResponse | never> {
  let url: string = generateUrl(baseUrl, "AppService", "getTags");

  return wxRequest({ url: url, data: params, method:'POST', header:{'Authorization': 'token ' + authToken}}).then(res => {
    if (typeof res.data === 'object') {
      try {
        return Promise.resolve(res.data as TagListResponse)
      } catch (e) {
        return Promise.reject(res.data);
      }
    }
    return Promise.reject(res.data);
  }).catch(err => {
    // handle error response
    return errorHandling(err)
  });
}


export function getProducts(params: ProductListRequest): Promise<ProductListResponse | never> {
  let url: string = generateUrl(baseUrl, "AppService", "getProducts");

  return wxRequest({ url: url, data: params, method:'POST', header:{'Authorization': 'token ' + authToken}}).then(res => {
    if (typeof res.data === 'object') {
      try {
        return Promise.resolve(res.data as ProductListResponse)
      } catch (e) {
        return Promise.reject(res.data);
      }
    }
    return Promise.reject(res.data);
  }).catch(err => {
    // handle error response
    return errorHandling(err)
  });
}


export function getServices(params: ServiceListRequest): Promise<ServiceListResponse | never> {
  let url: string = generateUrl(baseUrl, "AppService", "getServices");

  return wxRequest({ url: url, data: params, method:'POST', header:{'Authorization': 'token ' + authToken}}).then(res => {
    if (typeof res.data === 'object') {
      try {
        return Promise.resolve(res.data as ServiceListResponse)
      } catch (e) {
        return Promise.reject(res.data);
      }
    }
    return Promise.reject(res.data);
  }).catch(err => {
    // handle error response
    return errorHandling(err)
  });
}


export function searchServices(params: ServiceSearchRequest): Promise<ServiceListResponse | never> {
  let url: string = generateUrl(baseUrl, "AppService", "searchServices");

  return wxRequest({ url: url, data: params, method:'POST', header:{'Authorization': 'token ' + authToken}}).then(res => {
    if (typeof res.data === 'object') {
      try {
        return Promise.resolve(res.data as ServiceListResponse)
      } catch (e) {
        return Promise.reject(res.data);
      }
    }
    return Promise.reject(res.data);
  }).catch(err => {
    // handle error response
    return errorHandling(err)
  });
}


export function getKeyList(params: KeyListRequest): Promise<KeyListResponse | never> {
  let url: string = generateUrl(baseUrl, "AppService", "getKeyList");

  return wxRequest({ url: url, data: params, method:'POST', header:{'Authorization': 'token ' + authToken}}).then(res => {
    if (typeof res.data === 'object') {
      try {
        return Promise.resolve(res.data as KeyListResponse)
      } catch (e) {
        return Promise.reject(res.data);
      }
    }
    return Promise.reject(res.data);
  }).catch(err => {
    // handle error response
    return errorHandling(err)
  });
}


export function getKeyValueList(params: KeyValueListRequest): Promise<KeyValueListResponse | never> {
  let url: string = generateUrl(baseUrl, "AppService", "getKeyValueList");

  return wxRequest({ url: url, data: params, method:'POST', header:{'Authorization': 'token ' + authToken}}).then(res => {
    if (typeof res.data === 'object') {
      try {
        return Promise.resolve(res.data as KeyValueListResponse)
      } catch (e) {
        return Promise.reject(res.data);
      }
    }
    return Promise.reject(res.data);
  }).catch(err => {
    // handle error response
    return errorHandling(err)
  });
}


export function searchKeyValueList(params: SearchKeyValueListRequest): Promise<KeyValueListResponse | never> {
  let url: string = generateUrl(baseUrl, "AppService", "searchKeyValueList");

  return wxRequest({ url: url, data: params, method:'POST', header:{'Authorization': 'token ' + authToken}}).then(res => {
    if (typeof res.data === 'object') {
      try {
        return Promise.resolve(res.data as KeyValueListResponse)
      } catch (e) {
        return Promise.reject(res.data);
      }
    }
    return Promise.reject(res.data);
  }).catch(err => {
    // handle error response
    return errorHandling(err)
  });
}


export function updateKeyValue(params: KeyValueRequest): Promise<KeyValueResponse | never> {
  let url: string = generateUrl(baseUrl, "AppService", "updateKeyValue");

  return wxRequest({ url: url, data: params, method:'POST', header:{'Authorization': 'token ' + authToken}}).then(res => {
    if (typeof res.data === 'object') {
      try {
        return Promise.resolve(res.data as KeyValueResponse)
      } catch (e) {
        return Promise.reject(res.data);
      }
    }
    return Promise.reject(res.data);
  }).catch(err => {
    // handle error response
    return errorHandling(err)
  });
}


export function fetchKeyHistory(params: KVHistoryRequest): Promise<KVHistoryResponse | never> {
  let url: string = generateUrl(baseUrl, "AppService", "fetchKeyHistory");

  return wxRequest({ url: url, data: params, method:'POST', header:{'Authorization': 'token ' + authToken}}).then(res => {
    if (typeof res.data === 'object') {
      try {
        return Promise.resolve(res.data as KVHistoryResponse)
      } catch (e) {
        return Promise.reject(res.data);
      }
    }
    return Promise.reject(res.data);
  }).catch(err => {
    // handle error response
    return errorHandling(err)
  });
}
