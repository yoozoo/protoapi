import {
    msg,
} from './EchoServiceObjs';
import * as svr from './EchoService';

svr.SetBaseUrl("http://localhost:8080");

var msg: msg = {
    "msg": "hello world"
};

svr.echo(msg).then(m => {
    console.log(m.msg);
}).catch(e => {
    console.log(e);
});
