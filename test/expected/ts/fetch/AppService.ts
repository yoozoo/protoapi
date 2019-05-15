/**
* This file is generated by 'protoapi'
* The file contains frontend API code that work with fetch API for HTTP usages
* The generated code is written in TypeScript
* The code provides a basic usage for API call and may need adjustment according to specific project requirement and situation
* -------------------------------------------
* 该文件生成于protoapi
* 文件包含前端调用API的代码，并使用fetch做HTTP调用
* 文件内代码使用TypeScript
* 该生成文件只提供前端API调用基本代码，实际情况可能需要根据具体项目具体要求不同而作出更改
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

let baseUrl = "backend";
const headers = {
    "X-Requested-With": "XMLHttpRequest"
};

export function SetBaseUrl(url: string) {
    baseUrl = url;
}

export function setHeader(header: {[key: string]: string}) {
    return Object.assign(headers, header);
}// use fetch
async function call<InType, OutType>(
    service: string,
    method: string,
    params: InType
): Promise<OutType | never> {
    const url = generateUrl(baseUrl, service, method);
    try {
        const fetchResolve = await fetch(url, {
            method: 'POST',
            body: JSON.stringify(params),
            headers
        });

        const resolvedData = await fetchResolve.json();

        if (fetchResolve.statusText !== 'OK') {
            const parsedError = {
                headers: fetchResolve.headers,
                type: fetchResolve.type,
                statusText: fetchResolve.statusText,
                status: fetchResolve.status,
                ok: fetchResolve.ok,
                redirected: fetchResolve.redirected,
                url: fetchResolve.url,
                data: resolvedData,
            }
            throw (parsedError);
        }
        return resolvedData as OutType;
    }
    catch (err) {
        const handledError = await errorHandling(err);
        throw handledError;
    }
}
export function getEnv(params: EnvListRequest): Promise<EnvListResponse | never> {
    return call<EnvListRequest, EnvListResponse>("AppService", "getEnv", params);
}

export function registerService(params: RegisterServiceRequest): Promise<RegisterServiceResponse | never> {
    return call<RegisterServiceRequest, RegisterServiceResponse>("AppService", "registerService", params);
}

export function updateService(params: UpdateServiceRequest): Promise<UpdateServiceResponse | never> {
    return call<UpdateServiceRequest, UpdateServiceResponse>("AppService", "updateService", params);
}

export function uploadProtoFile(params: UploadProtoFileRequest): Promise<UploadProtoFileResponse | never> {
    return call<UploadProtoFileRequest, UploadProtoFileResponse>("AppService", "uploadProtoFile", params);
}

export function getTags(params: TagListRequest): Promise<TagListResponse | never> {
    return call<TagListRequest, TagListResponse>("AppService", "getTags", params);
}

export function getProducts(params: ProductListRequest): Promise<ProductListResponse | never> {
    return call<ProductListRequest, ProductListResponse>("AppService", "getProducts", params);
}

export function getServices(params: ServiceListRequest): Promise<ServiceListResponse | never> {
    return call<ServiceListRequest, ServiceListResponse>("AppService", "getServices", params);
}

export function searchServices(params: ServiceSearchRequest): Promise<ServiceListResponse | never> {
    return call<ServiceSearchRequest, ServiceListResponse>("AppService", "searchServices", params);
}

export function getKeyList(params: KeyListRequest): Promise<KeyListResponse | never> {
    return call<KeyListRequest, KeyListResponse>("AppService", "getKeyList", params);
}

export function getKeyValueList(params: KeyValueListRequest): Promise<KeyValueListResponse | never> {
    return call<KeyValueListRequest, KeyValueListResponse>("AppService", "getKeyValueList", params);
}

export function searchKeyValueList(params: SearchKeyValueListRequest): Promise<KeyValueListResponse | never> {
    return call<SearchKeyValueListRequest, KeyValueListResponse>("AppService", "searchKeyValueList", params);
}

export function updateKeyValue(params: KeyValueRequest): Promise<KeyValueResponse | never> {
    return call<KeyValueRequest, KeyValueResponse>("AppService", "updateKeyValue", params);
}

export function fetchKeyHistory(params: KVHistoryRequest): Promise<KVHistoryResponse | never> {
    return call<KVHistoryRequest, KVHistoryResponse>("AppService", "fetchKeyHistory", params);
}
