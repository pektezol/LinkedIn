<template>
  <b-container class="bv-example-row">
    <b-row>
      <h4 class="pl-4">Networks</h4>
    </b-row>
    <b-row>

      <b-col>
        <networkInfos :connection_count_prop="connection_count"></networkInfos>
      </b-col>

      <b-col cols="8">
        <div class="row">
          <b-col cols="12">
            <h4>Connection Reqests</h4>
          </b-col>
          <b-row class="ml-4">
            <div v-for="(connection) in connections_data" :key="connection.id">
              <connectRequest :connection-req="connection" req_or_not="req"></connectRequest>
            </div>
          </b-row> 
        </div>
        <div class="row">
          <b-col cols="12">
            <h4>All Connections</h4>
          </b-col>
          <b-row class="ml-4">
            <div v-for="(connection) in connections_data" :key="connection.id">
              <connectRequest :connection-req="connection" req_or_not="not"></connectRequest>
            </div>
          </b-row> 
        </div>

      </b-col>
    </b-row>
  </b-container>
</template>

<script>
import axios from 'axios';
import connectRequest from '../components/Network/connectRequest.vue';
import networkInfos from '../components/Network/networkInfos.vue';

export default {
  components: {
    connectRequest,
    networkInfos
  },
  data() {
    return {
      connections_data: null,
      connection_count: 0,
      connections: [
        { id: 1, name: 'John Doe' },
        { id: 2, name: 'Jane Smith' },
        { id: 3, name: 'Mike Johnson' }
      ]
    };
  },
  async mounted() {
    await axios.get(`https://software.ardapektezol.com/api/connections`, {
      headers: {
        Authorization: this.$cookies.get("token")
      }
    }).then(res => {
      this.connections_data = res.data.data.connections
    })
    this.calculate_connection_count()
  },
  methods: {
    calculate_connection_count() {
      console.log("değer", this.connections_data);
      for (let index = 0; index < this.connections_data.length; index++) {
        if (this.connections_data[index].status == true) {
          this.connection_count = this.connection_count + 1
        }
        console.log(this.connection_count);

      }
    }
  }
};
</script>