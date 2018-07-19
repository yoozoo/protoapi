package com.yoozoo.spring;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import java.util.ArrayList;
import java.util.List;

@SpringBootApplication
public class Application {
    protected static List<Group> groups;
    protected static List<App> apps;
    protected static List<KeyValue> keys;

    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
//        group data
        groups = new ArrayList<>();
        Group appGroup = new Group("1", "app_server");
        groups.add(appGroup);
        Group gameGroup = new Group("2", "game_server");
        groups.add(gameGroup);

//        app data
        apps = new ArrayList<>();
        App app1 = new App("1", "common");
        apps.add(app1);
        App app2 = new App("2", "platform");
        apps.add(app2);
        App app3 = new App("3", "streaming");
        apps.add(app3);

        App app4 = new App("4", "女神联盟");
        apps.add(app4);
        App app5 = new App("5", "游族通");
        apps.add(app5);
        App app6 = new App("6", "盗墓笔记");
        apps.add(app6);

//        TODO: key-value data
        keys = new ArrayList<>();
        KeyValue keyValue1 = new KeyValue("/key1", "value1");
        keys.add(keyValue1);
        KeyValue keyValue2 = new KeyValue("/key2", "value2");
        keys.add(keyValue2);
        KeyValue keyValue3 = new KeyValue("/key3", "value3");
        keys.add(keyValue3);
        KeyValue keyValue4 = new KeyValue("/key4", "value4");
        keys.add(keyValue4);
        KeyValue keyValue5 = new KeyValue("/key5", "value5");
        keys.add(keyValue5);
        KeyValue keyValue6 = new KeyValue("/key6", "value6");
        keys.add(keyValue6);
        keys.add(new KeyValue("/key7", "value7"));
        keys.add(new KeyValue("/key8", "value8"));
        keys.add(new KeyValue("/key9", "value9"));
        keys.add(new KeyValue("/key10", "value10"));
        keys.add(new KeyValue("/key11", "value11"));
        keys.add(new KeyValue("/key12", "value12"));

    }

}
