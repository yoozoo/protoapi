package com.yoozoo.spring;

import com.yoozoo.protoconf.ConfigurationReader;
import com.yoozoo.protoconf.EtcdReader;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

@SpringBootApplication
public class Application {

    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }

    @Bean
    public ConfigurationReader configurationReader() {
//        connect to etcd server
        return new ConfigurationReader(new EtcdReader("http://192.168.115.57:2379", "root", "root"));
    }

}
