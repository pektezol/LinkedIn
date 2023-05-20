<template>
    <div class="card mb-2">
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
                            <b-col cols="9">
                                <b-row>
                                    <h5>{{user.name}}</h5>
                                </b-row>
                                <b-row> 
                                    <p class="font-weight-lighter">{{user.info}}</p> 
                                </b-row>
                            </b-col>
                        </b-row>
                    </b-col>
                    <b-col class="mt-3 pl-5" cols="4">
                        <div>
                            <b-button pill variant="outline-danger" class="mr-2">Ignore</b-button>
                            <b-button pill variant="outline-primary">Accept</b-button>
                        </div>
                    </b-col>
                </b-row>
            </b-container>
        </div>
    </div>
</template>
  
<script> 
export default {
    name: "connectReq",
    data() {
        return {
            user: {
                name: "Volkan Öztoklu",
                info: "Mef University"
            }
        }
    }
}
</script>
  