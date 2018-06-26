import { Component, Vue } from 'vue-property-decorator';
import VueResource from "vue-resource";

Vue.use(VueResource);


interface HelloRequest {
    
     greeting: string
    
}

interface HelloResponse {
    
     reply: string
    
}


@Component
export default class HelloService extends Vue {
    constructor() {
        super()
    }

    // GET 
    SayHello(param: HelloRequest): PromiseLike<HelloResponse[] | Error> {
        /**
        *
        return this.$http.get().then(
            response => {
                console.log(response);
                return response.data as HelloResponse;
            },
            error => {
                console.log(error);
                return new Error("Something wrong");
            }
        );
        */
        
    }
}