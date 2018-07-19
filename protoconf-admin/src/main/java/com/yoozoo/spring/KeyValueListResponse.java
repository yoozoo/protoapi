package com.yoozoo.spring;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class KeyValueListResponse {
    private final List<KeyValue> key_values;
    private final String app_id;
    private final String app_name;
    private final String group_id;
    private final String group_name;

    @JsonCreator
    public KeyValueListResponse(@JsonProperty("key_values") List<KeyValue> key_values, @JsonProperty("app_id") String app_id, @JsonProperty("app_name") String app_name, @JsonProperty("group_id") String group_id, @JsonProperty("group_name") String group_name) {
        this.key_values = key_values;
        this.app_id = app_id;
        this.app_name = app_name;
        this.group_id = group_id;
        this.group_name = group_name;
    }

    public List<KeyValue> getKey_values() {
        return key_values;
    }
    public String getApp_id() {
        return app_id;
    }
    public String getApp_name() {
        return app_name;
    }
    public String getGroup_id() {
        return group_id;
    }
    public String getGroup_name() {
        return group_name;
    }
    
}
