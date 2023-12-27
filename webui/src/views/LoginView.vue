
<script setup>
import { inject } from 'vue'
import { useRouter } from "vue-router";
import { setCookie } from "../utils/cookieUtils.js";
const { loggedIn, setLoggedIn } = inject('loggedIn')
const router = useRouter();
</script>  

<script>


export default {
    data() {
        return {
            name: ''
        };
    },
    setup() {
        // const router = useRouter();
        // return { router };

    },
    methods: {
        async login() {
            try {
                const response = await this.$axios.post('http://localhost:3000/login', { name: this.name });
                setCookie('identifier', response.data.identifier);
                this.setLoggedIn(true);
                this.router.push('/');
            } catch (error) {
                console.error(error);
            }
        }
    }
};
</script>
<template>

        <div class="login-container">
    
            <form @submit.prevent="login">
                <h2>Login</h2>
                <input type="text" v-model="name" placeholder="Enter your name" />
                <button type="submit">Submit</button>
            </form>
        </div>
</template>

  
<style>
.login-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
}

.login-container form {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.login-container input {
    margin: 10px;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
}

.login-container button {
    padding: 10px 20px;
    border: none;
    border-radius: 5px;
    background-color: #007bff;
    color: white;
    cursor: pointer;
}

.login-container button:hover {
    background-color: #0056b3;
}
</style>
  