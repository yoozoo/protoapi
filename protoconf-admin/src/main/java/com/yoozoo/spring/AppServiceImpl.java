package com.yoozoo.spring;

import com.yoozoo.protoconf.ConfigurationReader;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.CrossOrigin;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

@Controller
@CrossOrigin
public class AppServiceImpl extends AppServiceBase {
    @Autowired
    ConfigurationReader configurationReader;

    @Override
    GroupListResponse GetGroups(GroupListRequest in) {
        Map<String, String> keyValues = configurationReader.getValuesWithPrefix("/group");
        List<Group> groups = new ArrayList<>();

        if(keyValues == null || keyValues.isEmpty()) {
//            no key-value found
            return null;
        }

        for(Map.Entry<String, String> entry: keyValues.entrySet()) {
            if(entry.getKey().split("/").length == 2) {
//                get group only
                groups.add(new Group(entry.getKey(), entry.getValue()));
            }
        }

        if(groups.isEmpty()) {
//            no group found
            return null;
        }
        return new GroupListResponse(groups);
    }

    @Override
    AppListResponse GetApps(AppListRequest in) {
        Map<String, String> keyValues = configurationReader.getValuesWithPrefix(in.getGroup_id() + "/app");
        if(keyValues == null || keyValues.isEmpty()) {
//            no app found
            return null;
        }
        List<App> apps = new ArrayList<>();

        for(Map.Entry<String, String> entry: keyValues.entrySet()) {
            if(entry.getKey().split("/").length == 3) {
//                get app only
                apps.add(new App(entry.getKey(), entry.getValue()));
            }
        }

        String groupName = configurationReader.getValue(in.getGroup_id());
        return new AppListResponse(apps, in.getGroup_id(), groupName);
    }

    @Override
    KeyListResponse GetKeyList(KeyListRequest in) {
        Map<String, String> keyValues = configurationReader.getValuesWithPrefix(in.getApp_id() + "/");
        List<String> keys = new ArrayList<>();

        if(keyValues == null || keyValues.isEmpty()) {
//            no key found
            return null;
        }

        for(Map.Entry<String, String> entry: keyValues.entrySet()) {
            keys.add(entry.getKey());
        }

        String appName = configurationReader.getValue(in.getApp_id());
        return new KeyListResponse(keys, in.getApp_id(), appName);
    }

    @Override
    KeyValueListResponse GetKeyValueList(KeyValueListRequest in) {
        List<String> keys = in.getKeys();
        List<KeyValue> keyValues = new ArrayList<>();

        for(String key: keys) {
            keyValues.add(new KeyValue(key, configurationReader.getValue(key)));
        }

        if(keyValues.isEmpty()) {
//            no key found
            return null;
        }

        String[] keySplit = keys.get(0).split("/");
        if(keySplit.length < 3) {
//            wrong key format: correct format is /grp/app/*
            return null;
        }

        String groupId = "/" + keySplit[1];
        String groupName = configurationReader.getValue(groupId);

        String appId = "/" + keySplit[1] + "/" + keySplit[2];
        String appName = configurationReader.getValue(appId);

        return new KeyValueListResponse(keyValues, appId, appName, groupId, groupName);
    }

    @Override
    KeyValueResponse UpdateKeyValue(KeyValueRequest in) {
//        update key-value
        configurationReader.setValue(in.getKey_value().getKey(), in.getKey_value().getValue());

        String[] keySplit = in.getKey_value().getKey().split("/");
        if(keySplit.length < 3) {
//            wrong key format: correct format is /grp/app/*
            return null;
        }

        String groupId = "/" + keySplit[1];
        String groupName = configurationReader.getValue(groupId);

        String appId = "/" + keySplit[1] + "/" + keySplit[2];
        String appName = configurationReader.getValue(appId);

        return new KeyValueResponse(in.getKey_value(), appId, appName, groupId, groupName);
    }
}
