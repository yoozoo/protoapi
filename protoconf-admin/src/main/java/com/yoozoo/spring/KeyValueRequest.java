package com.yoozoo.spring;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class KeyValueRequest {
    private final KeyValue key_value;

    @JsonCreator
    public KeyValueRequest(@JsonProperty("key_value") KeyValue key_value) {
        this.key_value = key_value;
    }

    public KeyValue getKey_value() {
        return key_value;
    }
    
}
