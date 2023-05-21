<template>
    <div class="container">
        <img src="../assets/images/LI-Logo.png" alt="" style="height: 70px;">

        <div class="your-input">
            <div class="input">
                <input type="text" name="email" v-model="loginData.userName" id="email" required />
                <label for="email">Email</label>
            </div>
            <div class="input">
                <input type="password" name="password" v-model="loginData.password" id="password" required />
                <label for="password">
                    Password
                </label>
            </div>
        </div>

        <a href="#" class="forgot-password-link">
            Forgot Password?
        </a>
        <button @click="login()">Sign in</button>
        <p class="join-link" style="margin-left: 10%;">
            New to linkedin?
            <a href="/register" class="join-now">
                Join now
            </a>
        </p>

    </div>
</template>
  
<script>
import { useAuthStore } from '../store/auth'; 

export default {
    data() {
        return {
            loginData: {
                userName: '',
                password: ''
            },
        };
    },          
    methods: {
        login() {  
            const authStore = useAuthStore()
            const res = authStore.handleLogin(this.loginData)
            this.$cookies.set('token', authStore.getToken) 
        } 
    }
};
</script>

<style scoped>
* {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
    font-family: 'Roboto', sans-serif;
}

.container {
    width: 320px;
    height: 480px;
    background: #fff;
    box-shadow: 0 10px 15px rgba(179, 179, 179, 0.7);
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    border-radius: 7px;
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;
    padding: 1rem;
}

.container h2 {
    color: #0A66C3;
    font-size: 1.6rem;
}

.container h2 i {
    font-size: 1.8rem;
    padding-left: .1rem;
}

.text p {
    font-size: .8rem;
    padding-top: 0.3rem;
}

.input {
    position: relative;
    width: 100%;
    height: 50px;
    margin-bottom: 0.7em;
}

.your-input input {
    width: 100%;
    height: 50px;
    padding-top: 1.1rem;
    padding-left: 9px;
    outline: none;
    border: 1px solid #8c8c8c;
    border-radius: 3px;
    transition: .2s;
}

.your-input label {
    position: absolute;
    top: 30%;
    left: 10px;
    font-size: 1.1rem;
    color: #8c8c8c;
    transition: .2s;
}

.input input:focus~label,
.input input:valid~label {
    top: 10%;
    font-size: .8rem;
    color: #000;
}

.input input:focus {
    border-width: 2px;
    border-color: #0A66C3;
}

.forgot-password-link {
    width: 250px;
    padding: .2rem;
    text-align: center;
    text-decoration: none;
    font-weight: bolder;
    color: #0A66C3;
    transition: .3s;
    border-radius: 5px;
    margin-left: 7%;
}

.forgot-password-link:hover {
    background: rgba(10, 102, 195, 0.3);
}

button {
    height: 50px;
    background: #0A66C3;
    outline: none;
    border: none;
    border-radius: 30px;
    color: #fff;
    font-size: 1rem;
    font-weight: bolder;
}

.join-link {
    text-decoration: none;
    font-weight: bolder;
    color: #0A66C3;
}

.join-now {
    text-decoration: none;
    font-weight: bolder;
    color: #0A66C3;
    border-radius: 12px;
    transition: .3s;
    font-weight: bolder;
    padding: .2rem;
}

.join-now:hover {
    background: rgba(10, 102, 195, 0.3);
}
</style>
  