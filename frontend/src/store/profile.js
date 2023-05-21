import { defineStore } from "pinia" 
import profileService from "../services/profile"
import router from "../router/index"

export const useProfileStore = defineStore({
    id: "profile",
    state: () => ({
        profile_data: null, 
        error_message_update: null,
        connection_data: null,
        message_connection: null,
        error_message_connection: null, 
        notifications_data: null
    }),

    //login logout
    actions: { 
        getProfile(token){ 
            profileService.getProfile(token).then(res => {
                if (res.data.status == "error") {
                    router.push('/login')
                }
                if (res.data.status != "error"){
                    console.log("bbbb");
                    this.profile_data = res.data
                    return this.profile_data
                } 
                return this.profile_data
            })
        },
        updateProfile(payload, token){ 
            profileService.updateProfile(payload, token).then(res => { 
                if (res.data.status == "error") {
                    this.error_message_update = res.data
                    router.push('/login')
                }
                if (res.data.status != "error"){
                    this.profile_data = res.data
                }   
            })
         },
        connections(token){ 
            profileService.connections(token).then(res => { 
                if (res.data.status == "error") {
                    this.error_message_connection = res.data
                    router.push('/login')
                }
                if (res.data.status != "error"){
                    this.connection_data = res.data
                }   
            })
         },
        connectionRequest(user_name, token){ 
            profileService.connectionRequest(user_name, token).then(res => { 
                if (res.data.status == "error") {
                    this.error_message_connection = res.data 
                }
                if (res.data.status != "error"){
                    this.message_connection = res.data
                }   
            })
         },
        connectionAccept(user_name, token){ 
            profileService.connectionAccept(user_name, token).then(res => { 
                if (res.data.status == "error") {
                    this.error_message_connection = res.data
                }
                if (res.data.status != "error"){
                    this.message_connection = res.data
                }   
            })
         },
        connectionIgnore(user_name, token){ 
            profileService.connectionIgnore(user_name, token).then(res => { 
                if (res.data.status == "error") {
                    this.error_message_connection = res.data
                }
                if (res.data.status != "error"){
                    this.message_connection = res.data
                }   
            })
        },
        notifications(){ 
            profileService.notifications().then(res => { 
                if (res.data.status == "error") {
                    this.notifications_data = res.data
                }
                if (res.data.status != "error"){
                    this.notifications_data = res.data
                }   
            })
         },
    }, 
    getters: {
        getProfileData(state) {
            return state.profile_data
        },
        getConnectioneData(state) {
            return state.connection_data
        }, 
        getNotifiationData(state) {
            return state.notifications_data
        }   
    }
})