import { defineStore } from "pinia" 
import authService from "../services/auth" 
import router from "../router/index"

export const useAuthStore = defineStore({
    id: "auth",
    state: () => ({
        user: null,
        login_token: null, 
        login_message: null, 
        register_data: null, 
        register_message: null, 
    }),

    //login logout
    actions: {
        handleLogin(data){
           console.log("login",data) 
           authService.userLogin(data).then(res => {
            console.log(res.data);
            if (res.data.status != "error") {
                console.log("home a yönlenidr");
                this.login_message = null
                this.login_token = res.data.data
                router.push('/')
            }
            if (res.data.status == "error") {
                console.log("loginede kal hata mesajı ver");
                this.login_token = null
                this.login_message = res.data.data
            }
            
           })
        },
        handleRegister(data) {
            console.log(data);
            authService.userRegister(data).then(res => {
                console.log(res.data)
                if (res.data.status != "error") {
                    console.log("logine yönlenidr");
                    this.register_data = res.data.data
                    this.login_token = res.data.data
                    router.push('/')
                }
                if (res.data.status == "error") {
                    console.log("registerda kal hata mesajı ver");
                    this.register_data =  null
                    this.register_message = res.data.message
                }
            }) 
        },
        handleLogout() {
            this.login_token = null
            router.push('/login') 
        }
    },

    getters: {
        getToken(state) {
            return state.login_token
        }
    }
})