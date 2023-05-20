<template>
  <b-container class="bv-example-row">
        <b-row >
            <h4 class="pl-4">Networks</h4>
        </b-row>
        <b-row >
            <b-col> 
                <networkInfos></networkInfos> 
            </b-col>
            <b-col cols="8"> 
                <connectRequest></connectRequest>
                <connectRequest></connectRequest>
                <connectRequest></connectRequest>
                <connectRequest></connectRequest>
            </b-col> 
        </b-row>
    </b-container>
</template>

<script>
import connectRequest from '../components/Network/connectRequest.vue';
import networkInfos from '../components/Network/networkInfos.vue';
export default {
  components: {
    connectRequest,
    networkInfos
  },
    data() {
    return {
      connections: [
        { id: 1, name: 'John Doe' },
        { id: 2, name: 'Jane Smith' },
        { id: 3, name: 'Mike Johnson' }
      ]
    };
  }
};
</script>