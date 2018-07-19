package com.yoozoo.spring;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.CrossOrigin;

import java.util.ArrayList;
import java.util.List;

@Controller
@CrossOrigin
public class AppServiceImpl extends AppServiceBase {
    @Override
    GroupListResponse GetGroups(GroupListRequest in) {
        return new GroupListResponse(Application.groups);
    }

    @Override
    AppListResponse GetApps(AppListRequest in) {
        int startIndex = Integer.parseInt(in.getGroup_id()) - 1;
        return new AppListResponse(Application.apps.subList(startIndex * 3, (startIndex + 1) * 3 ), in.getGroup_id(), Application.groups.get(startIndex).getGroup_name());
    }

    @Override
    KeyListResponse GetKeyList(KeyListRequest in) {
        int appId = Integer.parseInt(in.getApp_id()) - 1;
        List<KeyValue> keyValues = Application.keys.subList(appId*2, appId*2 + 2);
        List<String> keys = new ArrayList<>();
        for (KeyValue keyValue : keyValues) {
            keys.add(keyValue.getKey());
        }
        return new KeyListResponse(keys, in.getApp_id(), Application.apps.get(appId).getApp_name());
    }

    @Override
    KeyValueListResponse GetKeyValueList(KeyValueListRequest in) {
        List<String> keys = in.getKeys();
        List<KeyValue> keyValues = new ArrayList<>();
        int keyId = 0;
        int i = 0;
        for (KeyValue keyValue : Application.keys) {
            i++;
            if (keys.contains(keyValue.getKey())) {
                keyValues.add(keyValue);
                keyId = i;
            }
        }
        int appId = (keyId - 1) /2 + 1;
        String appName = Application.apps.get(appId - 1).getApp_name();
        int groupId = (appId - 1) / 3 + 1;
        String groupName = Application.groups.get(groupId - 1).getGroup_name();
        return new KeyValueListResponse(keyValues, String.valueOf(appId), appName, String.valueOf(groupId), groupName);
    }

    @Override
    KeyValueResponse UpdateKeyValue(KeyValueRequest in) {
        int keyId = 0;
        for (KeyValue keyValue : Application.keys) {
            keyId++;
            if (keyValue.getKey().equalsIgnoreCase(in.getKey_value().getKey())) {
                Application.keys.set(keyId - 1, in.getKey_value());
                int appId = (keyId - 1) /2 + 1;
                String appName = Application.apps.get(appId - 1).getApp_name();
                int groupId = (appId - 1) / 3 + 1;
                String groupName = Application.groups.get(groupId - 1).getGroup_name();
                return new KeyValueResponse(in.getKey_value(), String.valueOf(appId), appName, String.valueOf(groupId), groupName);
            }
        }

        return null;
    }
}
