package com.yoozoo.spring;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class App {
    private final String app_id;
    private final String app_name;

    @JsonCreator
    public App(@JsonProperty("app_id") String app_id, @JsonProperty("app_name") String app_name) {
        this.app_id = app_id;
        this.app_name = app_name;
    }

    public String getApp_id() {
        return app_id;
    }
    public String getApp_name() {
        return app_name;
    }
    
}
