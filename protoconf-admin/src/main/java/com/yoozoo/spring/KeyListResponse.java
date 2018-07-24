package com.yoozoo.spring;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class KeyListResponse {
    private final List<String> keys;
    private final String app_id;
    private final String app_name;

    @JsonCreator
    public KeyListResponse(@JsonProperty("keys") List<String> keys, @JsonProperty("app_id") String app_id, @JsonProperty("app_name") String app_name) {
        this.keys = keys;
        this.app_id = app_id;
        this.app_name = app_name;
    }

    public List<String> getKeys() {
        return keys;
    }
    public String getApp_id() {
        return app_id;
    }
    public String getApp_name() {
        return app_name;
    }
    
}
