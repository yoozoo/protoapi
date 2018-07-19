package com.yoozoo.spring;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RequestBody;

public abstract class AppServiceBase {
    @PostMapping("/AppService.GetGroups")
    @ResponseBody
    public GroupListResponse GetGroupsPost(@RequestBody GroupListRequest in) {
        return GetGroups(in);
    }
    @GetMapping("/AppService.GetGroups")
    @ResponseBody
    public GroupListResponse GetGroupsGet(GroupListRequest in) {
        return GetGroups(in);
    }

    abstract GroupListResponse GetGroups(GroupListRequest in);
    
    @PostMapping("/AppService.GetApps")
    @ResponseBody
    public AppListResponse GetAppsPost(@RequestBody AppListRequest in) {
        return GetApps(in);
    }
    @GetMapping("/AppService.GetApps")
    @ResponseBody
    public AppListResponse GetAppsGet(AppListRequest in) {
        return GetApps(in);
    }

    abstract AppListResponse GetApps(AppListRequest in);
    
    @PostMapping("/AppService.GetKeyList")
    @ResponseBody
    public KeyListResponse GetKeyListPost(@RequestBody KeyListRequest in) {
        return GetKeyList(in);
    }
    @GetMapping("/AppService.GetKeyList")
    @ResponseBody
    public KeyListResponse GetKeyListGet(KeyListRequest in) {
        return GetKeyList(in);
    }

    abstract KeyListResponse GetKeyList(KeyListRequest in);
    
    @PostMapping("/AppService.GetKeyValueList")
    @ResponseBody
    public KeyValueListResponse GetKeyValueListPost(@RequestBody KeyValueListRequest in) {
        return GetKeyValueList(in);
    }
    @GetMapping("/AppService.GetKeyValueList")
    @ResponseBody
    public KeyValueListResponse GetKeyValueListGet(KeyValueListRequest in) {
        return GetKeyValueList(in);
    }

    abstract KeyValueListResponse GetKeyValueList(KeyValueListRequest in);
    
    @PostMapping("/AppService.UpdateKeyValue")
    @ResponseBody
    public KeyValueResponse UpdateKeyValuePost(@RequestBody KeyValueRequest in) {
        return UpdateKeyValue(in);
    }
    @GetMapping("/AppService.UpdateKeyValue")
    @ResponseBody
    public KeyValueResponse UpdateKeyValueGet(KeyValueRequest in) {
        return UpdateKeyValue(in);
    }

    abstract KeyValueResponse UpdateKeyValue(KeyValueRequest in);
    
}
