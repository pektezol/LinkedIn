 
<template>
    <div class="card mb-2">
        <div>
            <b-list-group>

                <b-list-group-item class="d-flex align-items-center">
                    <span class="mr-auto">Manage my network</span>
                </b-list-group-item>

                <b-list-group-item class="btn d-flex align-items-center">
                    <span class="mr-auto">Connections</span>
                    <b-badge>{{connection_count_prop}}</b-badge>
                </b-list-group-item>
                 
            </b-list-group>
        </div>
    </div>
</template>
  
<script>
export default {
    name: "networkInfo",
    props: ["connection_count_prop"],
    data() {
        return {
             
        }
    },
    created(){
        console.log("info", this.connectionReq);
    }
}
</script>
  