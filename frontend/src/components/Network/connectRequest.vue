<template>
    <div>
        
        <div class="card mb-2" v-if="connectionReq.status == false && req_or_not== 'req'"> 
        <div class="card-body">
            <b-container class="bv-example-row">
                <b-row>
                    <b-col cols="8">
                        <b-row>
                            <b-col>
                                <div class="mb-2">
                                    <b-avatar src="https://placekitten.com/300/300" size="4rem"></b-avatar>
                                </div>
                            </b-col>
                            <b-col cols="8">
                                <b-row>
                                    <h5>{{ connectionReq.sender.first_name }} {{ connectionReq.sender.last_name }}</h5>
                                </b-row>
                                <b-row>
                                    <p class="font-weight-lighter">{{ connectionReq.sender.headline }}</p>
                                </b-row>
                            </b-col>
                        </b-row>
                    </b-col>
                    <b-col class="mt-1 pl-5" cols="4">
                        <div>
                            <b-button @click="ignoreReq(connectionReq.sender.user_name)" pill variant="outline-danger"
                                class="mr-2">Ignore</b-button>
                            <b-button @click="acceptReq(connectionReq.sender.user_name)" pill class="mt-1"
                                variant="outline-primary">Accept</b-button>
                        </div>
                    </b-col>
                </b-row>
            </b-container>
        </div>
    </div>

    <div class="card mb-2" v-if=" req_or_not== 'not'"> 
        <div class="card-body">
            <b-container class="bv-example-row">
                <b-row>
                    <b-col cols="8">
                        <b-row>
                            <b-col>
                                <div class="mb-2">
                                    <b-avatar src="https://placekitten.com/300/300" size="4rem"></b-avatar>
                                </div>
                            </b-col>
                            <b-col cols="7">
                                <b-row>
                                   <a :href="`user/${connectionReq.sender.user_name}`">
                                    <h5>{{ connectionReq.sender.first_name }} {{ connectionReq.sender.last_name }}</h5>
                                   </a>
                                </b-row>
                                <b-row>
                                    <p class="font-weight-lighter">{{ connectionReq.sender.headline }}</p>
                                </b-row>
                            </b-col>
                            
                        </b-row>
                    </b-col>
                    <b-col class="mt-1 pl-5" cols="4">
                        <div>
                            <b-button @click="removeConnection(connectionReq.sender.user_name)" pill variant="outline-danger"
                                class="mr-2">Remove</b-button> 
                        </div>
                    </b-col>
                </b-row>
            </b-container>
        </div>
    </div>
    </div>
</template>
  
<script>
import axios from 'axios';
import router from '../../router';

export default {
    name: "connectReq",
    props: ["connectionReq", "req_or_not"],
    data() {
        return {

            user: {
                name: "Volkan Öztoklu",
                info: "Mef University"
            }
        }
    },
    methods: {
        ignoreReq(data) { 
            axios.delete(`https://software.ardapektezol.com/api/connections/${data}`, {
                headers: {
                    Authorization: this.$cookies.get("token")
                }
            }).then(res => {
                console.log(res );
                router.push("/network")
            })
        },
        acceptReq(data) { 
            axios.put(`https://software.ardapektezol.com/api/connections/${data}`,{} ,{
                headers: {
                    Authorization: this.$cookies.get("token")
                }
            }).then(res => {
                console.log(res );
                router.push("/network")
            })
        },
        removeConnection(data){
            axios.delete(`https://software.ardapektezol.com/api/connections/${data}`, {
                headers: {
                    Authorization: this.$cookies.get("token")
                }
            }).then(res => {
                console.log(res );
                router.push("/network")
            })
        }
    }

}
</script>
  