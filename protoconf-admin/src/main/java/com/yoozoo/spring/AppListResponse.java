package com.yoozoo.spring;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class AppListResponse {
    private final List<App> apps;
    private final String group_id;
    private final String group_name;

    @JsonCreator
    public AppListResponse(@JsonProperty("apps") List<App> apps, @JsonProperty("group_id") String group_id, @JsonProperty("group_name") String group_name) {
        this.apps = apps;
        this.group_id = group_id;
        this.group_name = group_name;
    }

    public List<App> getApps() {
        return apps;
    }
    public String getGroup_id() {
        return group_id;
    }
    public String getGroup_name() {
        return group_name;
    }
    
}
