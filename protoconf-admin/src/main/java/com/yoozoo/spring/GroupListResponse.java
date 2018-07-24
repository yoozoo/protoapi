package com.yoozoo.spring;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class GroupListResponse {
    private final List<Group> groups;

    @JsonCreator
    public GroupListResponse(@JsonProperty("groups") List<Group> groups) {
        this.groups = groups;
    }

    public List<Group> getGroups() {
        return groups;
    }
    
}
