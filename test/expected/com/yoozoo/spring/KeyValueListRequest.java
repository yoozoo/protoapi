// Code generated by protoapi; DO NOT EDIT.

package com.yoozoo.spring;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class KeyValueListRequest {
    private final int service_id;
    private final List<Key> keys;

    @JsonCreator
    public KeyValueListRequest(@JsonProperty("service_id") int service_id, @JsonProperty("keys") List<Key> keys) {
        this.service_id = service_id;
        this.keys = keys;
    }

    public int getService_id() {
        return service_id;
    }
    public List<Key> getKeys() {
        return keys;
    }
    
}
