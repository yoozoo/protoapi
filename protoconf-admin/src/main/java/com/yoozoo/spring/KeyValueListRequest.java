package com.yoozoo.spring;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class KeyValueListRequest {
    private final List<String> keys;

    @JsonCreator
    public KeyValueListRequest(@JsonProperty("keys") List<String> keys) {
        this.keys = keys;
    }

    public List<String> getKeys() {
        return keys;
    }
    
}
